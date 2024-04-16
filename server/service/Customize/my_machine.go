package Customize

import (
	"errors"
	"fmt"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/Customize"
	"github.com/go-ping/ping"
	"golang.org/x/crypto/ssh"
	"strings"
	"sync"
	"time"
)

type MyMachineService struct {
}

var MachinesMap sync.Map // id->ip_addr

func init() {
	//tmpMachines := make([]struct {
	//	ID     uint
	//	IPAddr string
	//}, 0)
	//global.GVA_DB.Model(&Customize.Machine{}).Select("id", "ip_addr").Find(&tmpMachines)
	//for _, machine := range tmpMachines {
	//	MachinesMap.Store(machine.ID, machine.IPAddr)
	//}
	//
	//myMachineService := MyMachineService{}
	//myMachineService.MachineHeartBeat()
}

func ConvertTimestamp(timestamp string) string {
	// 2006-01-02 15:04:05 to 2006-01-02T15:04:05Z
	times := strings.Split(timestamp, " ")
	timestamp = times[0] + "T" + times[1] + "Z"
	return timestamp
}

func (m *MyMachineService) AddMachineHook(machine Customize.Machine) {
	MachinesMap.Store(machine.ID, machine.IPAddr)
}

// to: xxxx@qq.com
func (machineService *MyMachineService) SendEmail(to string, body string) (err error) {
	fmt.Println("package sendFunc sendMailSimple")
	// logrus.Info("Panic info is: %v", err)
	global.GVA_LOG.Info("SendMailSimple func start ")

	//auth := smtp.PlainAuth(
	//	"",                  //传空值即可，用不上
	//	"2476100824@qq.com", // **发送邮件的邮箱地址
	//	"zclhprbppvlgdjhf",  //POP3/IMAP/SMTP/Exchange/CardDAV/CalDAV服务的授权码，需要在QQ邮箱设置中开启
	//	"smtp.qq.com",       // qq邮箱SMTP 服务器地址
	//)

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

	//msg := []byte(s)

	//err = smtp.SendMail(
	//	"smtp.qq.com:587",
	//	auth,
	//	"2476100824@qq.com", //发送邮件的邮箱地址
	//	[]string{to},        //接收邮箱
	//	// []string{"xxx@qq.com", "xxx@qq.com", "xxx@qq.com"}, //接收邮箱
	//	msg,
	//)

	if err != nil {
		global.GVA_LOG.Error("smtp.SendMail error: " + err.Error())
	} else {
		global.GVA_LOG.Info("smtp.SendMail success! ")
	}
	return err
}

func (machineService *MyMachineService) ExecuteSSH(machine *Customize.Machine, cmds []string) (err error) {
	sshHost := machine.IPAddr
	sshUser := "root"
	sshPassword := machine.Password
	//sshType := "password"
	sshPort := 22

	//创建sshp登陆配置
	config := &ssh.ClientConfig{
		Timeout:         time.Second, //ssh 连接time out 时间一秒钟, 如果ssh验证错误 会在一秒内返回
		User:            sshUser,
		HostKeyCallback: ssh.InsecureIgnoreHostKey(), //这个可以, 但是不够安全
		//HostKeyCallback: hostKeyCallBackFunc(h.Host),
	}
	config.Auth = []ssh.AuthMethod{ssh.Password(sshPassword)}

	//dial 获取ssh client
	addr := fmt.Sprintf("%s:%d", sshHost, sshPort)
	sshClient, err := ssh.Dial("tcp", addr, config)
	if err != nil {
		global.GVA_LOG.Fatal("创建ssh client 失败" + err.Error())
	}
	defer sshClient.Close()

	//创建ssh-session
	session, err := sshClient.NewSession()
	if err != nil {
		global.GVA_LOG.Fatal("创建ssh session 失败" + err.Error())
	}
	defer session.Close()
	//执行远程命令

	combo, err := session.CombinedOutput(strings.Join(cmds, ";"))
	if err != nil {
		global.GVA_LOG.Fatal("远程执行cmd 失败" + err.Error())
	}
	global.GVA_LOG.Info("命令输出:" + string(combo))
	return nil
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
		MachinesMap.Range(func(machineID, MachineAddr interface{}) bool {
			machine := Customize.Machine{}
			global.GVA_DB.Where("id = ?", machineID).First(&machine)
			err := machineService.PingMachine(MachineAddr.(string))
			newStatus := err == nil
			if newStatus != *machine.Status {
				machine.Status = &newStatus
				global.GVA_DB.Model(&Customize.Machine{}).Save(&machine)
			}
			return true
		})
		time.Sleep(10 * time.Second) // ping duration
	}
}
