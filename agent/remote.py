#!python

import subprocess, argparse, importlib, sys, os, requests


def install_module(module_name: str):
    try:
        # 尝试导入模块
        importlib.import_module(module_name)
    except ImportError:
        # 如果导入失败则安装模块
        source = "https://pypi.tuna.tsinghua.edu.cn/simple"
        cmd = f"pip install {module_name} -i {source} -U"
        subprocess.check_call(cmd.split())


parser = argparse.ArgumentParser()
subparsers = parser.add_subparsers(dest="subcommand")

# 常规参数解析
parser.add_argument("--host", type=str, required=True, help="远程主机IP地址")
parser.add_argument("--port", type=int, default=22, help="远程主机端口 (default: 22)")
parser.add_argument(
    "--username", type=str, default="root", help="远程主机用户名 (default: root)"
)
parser.add_argument("--password", type=str, help="远程主机密码 (已安装 Agent 可留空)")

# 安装参数解析
install_parser = subparsers.add_parser("install", help="在远程主机上安装 Agent")
install_parser.add_argument("--machine-id", type=int, required=True, help="machineID")

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

# 运行命令参数解析
run_parser = subparsers.add_parser("run", help="在远程主机上执行命令")
run_parser.add_argument("--command", type=str, required=True, help="执行的命令")

# 停止监控命令参数解析
stop_parser = subparsers.add_parser("stop", help="在远程主机上停止 Agent 监控")

# 卸载命令参数解析
uninstall_parser = subparsers.add_parser("uninstall", help="在远程主机上卸载 Agent")

args = parser.parse_args()
print(args)

install_module("paramiko")
import paramiko

ssh_client = paramiko.SSHClient()
ssh_client.set_missing_host_key_policy(paramiko.AutoAddPolicy())
ssh_client.connect(
    hostname=args.host,
    port=args.port,
    username=args.username,
    password=args.password,
)


def remote_exec(command: str, silent=False):
    stdin, stdout, stderr = ssh_client.exec_command(command)
    out = stdout.read().decode()
    err = stderr.read().decode()
    if not silent:
        if out:
            print(out)
        if err:
            print(err)
    return out, err


if args.subcommand == "install":

    print("正在安装 Agent ...")

    # 获取运行本文件的用户 home 目录
    home = subprocess.check_output("echo $HOME", shell=True).decode().strip()
    # 检查本机是否存在 ~/.ssh/id_rsa.pub 文件
    if not os.path.exists(f"{home}/.ssh/id_rsa.pub"):
        print("本机不存在 ~/.ssh/id_rsa.pub 文件，先生成密钥对")
        # 生成密钥对并设置权限
        subprocess.check_call(
            f"mkdir {home}/.ssh && chmod 700 {home}/.ssh && ssh-keygen -t rsa -f {home}/.ssh/id_rsa -N '' && chmod 600 {home}/.ssh/id_rsa",
            shell=True,
        )
        print("密钥对已生成")

    # 获取本机公钥
    pubkey = open(f"{home}/.ssh/id_rsa.pub").read().strip().split(" ")[1]

    # 读取远程主机的 authorized_keys 文件，若尚不本机公钥则添加
    authorized_keys, _ = remote_exec("cat ~/.ssh/authorized_keys", silent=True)

    if pubkey not in authorized_keys:
        print("本机公钥未添加到远程主机，正尝试添加")
        # 安装本机公钥到远程主机
        remote_exec(
            f'mkdir -p ~/.ssh && chmod 700 ~/.ssh && echo "{pubkey}" >> ~/.ssh/authorized_keys && chmod 600 ~/.ssh/authorized_keys'
        )
        print("本机公钥已成功添加到远程主机")

    # 将 agent.py 下载到远程主机的 /usr/local/monit 目录中
    print("从 https://file.jiangyj.tech/proj/monit/agent.py 下载 agent.py ...")
    remote_exec(
        "mkdir -p /usr/local/monit && curl -s https://file.jiangyj.tech/proj/monit/agent.py -o /usr/local/monit/agent.py && chmod +x /usr/local/monit/agent.py"
    )
    print("agent.py 已安装至 /usr/local/monit")

    # 安装 agent.py 所需依赖
    print("初始化 Agent ...\n")
    # os, sys, psutil, time, json, requests, datetime, subprocess, re, sqlite3, redis, pymongo, pymysql, argparse, hashlib, threading, yaml, logging
    agent_modules = {"psutil", "requests", "redis", "pymongo", "PyMySQL", "PyYAML"}
    # 检查 Agent 所需依赖 module 是否已安装，若未安装则安装
    out, _ = remote_exec("pip list", silent=True)
    installed_modules = set(out.strip().split()[2:])
    # print("已安装模块:", installed_modules)
    if not agent_modules.issubset(installed_modules):
        print("正在安装 Agent 所需依赖 ...")
        remote_exec(
            f"pip install {' '.join(agent_modules - installed_modules)} -i https://pypi.tuna.tsinghua.edu.cn/simple -U"
        )
        print("Agent 所需依赖已安装完毕")

    # 获取本机的公网地址
    ip_api = "https://searchplugin.csdn.net/api/v1/ip/get"
    server_ip = requests.get(ip_api).json()["data"]["ip"]
    print("服务端 IP 地址为:", server_ip)
    print(f"后续监控数据将发送至 {server_ip}:8888")
    print("请注意开启服务端的公网 8888 端口，否则 Agent 发送数据将失败")

    # 将 machine_id 传给 agent.py 存储
    print("machine_id:", args.machine_id)
    print("password:", args.password)
    remote_exec(
        f"python /usr/local/monit/agent.py init --server-ip {server_ip} --machine-id {args.machine_id} --password {args.password}"
    )

if args.subcommand == "configure":
    print("正在配置 Agent ...\n")
    remote_exec(
        f"python /usr/local/monit/agent.py configure --mysql-enable {args.mysql_enable} --mysql-port {args.mysql_port} --mysql-user {args.mysql_user} --mysql-password {args.mysql_password} --redis-enable {args.redis_enable} --redis-port {args.redis_port} --redis-password {args.redis_password} --nginx-enable {args.nginx_enable} --nginx-path {args.nginx_path} --php-fpm-enable {args.php_fpm_enable} --php-fpm-path {args.php_fpm_path}"
    )

if args.subcommand == "run":
    print(f"在远程主机 {args.host} 上执行命令: {args.command}")
    remote_exec(args.command)
    print("命令执行完毕")

if args.subcommand == "stop":
    print("正在停止 Agent ...")
    remote_exec("python /usr/local/monit/agent.py stop")
    print("Agent 已停止")

if args.subcommand == "uninstall":
    print("正在卸载 Agent ...")
    remote_exec("python /usr/local/monit/agent.py uninstall")
    print("Agent 已卸载")

ssh_client.close()
