package Customize

import (
	"bytes"
	"errors"
	"fmt"
	"github.com/go-ping/ping"
	"my-server/global"
	"my-server/model/Customize"
	"net/smtp"
	"os/exec"
	"strings"
	"sync"
	"time"
)

type MyMachineService struct {
}

var MachinesMap sync.Map // id->ip_addr

func init() {
	go func() {
		time.Sleep(10 * time.Second) // wait for db init
		myMachineService := MyMachineService{}
		myMachineService.MachineHeartBeat()
	}()
}

func ConvertTimestamp(timestamp string) string {
	// 2006-01-02 15:04:05 to 2006-01-02T15:04:05Z
	times := strings.Split(timestamp, " ")
	timestamp = times[0] + "T" + times[1] + "Z"
	return timestamp
}

func (machineService *MyMachineService) AddMachineHook(machine Customize.Machine) {
	MachinesMap.Store(machine.ID, machine.IPAddr)
}

// to: xxxx@qq.com
func (machineService *MyMachineService) SendEmail(to string, body string) (err error) {
	auth := smtp.PlainAuth(
		"",                  //传空值即可，用不上
		"2476100824@qq.com", // **发送邮件的邮箱地址
		"zclhprbppvlgdjhf",  //POP3/IMAP/SMTP/Exchange/CardDAV/CalDAV服务的授权码，需要在QQ邮箱设置中开启
		"smtp.qq.com",       // qq邮箱SMTP 服务器地址
	)

	// 以下的参数组成msg内容(即为邮件内容)
	/*
		To:%s  //接收邮箱
		From:%s<%s> // user<发送邮箱>
		Subject:%s //主题： 邮件主题
		mailtype %s //Content-Type: text/html; charset=UTF-8 内容格式
		body %s  // 邮件正文内容
	*/
	user := "2476100824" //用户名 可任意字符串
	subject := "email from go"
	// mailtype := "Content-Type: text/plain;charset=UTF-8" //body内容的格式  此为纯文本显示
	mailtype := "Content-Type: text/html; charset=UTF-8" //body内容的格式  此为html格式渲染
	//body := "<h1>This is the body</h1> of my email"      //邮件正文

	//！！！msg的格式很重要！！！
	//要发送的消息，可以直接写在[]bytes里，但是看着太乱，因此使用格式化
	s := fmt.Sprintf("To:%s\r\nFrom:%s<%s@qq.com>\r\nSubject:%s\r\n%s\r\n\r\n%s",
		to, user, user, subject, mailtype, body)
	global.GVA_LOG.Info("send email: \n" + s)

	msg := []byte(s)

	err = smtp.SendMail(
		"smtp.qq.com:587",
		auth,
		"2476100824@qq.com", //发送邮件的邮箱地址
		[]string{to},        //接收邮箱
		//[]string{"xxx@qq.com", "xxx@qq.com", "xxx@qq.com"}, //接收邮箱
		msg,
	)

	if err != nil {
		global.GVA_LOG.Error("smtp.SendMail error: " + err.Error())
	} else {
		global.GVA_LOG.Info("smtp.SendMail success! ")
	}
	return err
}

func (machineService *MyMachineService) FormCmdParams(host string, params ...string) []string {
	cmdParams := make([]string, 0, len(params)+1)
	cmdParams = append(cmdParams, "-")
	cmdParams = append(cmdParams, "--host="+host)
	cmdParams = append(cmdParams, "--port=22")
	cmdParams = append(cmdParams, params...)
	return cmdParams
}

