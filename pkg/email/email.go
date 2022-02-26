/*
	对发送邮件进行封装
*/
package email

import (
	"crypto/tls"
	"gopkg.in/gomail.v2"
)

type Email struct {
	*SMTPInfo
}

type SMTPInfo struct {
	Host     string
	Port     int
	IsSSl    bool
	UserName string
	Password string
	From     string
}

func NewMail(info *SMTPInfo) *Email {
	return &Email{SMTPInfo: info}
}

/*
	SendMail 发送邮件
		首先调用 NewMessage 创建一个消息实例，用于设置邮件的必要信息 发件人From 收件人To 邮件主题Subject 正文Body
		接着 NewDialer 创建一个新的 SMTP 拨号实例，设置对应的拨号信息 用于连接 SMTP 服务器最后调用 DialAndSend 打开与 SMTP
		服务器的连接并发送电子邮件
*/
func (e *Email) SendMail(to []string, subject, body string) error {
	m := gomail.NewMessage()
	m.SetHeader("From", e.From)
	m.SetHeader("To", to...)
	m.SetHeader("Subject", subject)
	m.SetBody("text/html", body)

	dialer := gomail.NewDialer(e.Host, e.Port, e.UserName, e.Password)
	dialer.TLSConfig = &tls.Config{InsecureSkipVerify: e.IsSSl}
	return dialer.DialAndSend(m)
}
