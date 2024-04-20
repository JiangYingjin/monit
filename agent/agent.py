#!python

import os, sys, psutil, time, json, requests, datetime, subprocess, re, sqlite3, redis, pymongo, pymysql, argparse, hashlib, threading, yaml, logging, fcntl


# ===================== 版本检查更新 =====================

# 获取本文件的位置
cwd = os.path.dirname(os.path.abspath(__file__))
# 如果本文件位于 /usr/local/monit
if cwd == "/usr/local/monit":
    print("检查 Agent 是否有新版本")
    # 计算本文件的 md5sum 与 https://file.jiangyj.tech/proj/monit/agent.py.checksum 进行比对
    local_md5sum = (
        subprocess.run(f"md5sum {cwd}/agent.py", shell=True, stdout=subprocess.PIPE)
        .stdout.decode()
        .split()[0]
    )
    latest_md5sum = requests.get(
        "https://file.jiangyj.tech/proj/monit/agent.py.checksum"
    ).text.strip()
    print(f"本地版本：{local_md5sum}")
    print(f"最新版本：{latest_md5sum}")

    # 如果不一致，则下载最新的 agent.py 并替换本文件
    if local_md5sum != latest_md5sum:
        print("Agent 有新版本！正在更新 ...")
        # 从 https://file.jiangyj.tech/proj/monit/agent.py 下载最新版本的 agent.py 替换本文件
        # 原封不动地保留本文件的所有入参并重新执行
        subprocess.run(
            f"wget -L https://file.jiangyj.tech/proj/monit/agent.py -O /usr/local/monit/agent.py && chmod +x /usr/local/monit/agent.py && python /usr/local/monit/agent.py {' '.join(sys.argv[1:])}",
            shell=True,
            # stdout=subprocess.DEVNULL,
            # stderr=subprocess.DEVNULL,
        )
        print("Agent 更新完成！重新启动 ...")
        exit(0)
    else:
        print("Agent 已是最新版本！")


# ===================== 维持保活 =====================

# 参考 * * * * * root flock -xn /tmp/stargate.lock -c '/usr/local/qcloud/stargate/admin/start.sh > /dev/null 2>&1 &'
# flock 为文件锁命令，-xn 以非阻塞模式获取锁，若无法获取锁则立即退出，若可获取则执行 -c 的入参命令
open("/etc/cron.d/monit", "w").write(
    "* * * * * root flock -xn /tmp/monit.lock -c 'python /usr/local/monit/agent.py monit --cron > /dev/null 2>&1 &'\n"
)


# ===================== 日志流配置 =====================

logging.basicConfig(
    level=logging.INFO, format="%(asctime)s - %(levelname)s - %(message)s"
)  # 目前为打印到屏幕


# ===================== 声明入参解析 =====================

parser = argparse.ArgumentParser()
subparsers = parser.add_subparsers(dest="subcommand")

# 初始化 Agent
init_parser = subparsers.add_parser("init", help="初始化 Agent")
init_parser.add_argument("--server-ip", type=str, help="服务器 IP 地址")
init_parser.add_argument("--machine-id", type=int, default=True, help="机器 ID")
init_parser.add_argument("--password", type=str, default=True, help="机器密码")

# 配置参数解析
configure_parser = subparsers.add_parser("configure", help="配置 Agent")

configure_parser.add_argument(
    "--mysql-enable", type=bool, required=True, help="是否启用 MySQL 监控"
)
configure_parser.add_argument("--mysql-port", type=int, default=3306, help="MySQL 端口")
configure_parser.add_argument(
    "--mysql-user", type=str, default="root", help="MySQL 用户"
)
configure_parser.add_argument("--mysql-password", type=str, help="MySQL 密码")

configure_parser.add_argument(
    "--redis-enable", type=bool, required=True, help="是否启用 Redis 监控"
)
configure_parser.add_argument("--redis-port", type=int, default=6379, help="Redis 端口")
configure_parser.add_argument("--redis-password", type=str, help="Redis 密码")

configure_parser.add_argument(
    "--nginx-enable", type=bool, default=True, help="是否启用 Nginx 监控"
)
configure_parser.add_argument("--nginx-path", type=str, help="Nginx 安装路径")

configure_parser.add_argument(
    "--php-fpm-enable", type=bool, required=True, help="是否启用 PHP-FPM 监控"
)
configure_parser.add_argument("--php-fpm-path", type=str, help="PHP-FPM 安装路径")

# 开始监控参数解析
monit_parser = subparsers.add_parser("monit", help="开始监控")
monit_parser.add_argument("--cron", type=bool, default=False, help="保活")

# 停止监控参数解析
stop_parser = subparsers.add_parser("stop", help="停止监控")

# 卸载命令参数解析
uninstall_parser = subparsers.add_parser("uninstall", help="在远程主机上卸载 Agent")

args = parser.parse_args()
logging.info(f"入参：{args}")


# ===================== Agent 类 =====================


