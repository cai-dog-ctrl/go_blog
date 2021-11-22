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
	IsSSL    bool
	UserName string
	Password string
	From     string
}

// NewEmail 新建Email实体
func NewEmail(info *SMTPInfo) *Email {
	return &Email{SMTPInfo: info}
}

// SendMail 发送邮件
func (e *Email) SendMail(to []string, subject, body string) error {
	m := gomail.NewMessage()
	m.SetHeader("From", e.From)                                        //设置发件人
	m.SetHeader("To", to...)                                           //设置收件人
	m.SetHeader("Subject", subject)                                    //邮件主题
	m.SetBody("text/html", body)                                       //邮件正文
	dialer := gomail.NewDialer(e.Host, e.Port, e.UserName, e.Password) //返回一个新的 SMTP 拨号程序。
	// 给定的参数用于连接到 SMTP 服务器。
	dialer.TLSConfig = &tls.Config{InsecureSkipVerify: e.IsSSL}
	return dialer.DialAndSend(m)
}
