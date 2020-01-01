package email

import (
	"fmt"
	"strings"

	"github.com/go-gomail/gomail"
)

func defaultSend() {
	m := gomail.NewMessage()
	m.SetAddressHeader("From", "", "")
	m.SetAddressHeader("To", "", "")
	m.SetAddressHeader("Cc", "", "")
	m.SetHeader("Subject", "Hello!")
	m.SetBody("text/html", "Hello <b>Bob</b> and <i>Cora</i>!")
	m.Attach("./email/附件/fujian.jpg")
	d := gomail.NewDialer("smtp.exmail.qq.com", 465, "", "")
	fmt.Println("begin...")
	// Send the email to Bob, Cora and Dan.
	err := d.DialAndSend(m)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("done...")
}

//  GetMessage 构造邮件
func GetMessage(from, tos, ccs, subject, body, filenames string) *gomail.Message {
	m := gomail.NewMessage()
	m.SetHeader("From", from)
	m.SetHeader("To", strings.Split(tos, ",")...)
	if ccs != "" {
		m.SetHeader("Cc", strings.Split(ccs, ",")...)
	}
	m.SetHeader("Subject", subject)
	m.SetBody("text/html", body)
	for _, value := range strings.Split(filenames, ",") {
		m.Attach(value)
	}
	return m
}

// Send 发送
func Send(host, acc, pwd string, port int, messages ...*gomail.Message) {
	d := gomail.NewDialer(host, port, acc, pwd)
	fmt.Println("begin...")
	for _, value := range messages {
		err := d.DialAndSend(value)
		if err != nil {
			fmt.Println(err)
		}
	}
	fmt.Println("done...")
}

// TencentEmailSend 腾讯企业邮
func TencentEmailSend(messages ...*gomail.Message) {
	Send("smtp.exmail.qq.com", "", "", 465, messages...)
}