class Agent:
    def __init__(self, machine_id: int = None, password: str = None):

        # 应尽可能早地初始化数据库
        self.cwd = os.path.dirname(os.path.abspath(__file__))
        self._db_init()

        # 初始化 Agent
        if machine_id and password:
            self.db_dct("machine_id", machine_id)
            self.db_dct("password", password)

        self.parse_config()

        # 初始化时从本地数据库中取出一次 token（运行过程中不再从数据库中取出）
        self._machine_id = self.db_dct("machine_id")
        self._token = self.db_dct("Token")
        self._token_expires = self.db_dct("Token.ExpiresAt")

        self.detect_services = {
            "mysqld": "mysql",
            "redis-server": "redis",
            "nginx": "nginx",
            "php-fpm": "php-fpm",
        }

        self.base_url = f"http://{args.server_ip}:8888"
        self.uri_dct = {
            "sign": "/machine/login",
            "upload": "/machine/uploadDataMulti",
            "service": {
                "create": "/machine/createMachineService",
                "update": "/machine/updateMachineService",
            },
        }

        self.running_services = None
        self.update_service_status()

    def parse_config(self):
        if os.path.exists(f"{self.cwd}/config.yml"):
            with open(f"{self.cwd}/config.yml", "r", encoding="utf-8") as f:
                self.config = dict(yaml.safe_load(f))
                logging.info(f"本地配置（config.yml）: {self.config}")

    def save_config(self):
        with open(f"{self.cwd}/config.yml", "w", encoding="utf-8") as f:
            yaml.dump(self.config, f)

    def token(self):
        if self._token:
            # Token 的有效期是一个星期，如果有效期仍大于两个小时，则不再重新签名
            if time.time() - self._token_expires / 1000 < 7200:
                # print(datetime.datetime.fromtimestamp(time.time()))
                # print(datetime.datetime.fromtimestamp(self._token_expires / 1000))
                return self._token

        # 否则重新签名
        r = requests.post(
            self.base_url + self.uri_dct["sign"],
            json={
                "machine_id": str(self._machine_id),
                "password": self.db_dct("password"),
            },
        )

        """
        {"code":0,"data":{"ExpiresAt":1713674791000,"Machine":{"ID":3,"CreatedAt":"2024-04-14T11:08:00.813+08:00","UpdatedAt":"2024-04-14T11:08:00.813+08:00","name":"041411","description":"1","ip_addr":"111.230.30.196","password":"admin","status":true,"CreatedBy":1,"UpdatedBy":0,"DeletedBy":0},"Token":"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJNYWNoaW5lSUQiOiIzIiwiQnVmZmVyVGltZSI6ODY0MDAsImlzcyI6InFtUGx1cyIsImF1ZCI6WyJHVkEiXSwiZXhwIjoxNzEzNjc0NzkxLCJuYmYiOjE3MTMwNjk5OTF9.a1jVZ52Y-BnAbj0ZGLzioA-CsjNhxybsngIGAC8NczU"},"msg":"登录成功"}"""

        # 如果请求成功
        if r.status_code == 200:
            res = r.json()["data"]
            # 重新更新对象变量中的 Token 和 ExpiresAt
            self._token = res["Token"]
            self._token_expires = res["ExpiresAt"]
            # 重新更新数据库中的 Token 和 ExpiresAt
            self.db_dct("Token", self._token)
            self.db_dct("Token.ExpiresAt", self._token_expires)
            return self._token
        else:
            logging.error("签名失败！")
            logging.error(r.text)
            return self._token

    def cpu(self):
        # load_avg = psutil.getloadavg()
        cpu = {
            "cpu_percent": psutil.cpu_percent(),
            "cpu_freq": psutil.cpu_freq().current,
            # "load_avg_1min": load_avg[0],
            # "load_avg_5min": load_avg[1],
            # "load_avg_15min": load_avg[2],
        }
        self.db_insert("cpu", cpu)
        """
{"cpu_percent": 7.0, "cpu_freq": 2494.142, "load_avg_1min": 0.37109375, "load_avg_5min": 0.5419921875, "load_avg_15min": 0.5576171875}

cpu_percent     keep    CPU 利用率
cpu_freq        keep    CPU 频率
load_avg_1min   keep    1 分钟负载
load_avg_5min   keep    5 分钟负载
load_avg_15min  keep    15 分钟负载
"""
        self.add_packet("cpu.cpu_percent", cpu["cpu_percent"])
        self.add_packet("cpu.cpu_freq", cpu["cpu_freq"])
        # self.add_packet("cpu.load_avg_1min", cpu["load_avg_1min"])
        # self.add_packet("cpu.load_avg_5min", cpu["load_avg_5min"])
        # self.add_packet("cpu.load_avg_15min", cpu["load_avg_15min"])

    def memory(self):
        memory = psutil.virtual_memory()._asdict()
        swap_memory = psutil.swap_memory()._asdict()
        # 在 swap_memory 的键名前加上 swap_
        swap_memory = {f"swap_{k}": v for k, v in swap_memory.items()}
        # 将 virtual_memory 和 swap_memory 合并成一个字典
        memory.update(swap_memory)
        self.db_insert("memory", memory)
        """
{"total": 7786668032, "available": 1870245888, "percent": 76.0, "used": 5496979456, "free": 191303680, "active": 1199513600, "inactive": 5651271680, "buffers": 319791104, "cached": 1778593792, "shared": 97697792, "slab": 554811392, "swap_total": 4294963200, "swap_used": 105005056, "swap_free": 4189958144, "swap_percent": 2.4, "swap_sin": 16244736, "swap_sout": 234557440}

memory.total        keep    总内存
memory.available    keep    可用内存
memory.percent      keep    内存利用率
memory.used         keep    已用内存
memory.free         keep    空闲内存
memory.active       keep    活跃内存
memory.inactive     keep    非活跃内存
memory.buffers      keep    缓冲区
memory.cached       keep    缓存
memory.shared       keep    共享内存
memory.slab         keep    内核缓存
memory.swap_total   keep    总交换内存
memory.swap_used    keep    已用交换内存
memory.swap_free    keep    空闲交换内存
memory.swap_percent keep    交换内存利用率
memory.swap_sin     keep    交换内存扇入
memory.swap_sout    keep    交换内存扇出
"""
        MB = 1024**2
        self.add_packet("memory.total", memory["total"] / MB)
        self.add_packet("memory.available", memory["available"] / MB)
        self.add_packet("memory.percent", memory["percent"])
        self.add_packet("memory.used", memory["used"] / MB)
        self.add_packet("memory.free", memory["free"] / MB)
        self.add_packet("memory.active", memory["active"] / MB)
        self.add_packet("memory.inactive", memory["inactive"] / MB)
        self.add_packet("memory.buffers", memory["buffers"] / MB)
        self.add_packet("memory.cached", memory["cached"] / MB)
        self.add_packet("memory.shared", memory["shared"] / MB)
        self.add_packet("memory.slab", memory["slab"] / MB)
        self.add_packet("memory.swap_total", memory["swap_total"] / MB)
        self.add_packet("memory.swap_used", memory["swap_used"] / MB)
        self.add_packet("memory.swap_free", memory["swap_free"] / MB)
        self.add_packet("memory.swap_percent", memory["swap_percent"])
        self.add_packet("memory.swap_sin", memory["swap_sin"] / MB)
        self.add_packet("memory.swap_sout", memory["swap_sout"] / MB)

    def disk(self):
        disk = psutil.disk_usage("/")._asdict()
        self.db_insert("disk", disk)
        """
{"total": 499963174912, "used": 1233622016, "free": 473648056576, "percent": 0.2}

disk.total      keep        磁盘总大小
disk.used       keep        磁盘已用大小
disk.free       keep        磁盘空闲大小
disk.percent    keep    磁盘利用率
"""
        GB = 1024**3
        self.add_packet("disk.total", disk["total"] / GB)
        self.add_packet("disk.used", disk["used"] / GB)
        self.add_packet("disk.free", disk["free"] / GB)
        self.add_packet("disk.percent", disk["percent"])

    def disk_io(self):
        """
        {"read_count": 328322, "write_count": 4106315, "read_bytes": 5682122240, "write_bytes": 52604564480, "read_time": 294060, "write_time": 5195645, "read_merged_count": 69506, "write_merged_count": 4112906, "busy_time": 5106976}

        disk_io.read_count          subtract    磁盘读次数
        disk_io.write_count         subtract    磁盘写次数
        disk_io.read_bytes          subtract    磁盘读大小（字节）
        disk_io.write_bytes         subtract    磁盘写大小（字节）
        disk_io.read_time           subtract    磁盘读时长（ms）
        disk_io.write_time          subtract    磁盘写时长（ms）
        disk_io.read_merged_count   subtract    磁盘合并读次数
        disk_io.write_merged_count  subtract    磁盘合并写次数
        disk_io.busy_time           subtract    磁盘繁忙时长（ms）
        """
        disk_io = psutil.disk_io_counters()._asdict()
        self.db_insert("disk_io", disk_io)

        curr_record = disk_io
        last_record = self._last_record("disk_io")
        # print(last_record)
        if last_record:
            self.add_packet(
                "disk_io.read_count",
                curr_record["read_count"] - last_record["read_count"],
            )
            self.add_packet(
                "disk_io.write_count",
                curr_record["write_count"] - last_record["write_count"],
            )
            self.add_packet(
                "disk_io.read_bytes",
                curr_record["read_bytes"] - last_record["read_bytes"],
            )
            self.add_packet(
                "disk_io.write_bytes",
                curr_record["write_bytes"] - last_record["write_bytes"],
            )
            self.add_packet(
                "disk_io.read_time",
                curr_record["read_time"] - last_record["read_time"],
            )
            self.add_packet(
                "disk_io.write_time",
                curr_record["write_time"] - last_record["write_time"],
            )
            self.add_packet(
                "disk_io.read_merged_count",
                curr_record["read_merged_count"] - last_record["read_merged_count"],
            )
            self.add_packet(
                "disk_io.write_merged_count",
                curr_record["write_merged_count"] - last_record["write_merged_count"],
            )
            self.add_packet(
                "disk_io.busy_time",
                curr_record["busy_time"] - last_record["busy_time"],
            )

    def net_io(self):
        net_io = psutil.net_io_counters()._asdict()
        self.db_insert("net_io", net_io)
        """
{"bytes_sent": 4371581206, "bytes_recv": 3883213301, "packets_sent": 15500445, "packets_recv": 15033879, "errin": 0, "errout": 0, "dropin": 0, "dropout": 0}

net_io.bytes_sent       subtract    网络发送字节数
net_io.bytes_recv       subtract    网络接收字节数
net_io.packets_sent     subtract    网络发送包数
net_io.packets_recv     subtract    网络接收包数
net_io.errin            subtract    网络接收错误数
net_io.errout           subtract    网络发送错误数
net_io.dropin           subtract    网络接收丢包数
net_io.dropout          subtract    网络发送丢包数
"""
        curr_record = net_io
        last_record = self._last_record("net_io")
        if last_record:
            self.add_packet(
                "net_io.bytes_sent",
                curr_record["bytes_sent"] - last_record["bytes_sent"],
            )
            self.add_packet(
                "net_io.bytes_recv",
                curr_record["bytes_recv"] - last_record["bytes_recv"],
            )
            self.add_packet(
                "net_io.packets_sent",
                curr_record["packets_sent"] - last_record["packets_sent"],
            )
            self.add_packet(
                "net_io.packets_recv",
                curr_record["packets_recv"] - last_record["packets_recv"],
            )
            self.add_packet("net_io.errin", curr_record["errin"] - last_record["errin"])
            self.add_packet(
                "net_io.errout", curr_record["errout"] - last_record["errout"]
            )
            self.add_packet(
                "net_io.dropin", curr_record["dropin"] - last_record["dropin"]
            )
            self.add_packet(
                "net_io.dropout", curr_record["dropout"] - last_record["dropout"]
            )

    def mysql_status(self):
        key_index = [
            "Queries",
            "Questions",
            "Com_insert",
            "Com_update",
            "Com_update_multi",
            "Com_delete",
            "Com_delete_multi",
            "Com_select",
            "innodb_buffer_pool_reads",
            "innodb_buffer_pool_read_ahead",
            "innodb_buffer_pool_read_requests",
            "Slow_queries",
            "Aborted_connects",
            "Connection_errors_max_connections",
            "Threads_connected",
            "Threads_running",
            "Connections",
            # 每建立一个连接，都需要一个线程来与之匹配
            "Max_used_connections",
            "max_connections",
        ]
        key_index_str = "|".join(key_index)
        sql = f"show global status where variable_name regexp '^({key_index_str})$';"
        mysql_status = self.mysql_exec(
            sql,
            port=self.config["mysql"]["port"],
            passwd=self.config["mysql"]["password"],
        )

        # 重新组织 mysql_status，将其转换成字典
        mysql_status = {
            d["Variable_name"]: (
                int(d["Value"]) if type(d["Value"]) != type(0) else d["Value"]
            )
            for d in mysql_status
        }
        self.db_insert("mysql", mysql_status)

        """
        show global status where variable_name in ('Queries', 'uptime');
            请求数量、查询数量
            QPS = (Queries2 -Queries1) / (uptime2 - uptime1)
        show global status where variable_name in ('com_insert' , 'com_delete' , 'com_update', 'uptime');
            已经执行的事务数
            事务数TC ≈'com_insert' + 'com_delete' + 'com_update' + 'com_select'
            TPS  ≈ (TC2 -TC1) / (uptime2 - uptime1)
        并发数  Threads_running
        当前连接数  Threads_connected
        最大连接数  max_connections（可用于计算连接数百分比）
        innodb缓冲池查询总数  innodb_buffer_pool_read_requests
        innodb从硬盘中读取数  innodb_buffer_pool_reads
        reads / read_requests = 穿透比例（比例越高，性能越差）
        慢查询总数  select COUNT(*) from information_schema.processlist;
        慢查询  Slow_queries
        show global status where Variable_name regexp 'Com_insert|Com_update|Com_delete|Com_select|Questions|Queries';
        超过最大连接数的报错次数    Connection_errors_max_connections
        失败的连接总数  Aborted_connects
        show global status where Variable_name regexp 'Connection_errors_max_connections|Aborted_connects';
        
        com_ 加合除以时间就是 QPS
        Queries 除以时间就是 TPS
        reads / read_requests 是 innodb 缓冲池穿透比例
        """
        curr_record = mysql_status
        last_record = self._last_record("mysql")
        if last_record:
            self.add_packet(
                "mysql.qps",
                curr_record["Com_insert"]
                + curr_record["Com_update"]
                + curr_record["Com_delete"]
                + curr_record["Com_select"]
                + curr_record["Com_update_multi"]
                + curr_record["Com_delete_multi"]
                - last_record["Com_insert"]
                - last_record["Com_update"]
                - last_record["Com_delete"]
                - last_record["Com_select"]
                - last_record["Com_update_multi"]
                - last_record["Com_delete_multi"],
            )
            self.add_packet(
                "mysql.tps", curr_record["Queries"] - last_record["Queries"]
            )
            self.add_packet(
                "mysql.questions", curr_record["Questions"] - last_record["Questions"]
            )
            self.add_packet(
                "mysql.com_insert",
                curr_record["Com_insert"] - last_record["Com_insert"],
            )
            self.add_packet(
                "mysql.com_update",
                curr_record["Com_update"] - last_record["Com_update"],
            )
            self.add_packet(
                "mysql.com_update_multi",
                curr_record["Com_update_multi"] - last_record["Com_update_multi"],
            )
            self.add_packet(
                "mysql.com_delete",
                curr_record["Com_delete"] - last_record["Com_delete"],
            )
            self.add_packet(
                "mysql.com_delete_multi",
                curr_record["Com_delete_multi"] - last_record["Com_delete_multi"],
            )
            self.add_packet(
                "mysql.com_select",
                curr_record["Com_select"] - last_record["Com_select"],
            )

            self.add_packet(
                "mysql.innodb_buffer_penetration_percent",
                (
                    (
                        +curr_record["Innodb_buffer_pool_read_ahead"]
                        + curr_record["Innodb_buffer_pool_reads"]
                        - last_record["Innodb_buffer_pool_read_ahead"]
                        - last_record["Innodb_buffer_pool_reads"]
                    )
                    / (
                        curr_record["Innodb_buffer_pool_read_requests"]
                        + curr_record["Innodb_buffer_pool_read_ahead"]
                        + curr_record["Innodb_buffer_pool_reads"]
                        - last_record["Innodb_buffer_pool_read_requests"]
                        - last_record["Innodb_buffer_pool_read_ahead"]
                        - last_record["Innodb_buffer_pool_reads"]
                    )
                    * 100
                    if (
                        curr_record["Innodb_buffer_pool_read_requests"]
                        + curr_record["Innodb_buffer_pool_read_ahead"]
                        + curr_record["Innodb_buffer_pool_reads"]
                        - last_record["Innodb_buffer_pool_read_requests"]
                        - last_record["Innodb_buffer_pool_read_ahead"]
                        - last_record["Innodb_buffer_pool_reads"]
                    )
                    else 0
                ),
            )

            self.add_packet(
                "mysql.innodb_buffer_pool_read_requests",
                curr_record["Innodb_buffer_pool_read_requests"]
                - last_record["Innodb_buffer_pool_read_requests"],
            )
            self.add_packet(
                "mysql.innodb_buffer_pool_read_ahead",
                curr_record["Innodb_buffer_pool_read_ahead"]
                - last_record["Innodb_buffer_pool_read_ahead"],
            )
            self.add_packet(
                "mysql.innodb_buffer_pool_reads",
                curr_record["Innodb_buffer_pool_reads"]
                - last_record["Innodb_buffer_pool_reads"],
            )
            self.add_packet("mysql.slow_queries", curr_record["Slow_queries"])
            self.add_packet(
                "mysql.aborted_connects",
                curr_record["Aborted_connects"] - last_record["Aborted_connects"],
            )
            # self.add_packet(
            #     "mysql.connection_errors_max_connections",
            #     curr_record["Connection_errors_max_connections"]
            #     - last_record["Connection_errors_max_connections"],
            # )
            self.add_packet(
                "mysql.threads_connected",
                curr_record["Threads_connected"] - last_record["Threads_connected"],
            )
            self.add_packet(
                "mysql.threads_running",
                curr_record["Threads_running"] - last_record["Threads_running"],
            )
            self.add_packet(
                "mysql.connections",
                curr_record["Connections"] - last_record["Connections"],
            )
            self.add_packet(
                "mysql.max_used_connections",
                curr_record["Max_used_connections"]
                - last_record["Max_used_connections"],
            )

    def redis_info(self):
        rd = redis.Redis(
            port=self.config["redis"]["port"],
            password=self.config["redis"]["password"],
        )
        key_index = [
            "used_memory",
            "instantaneous_ops_per_sec",
            "total_reads_processed",
            "total_writes_processed",
            "evicted_keys",
            "blocked_clients",
            "connected_clients",
            "rejected_connections",
            "keyspace_hits",
            "keyspace_misses",
            "rejected_connections",
            "total_net_input_bytes",
            "total_net_output_bytes",
            "total_connections_received",
            "total_commands_processed",
            "allocator_frag_ratio",
            "allocator_rss_ratio",
            "rss_overhead_ratio",
        ]
        rd_info = rd.info()
        rd_info = {k: rd_info[k] for k in key_index}
        self.db_insert("redis", rd_info)

        MB = 1024**2
        self.add_packet("redis.used_memory", rd_info["used_memory"] / MB)
        self.add_packet(
            "redis.instantaneous_ops_per_sec", rd_info["instantaneous_ops_per_sec"]
        )
        self.add_packet("redis.connected_clients", rd_info["connected_clients"])
        self.add_packet("redis.blocked_clients", rd_info["blocked_clients"])
        self.add_packet("redis.allocator_frag_ratio", rd_info["allocator_frag_ratio"])
        self.add_packet("redis.allocator_rss_ratio", rd_info["allocator_rss_ratio"])
        self.add_packet("redis.rss_overhead_ratio", rd_info["rss_overhead_ratio"])

        curr_record = rd_info
        last_record = self._last_record("redis")
        if last_record:
            self.add_packet(
                "redis.reads_processed",
                curr_record["total_reads_processed"]
                - last_record["total_reads_processed"],
            )
            self.add_packet(
                "redis.writes_processed",
                curr_record["total_writes_processed"]
                - last_record["total_writes_processed"],
            )
            self.add_packet(
                "redis.evicted_keys",
                curr_record["evicted_keys"] - last_record["evicted_keys"],
            )
            self.add_packet(
                "redis.rejected_connections",
                curr_record["rejected_connections"]
                - last_record["rejected_connections"],
            )
            self.add_packet(
                "redis.keyspace_hits",  # 启动以来命中已存在键的次数
                curr_record["keyspace_hits"] - last_record["keyspace_hits"],
            )
            self.add_packet(
                "redis.keyspace_misses",
                curr_record["keyspace_misses"] - last_record["keyspace_misses"],
            )
            self.add_packet(
                "redis.net_input_bytes",
                curr_record["total_net_input_bytes"]
                - last_record["total_net_input_bytes"],
            )
            self.add_packet(
                "redis.net_output_bytes",
                curr_record["total_net_output_bytes"]
                - last_record["total_net_output_bytes"],
            )
            self.add_packet(
                "redis.connections_received",  # 启动以来接受的总连接数
                curr_record["total_connections_received"]
                - last_record["total_connections_received"],
            )
            self.add_packet(
                "redis.commands_processed",
                curr_record["total_commands_processed"]
                - last_record["total_commands_processed"],
            )

    def mongodb_status(self):
        mongo = pymongo.MongoClient("127.0.0.1", 27017)
        # 获取 mongodb 的性能指标
        print(mongo.admin.command("serverStatus").get("metrics"))
        print(mongo.admin.command("dbStats"))
        # 获取内存使用情况
        print(mongo.admin.command("top"))

    def nginx_status(self):
        nginx_status_url = "http://status.nginx.local"
        nginx_status = requests.get(nginx_status_url).text
        # 将 nginx_status 解析成字典
        """
        Active connections: 当前活跃连接数，即当前与 Nginx 服务器建立的活跃连接数。
        Server accepts handled requests:
        accepts: Nginx 已经接受的连接总数。
        handled: Nginx 已经处理的连接总数。
        requests: Nginx 已经处理的请求数。
        Reading:
        当前正在读取客户端请求头的连接数。
        Writing:
        当前正在向客户端发送响应的连接数。
        Waiting:
        当前空闲的客户端连接数，等待请求。
        """
        nginx_status_lst = []
        for s in nginx_status.split():
            # 用正则表达式匹配，如果 s 是纯数字，则转换成 int 类型
            if re.match(r"\d+", s):
                nginx_status_lst.append(int(s))

        nginx_status = {
            "Active connections": nginx_status_lst[0],
            "Server accepts": nginx_status_lst[1],
            "Server handled": nginx_status_lst[2],
            "Server requests": nginx_status_lst[3],
            "Reading": nginx_status_lst[4],
            "Writing": nginx_status_lst[5],
            "Waiting": nginx_status_lst[6],
        }
        # print(nginx_status_dct)
        self.db_insert("nginx", nginx_status)

        curr_record = nginx_status
        last_record = self._last_record("nginx")
        if last_record:
            self.add_packet(
                "nginx.active_connections", curr_record["Active connections"]
            )
            self.add_packet(
                "nginx.server_accepts",
                curr_record["Server accepts"] - last_record["Server accepts"],
            )
            self.add_packet(
                "nginx.server_handled",
                curr_record["Server handled"] - last_record["Server handled"],
            )
            self.add_packet(
                "nginx.server_requests",
                curr_record["Server requests"] - last_record["Server requests"],
            )
            self.add_packet("nginx.reading", curr_record["Reading"])
            self.add_packet("nginx.writing", curr_record["Writing"])
            self.add_packet("nginx.waiting", curr_record["Waiting"])

    def phpfpm_status(self):
        phpfpm_status_url = "http://status.phpfpm.local?json"
        phpfpm_status = requests.get(phpfpm_status_url).json()
        """
        {'pool': 'www', 'process manager': 'dynamic', 'start time': 1712981210, 'start since': 27144, 'accepted conn': 515, 'listen queue': 0, 'max listen queue': 0, 'listen queue len': 0, 'idle processes': 3, 'active processes': 0, 'total processes': 3, 'max active processes': 3, 'max children reached': 0, 'slow requests': 0}
        """
        key_index = [
            "accepted conn",
            "listen queue",
            "max listen queue",
            "listen queue len",
            "active processes",
            "idle processes",
            "total processes",
            "max active processes",  # 从fpm启动以来，活动进程的最大个数，如果这个值小于当前的max_children，可以调小此值
            "max children reached",  # 当pm尝试启动更多的进程，却因为max_children的限制，没有启动更多进程的次数。如果这个值非0，那么可以适当增加fpm的进程数
            "slow requests",  # 慢请求的次数，一般如果这个值未非0，那么可能会有慢的php进程，一般一个不好的mysql查询是最大的祸首。
        ]
        phpfpm_status = {k: phpfpm_status[k] for k in key_index}
        self.db_insert("phpfpm", phpfpm_status)

        curr_record = phpfpm_status
        last_record = self._last_record("phpfpm")
        if last_record:
            self.add_packet(
                "phpfpm.accepted_conn",
                curr_record["accepted conn"] - last_record["accepted conn"],
            )
            self.add_packet("phpfpm.listen_queue", curr_record["listen queue"])
            self.add_packet("phpfpm.max_listen_queue", curr_record["max listen queue"])
            self.add_packet("phpfpm.listen_queue_len", curr_record["listen queue len"])
            self.add_packet("phpfpm.active_processes", curr_record["active processes"])
            self.add_packet("phpfpm.idle_processes", curr_record["idle processes"])
            self.add_packet("phpfpm.total_processes", curr_record["total processes"])
            self.add_packet(
                "phpfpm.max_active_processes", curr_record["max active processes"]
            )
            self.add_packet(
                "phpfpm.max_children_reached", curr_record["max children reached"]
            )
            self.add_packet("phpfpm.slow_requests", curr_record["slow requests"])

    def docker_stats(self):
        # 获取 docker 的状态
        # docker stats（耗时2.54秒，且IONet为0B，时间段统计起来比较困难）
        t = time.time()
        print(
            subprocess.run(
                'docker stats --no-stream --format "{{ json . }}"',
                shell=True,
                stdout=subprocess.PIPE,
            ).stdout.decode()
        )
        print("耗时：", time.time() - t)

    def send_packets(self):

        # 查询本地数据库中状态为 wait 的最小、最大 id
        sql = "select min(id), max(id) from packet where send_time = 0"
        min_id, max_id = self.db_exec(sql)[0]

        # 无数据则会返回 [(None, None)]
        if not min_id or not max_id:
            logging.info("无待发送数据")
            return

        # print(min_id, max_id)
        # 如果数据量大于 128 条，则只发送 128 条
        if max_id - min_id >= 128:
            max_id = min_id + 127
        # print(min_id, max_id)

        # 从本地数据库中取出相应数据
        packets = self.db_exec(
            f"select time, uri, value from packet where id between {min_id} and {max_id}"
        )
        # 转为列表字典的形式
        packets_to_send = [
            {
                "created_at": p[0],
                "DataTypeID": self._uri_to_id(p[1]),
                "value": p[2],
                "MachineID": self._machine_id,
            }
            for p in packets
        ]
        # print(json.dumps(packets_to_send, indent=4))

        # 将这些数据的 send_time 标记为当前时间戳
        self.db_exec(
            f"update packet set send_time = {int(time.time())} where id between {min_id} and {max_id}"
        )

        logging.info(
            f"发送数据：{min_id} - {max_id}（共 {max_id-min_id+1} 个 packets）"
        )
        # 向服务端发送数据
        r = requests.post(
            self.base_url + self.uri_dct["upload"],
            headers={"x-token": self.token()},
            json={"data": packets_to_send},
        )

        if r.status_code == 200:
            # 删除本地数据库中的相应数据
            self.db_exec(f"delete from packet where id between {min_id} and {max_id}")
            logging.info(f"发送成功：{min_id} - {max_id}（packets 已删除）")
        else:
            logging.warning(r.text)

    def resend_packets(self):
        # 获取当前的时间戳
        now = int(time.time())

        # 从本地数据库中取出 send_time 与当前时间戳相差超过 10 秒的数据（限制 128 个）
        # 由于网络连接存在 timeout，而服务器端与云数据库通信存在延迟
        # 故单次发送的数据需要限制在一定数量，否则网络连接会被服务器端断开
        packets = self.db_exec(
            f"select id, time, uri, value, send_time from packet where send_time != 0 and {now} - send_time > 10 limit 128"
        )

        # 确定是否存在需要补发的数据
        if not packets:
            logging.info("无待补发数据")
            return

        # 转为列表字典的形式
        packets_to_send = [
            {
                "created_at": p[1],
                "DataTypeID": self._uri_to_id(p[2]),
                "value": p[3],
                "MachineID": self._machine_id,
            }
            for p in packets
        ]
        # print(json.dumps(packets_to_send, indent=4))

        # 将这些数据的 send_time 标记为当前时间戳
        self.db_exec(
            f"update packet set send_time = {now} where id in ({','.join([str(p[0]) for p in packets])})"
        )

        logging.info(f"补发数据：{len(packets)} 个 packets")
        # 向服务端发送数据
        r = requests.post(
            self.base_url + self.uri_dct["upload"],
            headers={"x-token": self.token()},
            json={"data": packets_to_send},
        )

        if r.status_code == 200:
            # 删除本地数据库中的相应数据
            self.db_exec(
                f"delete from packet where id in ({','.join([str(p[0]) for p in packets])})"
            )
            logging.info(f"补发成功：{len(packets)} 个 packets（packets 已删除）")
        else:
            logging.warning(r.text)

    def update_service_status(self):
        running_services = set()
        # 根据进程的名称（name）来判断是否有服务在运行
        for process in psutil.process_iter():
            for service_name in self.detect_services:
                if service_name in process.name():
                    running_services.add(self.detect_services[service_name])

        # 若运行服务没有发生变化，则直接返回
        if running_services == self.running_services:
            return

        # 服务端的特别需求
        service_status = {}
        for service in running_services:
            # 读入配置获取是否启用
            if self.config.get(service):
                service_status[service] = int(self.config[service].get("enable", False))
            else:
                service_status[service] = 0

        logging.info(f"更新正在运行的服务：{service_status}")

        # 向服务器端发送正在运行的服务
        r = requests.post(
            self.base_url + self.uri_dct["service"]["update"],
            headers={"x-token": self.token()},
            json={
                "MachineID": self._machine_id,
                "Services": json.dumps(service_status),
            },
        )
        logging.info(f"更新服务：{service_status}")

        if r.json()["code"] == 7 and r.json()["msg"] == "该数据不存在":
            # 为该机器创建 services 项
            r = requests.post(
                self.base_url + self.uri_dct["service"]["create"],
                headers={"x-token": self.token()},
                json={
                    "MachineID": self._machine_id,
                    "Services": json.dumps(list(running_services)),
                },
            )
            logging.info(f"创建服务：{running_services}")

        """
        POST /machine/createMachineService
             /machine/updateMachineService
        Header:
            x-token: token
        Body:
            {
                "machineID": 3,
                "services": ["nginx", "phpfpm", "mysql", "redis"]
            }
        """

        # # 查看进程名称列表（只取 name、exe、cmline 字段）
        # ps = {}
        # for process in psutil.process_iter():
        #     info = process.as_dict(attrs=["pid", "name", "exe", "cmdline"])
        #     pid = info["pid"]
        #     name = info["name"]
        #     exe = info["exe"]
        #     cmdline = info["cmdline"]
        #     cmdline = " ".join(" ".join(cmdline).strip().split())
        #     for string in [name, exe, cmdline]:
        #         for service in all_services:
        #             if service in string:
        #                 ps[pid] = (name, exe, cmdline)
        #                 break
        # # 写入到 json 文件
        # with open(self.cwd + "/process.json", "w", encoding="utf-8") as f:
        #     json.dump(ps, f, indent=4)

    def modify_hosts(self):
        hosts_path = "/etc/hosts"
        hosts = open(hosts_path, "r", encoding="utf-8").read()
        # 在 /etc/hosts 中检查是否有 # monit 标记，若不含该标记
        if not re.search(r"# monit", hosts):
            hosts += """
# monit - DO NOT EDIT
127.0.0.1 status.nginx.local
127.0.0.1 status.phpfpm.local
# monit - DO NOT EDIT"""
            open(hosts_path, "w", encoding="utf-8").write(hosts)

    def is_conf_modified(self, conf_path):
        # 打开文件，读取文件内容，搜索是否有标记 # monit
        nginx_conf = open(conf_path, "r", encoding="utf-8").read()
        res = re.search(r"[#;] monit", nginx_conf)
        if res:
            logging.info(f"配置文件 {conf_path} 已开启监控模块")
        else:
            logging.info(f"配置文件 {conf_path} 尚未开启监控模块，正在修改中 ...")
        return res

    def modify_nginx_conf(self):
        nginx_conf_path = self.config["nginx"]["path"]
        nginx_conf = open(nginx_conf_path, "r", encoding="utf-8").read()

        # 开始添加的标记
        nginx_conf_reg = r"\n\s*http\s*{"
        nginx_conf_reg_sub = """
http {
# monit - DO NOT EDIT
server {
    server_name status.nginx.local;
    location / {
        stub_status;
    }
}
server {
    server_name status.phpfpm.local;
    location / {
        fastcgi_pass unix:/run/php/fpm-status.sock;
        fastcgi_split_path_info ^(.+?\.php)(/.*)$;
        fastcgi_param PATH_INFO $fastcgi_path_info;
        fastcgi_param SCRIPT_NAME $fastcgi_script_name;
        fastcgi_param SCRIPT_FILENAME $document_root$fastcgi_script_name;
        fastcgi_param REQUEST_URI $request_uri;
        fastcgi_param REQUEST_METHOD $request_method;
        fastcgi_param REQUEST_SCHEME $scheme;
        fastcgi_param QUERY_STRING $query_string;
        fastcgi_param SERVER_ADDR $server_addr;
        fastcgi_param SERVER_PORT $server_port;
        fastcgi_param SERVER_NAME $server_name;
    }
}
# monit - DO NOT EDIT
"""
        nginx_conf = re.sub(nginx_conf_reg, nginx_conf_reg_sub, nginx_conf)
        # 备份 nginx 配置文件（获取今天的日期 2024-04-13）
        today_str = datetime.datetime.now().strftime("%Y-%m-%d")
        os.system(f"cp -f {nginx_conf_path} {nginx_conf_path}-{today_str}.monit.bak")
        open(nginx_conf_path, "w", encoding="utf-8").write(nginx_conf)
        logging.info(f"配置文件 {nginx_conf_path} 现已开启监控模块")
        # 重启 nginx
        os.system("nginx -s reload")
        logging.info("nginx 已重载配置文件")

    def modify_phpfpm_conf(self):
        # 直接在文件末尾加上对应监控模块的配置
        phpfpm_conf_path = self.config["php-fpm"]["path"]
        # 读取 phpfpm_conf 文件
        phpfpm_conf = open(phpfpm_conf_path, "r", encoding="utf-8").read()
        phpfpm_conf += """
; monit - DO NOT EDIT
pm.status_path = /
pm.status_listen = /run/php/fpm-status.sock
; monit - DO NOT EDIT
"""
        open(phpfpm_conf_path, "w", encoding="utf-8").write(phpfpm_conf)
        logging.info(f"配置文件 {phpfpm_conf_path} 现已开启监控模块")
        # 重启 php-fpm，在所有的系统服务中搜索含有 php 和 fpm 的服务，重启该服务
        # 以非交互式方式列出所有系统服务
        system_services = subprocess.run(
            "systemctl list-units --type=service --no-pager",
            shell=True,
            stdout=subprocess.PIPE,
        ).stdout.decode()
        # 获取所有的 php-fpm 服务
        phpfpm_services = re.findall(r"php.*fpm.*service", system_services)
        # 重启所有的 php-fpm 服务
        for s in phpfpm_services:
            os.system(f"systemctl restart {s}")
        logging.info("php-fpm 已重启")

        # # 如果找到 www.conf 文件，则修改 pm.status_path 和 pm.status_listen
        # if www_conf_path:
        #     # 读取  www.conf 文件
        #     www_conf = open(www_conf_path, "r", encoding="utf-8").read()
        #     # 如果 www.conf 文件中没有标记符号，则添加
        #     if not re.search(r"; monit", www_conf):
        #         www_conf = re.sub(r";\s*pm.status_path", "pm.status_path", www_conf)
        #         www_conf = re.sub(
        #             r"pm.status_path\s*=\s*[^\n]*",
        #             "; monit - DO NOT EDIT\npm.status_path = /",
        #             www_conf,
        #         )
        #         www_conf = re.sub(r";\s*pm.status_listen", "pm.status_listen", www_conf)
        #         www_conf = re.sub(
        #             r"pm.status_listen\s*=\s*[^\n]*",
        #             "; monit - NO NOT EDIT\npm.status_listen = /run/php/fpm-status.sock",
        #             www_conf,
        #         )
        #         open(www_conf_path, "w", encoding="utf-8").write(www_conf) if 1 else 1

    def mysql_exec(
        self,
        sql,
        args=None,
        database="mysql",
        host="127.0.0.1",
        port=3306,
        user="root",
        passwd=None,
    ):
        # 如果未指定 host、port、user、passwd，则使用默认值
        # if not host:
        #     if not self.config.get("mysql"):
        #         host = self.config.get("mysql").get("host", "127.0.0.1")
        #         logging.warning("[mysql] 未指定 MySQL 主机，默认使用")

        db = pymysql.connect(
            host=host,
            port=port,
            user=user,
            passwd=passwd,
            database=database,
            cursorclass=pymysql.cursors.DictCursor,
        )

        cursor = db.cursor()

        if type(sql) == type(""):
            cursor.execute(sql, args)
        elif type(sql) in [type((0,)), type([])]:
            num = len(sql)
            if args:  # 如果 args 是一个列表的话
                for i in range(num):
                    cursor.execute(sql[i], args[i])
            else:  # 如果 args 为 None
                for i in range(num):
                    cursor.execute(sql[i])

        try:
            dataDct = cursor.fetchall()
        except:
            dataDct = None

        db.commit()
        cursor.close()
        db.close()
        return dataDct

    def db_exec(self, sql, args=()):
        db_conn = sqlite3.connect(self.db_path)
        cursor = db_conn.execute(sql, args)
        res = cursor.fetchall()
        db_conn.commit()
        cursor.close()
        db_conn.close()
        return res

    def _db_init(self):
        self.db_path = f"{self.cwd}/monit.db"
        table_sql = [
            # 使用本地时间，而不是 UTC 时间
            """CREATE TABLE IF NOT EXISTS data (
                    time DATETIME DEFAULT (datetime('now', 'localtime')),
                    service TEXT,
                    data TEXT
            );""",
            """CREATE TABLE IF NOT EXISTS packet (
                    id INTEGER PRIMARY KEY AUTOINCREMENT,
                    time DATETIME DEFAULT (datetime('now', 'localtime')),
                    uri TEXT,
                    value FLOAT,
                    send_time INTEGER DEFAULT 0
            );""",
            """CREATE TABLE IF NOT EXISTS dict (
                    key TEXT PRIMARY KEY,
                    value TEXT
            );""",
        ]
        for sql in table_sql:
            self.db_exec(sql)

    def db_insert(self, service: str, data: dict):
        sql = "INSERT INTO data (service, data) VALUES (?, ?);"
        # # 获取当前时间（确保仅精确到秒，而不会保留毫秒纳秒的数值）
        # time_str = datetime.datetime.now().strftime("%Y-%m-%d %H:%M:%S")
        # 将 data 的字典转换成字符串
        data_str = json.dumps(data, ensure_ascii=False)
        self.db_exec(sql, (service, data_str))

    def add_packet(self, uri: str, value: float):
        sql = "INSERT INTO packet (uri, value) VALUES (?, ?);"
        self.db_exec(sql, (uri, value))

    def db_dct(self, key: str, value=None):
        if not value:
            sql = "SELECT value FROM dict WHERE key = ?;"
            data = self.db_exec(sql, (key,))
            res = data[0][0] if data else None
            return json.loads(res) if res else None

        if value:
            # 将 value 转换成 JSON 字符串
            value = json.dumps(value, ensure_ascii=False, cls=self._GeneralEncoder)
            # 检查 key 是否存在，若存在则更新，否则插入
            sql = "SELECT value FROM dict WHERE key = ?;"
            data = self.db_exec(sql, (key,))
            if data:
                sql = "UPDATE dict SET value = ? WHERE key = ?;"
                self.db_exec(sql, (value, key))
            else:
                sql = "INSERT INTO dict (key, value) VALUES (?, ?);"
                self.db_exec(sql, (key, value))

    class _GeneralEncoder(json.JSONEncoder):
        # 一个可以序列化时间的 JSONEncoder
        def default(self, obj):
            if isinstance(obj, datetime.datetime):
                return obj.strftime("%Y-%m-%d %H:%M:%S")
            if isinstance(obj, datetime.date):
                return obj.strftime("%Y-%m-%d")
            return json.JSONEncoder.default(self, obj)

    def _db_rotate(self):
        # 清空设定时间之前的数据（sqlite3 无法很好地处理时区差异，还是使用 Python 的）
        # now = datetime.datetime.now()
        # interval = datetime.timedelta(minutes=3)
        # expire = (now - interval).strftime("%Y-%m-%d %H:%M:%S")
        self.db_exec(
            f"DELETE FROM data WHERE time < datetime('now', 'localtime', '-60 seconds');"
        )
        self.db_exec(
            f"DELETE FROM packet WHERE time < datetime('now', 'localtime', '-15 minutes');"
        )

    def align_datatype(self):
        # 读入本地的 datatype 文件
        datatype = [
            line.strip().split(":")
            for line in open(
                f"{self.cwd}/datatype.txt", "r", encoding="utf-8"
            ).readlines()
            if line.strip()
        ]
        # 第一项是 uri（用作 name，并可计算 id），第二项是单位，第三项是描述
        # cpu.cpu_percent:%:CPU 利用率
        datatype = {
            self._uri_to_id(d[0]): {
                "name": d[0],
                "description": d[2],
                "units": d[1],
            }
            for d in datatype
        }

        data_type_db = {
            "host": "gz-cynosdbmysql-grp-5r17wgzb.sql.tencentcdb.com",
            "port": 26316,
            "user": "root",
            "passwd": "aA877783019",
            "database": "gva",
        }

        # 清空 MySQL 数据库
        self.mysql_exec("TRUNCATE TABLE data_type;", **data_type_db)

        # 将 datatype 写入 MySQL 数据库
        for k, v in datatype.items():
            self.mysql_exec(
                "INSERT INTO data_type (id, name, description, units, created_at, updated_at) VALUES (%s, %s, %s, %s, %s, %s);",
                (
                    k,
                    v["name"],
                    v["description"],
                    v["units"],
                    datetime.datetime.now(),
                    datetime.datetime.now(),
                ),
                **data_type_db,
            )

    def _last_record(self, service: str):
        sql = "SELECT data FROM data WHERE service = ? AND time = datetime('now', 'localtime', '-1 seconds');"
        res = self.db_exec(sql, (service,))
        return json.loads(res[0][0]) if res else None

    def _uri_to_id(self, uri: str) -> int:
        # 只取前 SHA-256 的前 32 位，因为 Go 那边使用的是 int
        return int(hashlib.sha256(uri.encode()).hexdigest()[:8], 16)

    def _wait(self, start=0.0, interval=1):
        while True:
            time_precision = 0.0003
            time_segment = [0.0, 0.6, 0.9, 0.97, 0.995, 0.999, 0.9997, 1.0]
            wait_monit_cnt = 0
            for i, t in enumerate(time_segment[:-1]):
                while (
                    t
                    <= ((time.time() - start) % interval) / interval
                    < time_segment[i + 1]
                ):
                    if time_segment[i + 1] != 1:
                        time.sleep((1 - time_segment[i + 1]) * interval)
                    else:
                        time.sleep(time_precision * interval)
                    wait_monit_cnt += 1
            # print(time.time(), wait_monit_cnt)
            return

    def _timing(self, func, start=0.0, interval=1, exec_now=False):
        def run():
            try:
                threading.Thread(target=func).start()
            except Exception as e:
                logging.error(func.__name__, e)

        if exec_now:
            run()
        while True:
            self._wait(start, interval)
            run()

    def start_monit(self):
        """
        这个部分的代码一定要相当健壮
        """

        threading.Thread(target=self._timing, args=(self.cpu,)).start()
        threading.Thread(target=self._timing, args=(self.memory,)).start()
        threading.Thread(target=self._timing, args=(self.disk,)).start()
        threading.Thread(target=self._timing, args=(self.disk_io,)).start()
        threading.Thread(target=self._timing, args=(self.net_io,)).start()

        # MySQL
        try:
            if self.config.get("mysql"):
                if self.config["mysql"].get("enable"):
                    while not self.mysql_exec(
                        "show databases;",
                        port=self.config["mysql"].get("port"),
                        passwd=self.config["mysql"].get("password"),
                    ):
                        logging.warning("MySQL 连接失败，请检查端口、密码")
                        time.sleep(5)
                    threading.Thread(
                        target=self._timing, args=(self.mysql_status,)
                    ).start()
        except Exception as e:
            logging.error("启动 MySQL 监控失败", e)

        # redis
        try:
            if self.config.get("redis"):
                if self.config["redis"].get("enable"):
                    while not redis.Redis(
                        port=self.config["redis"].get("port"),
                        password=self.config["redis"].get("password"),
                    ):
                        logging.warning("Redis 连接失败，请检查端口、密码")
                        time.sleep(5)
                    threading.Thread(
                        target=self._timing, args=(self.redis_info,)
                    ).start()
        except Exception as e:
            logging.error("启动 Redis 监控失败", e)

        # nginx
        try:
            if self.config.get("nginx"):
                if self.config["nginx"].get("enable"):
                    nginx_conf_path = self.config["nginx"].get("path")
                    if not nginx_conf_path:
                        raise Exception("Nginx 配置文件路径未指定")
                    if not os.path.exists(nginx_conf_path):
                        raise Exception("Nginx 配置文件不存在")
                    # 如果 nginx_conf_path 是目录，则更新为 nginx_conf_path/nginx.conf
                    if os.path.isdir(nginx_conf_path):
                        self.config["nginx"]["path"] = os.path.join(
                            nginx_conf_path, "nginx.conf"
                        )

                    self.save_config()

                    # 检查 nginx 配置文件是否被修改
                    if not self.is_conf_modified(self.config["nginx"]["path"]):
                        self.modify_nginx_conf()
                        self.modify_hosts()

                    # 检查连通性
                    while requests.get("http://status.nginx.local").status_code != 200:
                        logging.warning("Nginx 监控页面无法访问，请检查 Nginx 配置")
                        time.sleep(5)

                    threading.Thread(
                        target=self._timing, args=(self.nginx_status,)
                    ).start()
        except Exception as e:
            logging.error("启动 Nginx 监控失败", e)

        # php-fpm
        try:
            if self.config.get("php-fpm"):
                if self.config["php-fpm"].get("enable"):
                    phpfpm_path = self.config["php-fpm"].get("path")
                    if not phpfpm_path:
                        raise Exception("PHP-FPM 路径未指定")
                    if not os.path.exists(phpfpm_path):
                        raise Exception("PHP-FPM 路径不存在")
                    # 如果 phpfpm_path 是目录，则遍历其目录下的文件夹，找到 www.conf
                    # 如果找不到 www.conf，就找 php-fpm.conf
                    if os.path.isdir(phpfpm_path):
                        supplement_path = [
                            "pool.d/www.conf",
                            "php-fpm.conf",
                            "fpm/pool.d/www.conf",
                            "fpm/php-fpm.conf",
                        ]
                        conf_exist = False
                        for sp in supplement_path:
                            phpfpm_full_path = f"{phpfpm_path}/{sp}"
                            if os.path.exists(phpfpm_full_path):
                                conf_exist = True
                                break
                        if not conf_exist:
                            raise Exception("PHP-FPM 配置文件不存在")
                        self.config["php-fpm"]["path"] = phpfpm_full_path

                    self.save_config()

                    # 检测 php-fpm 配置文件是否被修改
                    if not self.is_conf_modified(self.config["php-fpm"]["path"]):
                        self.modify_phpfpm_conf()
                        self.modify_hosts()

                    # 检查连通性
                    while requests.get("http://status.phpfpm.local").status_code != 200:
                        logging.warning("PHP-FPM 监控页面无法访问，请检查 PHP-FPM 配置")
                        time.sleep(5)

                    threading.Thread(
                        target=self._timing, args=(self.phpfpm_status,)
                    ).start()
        except Exception as e:
            logging.error("启动 PHP-FPM 监控失败", e)

        threading.Thread(
            target=self._timing, args=(self._db_rotate, 0.3, 5, True)
        ).start()
        threading.Thread(
            target=self._timing, args=(self.send_packets, 0.25, 0.5)
        ).start()
        threading.Thread(
            target=self._timing, args=(self.resend_packets, 0.4, 3, True)
        ).start()
        threading.Thread(
            target=self._timing,
            args=(self.update_service_status, 0.5, 5, True),
            daemon=True,
        ).start()