func (machineService *MyMachineService) ExecuteCmd(params []string) (string, error) {
	curlCmd := exec.Command("curl", "-sL", "file.jiangyj.tech/proj/monit/remote.py")
	pythonScript, _ := curlCmd.CombinedOutput()

	pythonCmd := exec.Command("python", params...)
	pythonCmd.Stdin = bytes.NewReader(pythonScript)

	//if params[0] == "-" {
	//	params[0] = "../agent/remote.py"
	//}
	//pythonCmd := exec.Command("python", params...)
	//pythonCmd := exec.Command("C:\\Users\\24761\\AppData\\Local\\Programs\\Python\\Python311\\python.exe", params...)

	// 执行python命令并等待结果
	outputByte, err := pythonCmd.CombinedOutput()

	if err != nil {
		global.GVA_LOG.Error("execute cmd error: " + err.Error() + ", output: " + string(outputByte))
		return string(outputByte), err
	} else {
		return string(outputByte), nil
	}

	//sshHost := machine.IPAddr
	//sshUser := "root"
	//sshPassword := machine.Password
	////sshType := "password"
	//sshPort := 22
	//
	////创建sshp登陆配置
	//config := &ssh.ClientConfig{
	//	Timeout:         time.Second, //ssh 连接time out 时间一秒钟, 如果ssh验证错误 会在一秒内返回
	//	User:            sshUser,
	//	HostKeyCallback: ssh.InsecureIgnoreHostKey(), //这个可以, 但是不够安全
	//	//HostKeyCallback: hostKeyCallBackFunc(h.Host),
	//}
	//config.Auth = []ssh.AuthMethod{ssh.Password(sshPassword)}
	//
	////dial 获取ssh client
	//addr := fmt.Sprintf("%s:%d", sshHost, sshPort)
	//sshClient, err := ssh.Dial("tcp", addr, config)
	//if err != nil {
	//	global.GVA_LOG.Fatal("创建ssh client 失败" + err.Error())
	//}
	//defer sshClient.Close()
	//
	////创建ssh-session
	//session, err := sshClient.NewSession()
	//if err != nil {
	//	global.GVA_LOG.Fatal("创建ssh session 失败" + err.Error())
	//}
	//defer session.Close()
	////执行远程命令
	//
	//combo, err := session.CombinedOutput(strings.Join(cmds, ";"))
	//if err != nil {
	//	global.GVA_LOG.Fatal("远程执行cmd 失败" + err.Error())
	//}
	//global.GVA_LOG.Info("命令输出:" + string(combo))
	//return nil
}

func (machineService *MyMachineService) PingMachine(machineIP string) (err error) {
	pinger, err := ping.NewPinger(machineIP)
	if err != nil {
		global.GVA_LOG.Error(fmt.Sprintf("NewPinger error:%s\n", err.Error()))
	}
	// 设置ping包数量
	pinger.Count = 5
	// 设置超时时间
	pinger.Timeout = time.Second * 10
	// 设置成特权模式
	pinger.SetPrivileged(true)
	// 运行pinger
	err = pinger.Run()
	if err != nil {
		fmt.Printf("ping异常：%s\n", err.Error())
	}
	stats := pinger.Statistics()
	// 如果回包大于等于1则判为ping通
	if stats.PacketsRecv >= 1 {
		return nil
	} else {
		return errors.New("IP can not reach" + machineIP)
	}
}

func (machineService *MyMachineService) MachineHeartBeat() {
	for {
		if global.GVA_DB == nil {
			time.Sleep(10 * time.Second)
			continue
		}

		tmpMachines := make([]struct {
			ID     uint   `gorm:"primarykey" json:"ID"`
			IPAddr string `json:"ip_addr" form:"ip_addr" gorm:"column:ip_addr;comment:;" binding:"required"`
			Status bool   `json:"status" form:"status" gorm:"column:status;comment:机器是否在线;"`
		}, 0)
		global.GVA_DB.Model(&Customize.Machine{}).Select("id", "ip_addr", "status").Find(&tmpMachines)
		for _, machine := range tmpMachines {
			err := machineService.PingMachine(machine.IPAddr)
			newStatus := err == nil
			if newStatus != machine.Status {
				global.GVA_DB.Model(&Customize.Machine{}).Where("id = ?", machine.ID).Updates(map[string]interface{}{"status": newStatus})
			}
		}

		time.Sleep(10 * time.Second) // ping duration
	}
}
