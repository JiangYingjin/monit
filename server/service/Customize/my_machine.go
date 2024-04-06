package Customize

import (
	"fmt"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"net/smtp"
)

type MyMachineService struct {
}

// to: xxxx@qq.com
func (machineService *MyMachineService) SendEmail(to string, body string) (err error) {
	fmt.Println("package sendFunc sendMailSimple")
	// logrus.Info("Panic info is: %v", err)
	global.GVA_LOG.Info("SendMailSimple func start ")

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
	fmt.Println(s)

	msg := []byte(s)

	err = smtp.SendMail(
		"smtp.qq.com:587",
		auth,
		"2476100824@qq.com", //发送邮件的邮箱地址
		[]string{to},        //接收邮箱
		// []string{"xxx@qq.com", "xxx@qq.com", "xxx@qq.com"}, //接收邮箱
		msg,
	)

	if err != nil {
		global.GVA_LOG.Error("smtp.SendMail error: " + err.Error())
	} else {
		global.GVA_LOG.Info("smtp.SendMail success! ")
	}
	return err
}