# ==================== 处理入参 ====================

if args.subcommand == "init":
    agent = Agent(args.machine_id, args.password)
    logging.info("初始化 Agent 完毕，开始监控 ...")
    subprocess.run(
        "nohup python /usr/local/monit/agent.py monit >/dev/null 2>&1 &", shell=True
    )

if args.subcommand == "configure":
    logging.info("更新 Agent 配置中 ...")
    agent = Agent()
    # 根据 agent.detect_services 和 vars(args) 生成 config.yml
    config = {}
    for service in agent.detect_services.values():
        for opt, val in vars(args).items():
            if opt.rsplit("_", 1)[0].replace("_", "-") == service:
                if service not in config:
                    config[service] = {}
                config[service][opt.rsplit("_", 1)[1]] = val
    # print(config)
    # 获取本文件的位置
    cwd = os.path.dirname(os.path.abspath(__file__))
    # 将所有相关的参数写入到 config.yml 中
    with open(f"{cwd}/config.yml", "w", encoding="utf-8") as f:
        yaml.dump(config, f)

    # 重载 Agent
    logging.info("Agent 配置已更新，准备重载 ...")
    subprocess.run(
        "nohup python /usr/local/monit/agent.py monit >/dev/null 2>&1 &", shell=True
    )

if args.subcommand == "monit":
    # 检测系统中的进程是否有 agent.py --monit 在运行，若有则中止
    # 中止进程后，文件锁 /tmp/agent.lock 会被释放
    for process in psutil.process_iter():
        # 如果该进程的 PID 不同于当前进程的 PID
        if process.pid != os.getpid():
            cmd = " ".join(" ".join(process.cmdline()).strip().split())
            if "agent.py monit" in cmd:
                if args.cron:
                    sys.exit(0)
                logging.info(f"已有 Agent 进程运行中：{cmd}")
                process.terminate()
                logging.info(f"重载 Agent 进程中 ...")

    # # 创建文件锁
    # lock = open("/tmp/agent.lock", "w")
    # try:
    #     # 非阻塞独占式加锁
    #     fcntl.flock(lock, fcntl.LOCK_EX | fcntl.LOCK_NB)

    agent = Agent()
    agent.start_monit()
    # except:
    #     logging.info("Agent 进程已在运行中")
    #     sys.exit(0)

if args.subcommand == "stop":
    # 移除保活 cron
    os.system("rm -f /etc/cron.d/monit")
    # 检测系统中的进程是否有 agent.py --monit 在运行，若有则中止
    for process in psutil.process_iter():
        # 如果该进程的 PID 不同于当前进程的 PID
        if process.pid != os.getpid():
            cmd = " ".join(" ".join(process.cmdline()).strip().split())
            if "agent.py monit" in cmd:
                logging.info(f"中止 Agent 进程：{cmd}")
                process.terminate()
                logging.info(f"Agent 进程已中止")
                sys.exit(0)

if args.subcommand == "uninstall":
    try:
        pattern = r"\n[#;] ?monit.*?[#;] ?monit[^\n]*"

        # 读入配置文件
        with open(f"/usr/local/monit/config.yml", "r", encoding="utf-8") as f:
            config = yaml.safe_load(f)

        # 删除 hosts 文件中的 monit 部分
        hosts_path = "/etc/hosts"
        hosts = open(hosts_path, "r", encoding="utf-8").read()
        print(re.search(pattern, hosts, flags=re.DOTALL))
        hosts = re.sub(pattern, "", hosts, flags=re.DOTALL)
        open(hosts_path, "w", encoding="utf-8").write(hosts)
        print("hosts 文件已清除 monit 痕迹")

        # 删除 nginx 配置文件中的 monit 部分
        try:
            nginx_conf_path = config["nginx"]["path"]
        except:
            nginx_conf_path = "/etc/nginx/nginx.conf"
        nginx_conf = open(nginx_conf_path, "r", encoding="utf-8").read()
        print(re.search(pattern, nginx_conf, flags=re.DOTALL))
        nginx_conf = re.sub(pattern, "", nginx_conf, flags=re.DOTALL)
        open(nginx_conf_path, "w", encoding="utf-8").write(nginx_conf)
        print("nginx 配置文件已清除 monit 痕迹")

        # 删除 php-fpm 配置文件中的 monit 部分
        phpfpm_conf_path = config["php-fpm"]["path"]
        phpfpm_conf = open(phpfpm_conf_path, "r", encoding="utf-8").read()
        print(re.search(pattern, phpfpm_conf, flags=re.DOTALL))
        phpfpm_conf = re.sub(pattern, "", phpfpm_conf, flags=re.DOTALL)
        open(phpfpm_conf_path, "w", encoding="utf-8").write(phpfpm_conf)
        print("php-fpm 配置文件已清除 monit 痕迹")

    except Exception as e:
        pass
        print(e)

    # 删除保活 cron
    print("正在移除保活措施 ...")
    os.system("rm -f /etc/cron.d/monit")

    # 终止所有 Agent 守护进程
    print("正在终止 Agent 守护进程 ...")
    for process in psutil.process_iter():
        # 如果该进程的 PID 不同于当前进程的 PID
        if process.pid != os.getpid():
            cmd = " ".join(" ".join(process.cmdline()).strip().split())
            if "agent.py monit" in cmd:
                process.terminate()
                print(f"Agent 守护进程 {cmd} 已终止")

    # 彻底移除 monit 安装目录
    print("正在彻底删除 monit Agent 安装目录 ...")
    os.system("rm -rf /usr/local/monit")
    print("monit Agent 已彻底删除，感谢您的使用！")

    # 停止 Agent 进程
    for process in psutil.process_iter():
        # 如果该进程的 PID 不同于当前进程的 PID
        if process.pid != os.getpid():
            cmd = " ".join(" ".join(process.cmdline()).strip().split())
            if "agent.py monit" in cmd:
                logging.info(f"中止 Agent 进程：{cmd}")
                process.terminate()
                logging.info(f"Agent 进程已中止")
                sys.exit(0)

if not args.subcommand:
    agent = Agent()
    parser.print_help()
