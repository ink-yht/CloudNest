package tencent

import (
	"context"
	"github.com/go-gomail/gomail"
)

type Service struct {
	authEmail string // 发送方的邮箱地址
	authPwd   string // 发送方的邮箱授权码（非QQ密码）
	smtpHost  string // SMTP服务器地址
	smtpPort  int    // SMTP服务器端口
}

func NewService(authEmail, authPwd, smtpHost string, smtpPort int) *Service {
	return &Service{
		authEmail: authEmail,
		authPwd:   authPwd,
		smtpHost:  smtpHost,
		smtpPort:  smtpPort,
	}
}

// Send 邮件发送方法
func (s *Service) Send(ctx context.Context, email string, subject, body string) error {
	d := gomail.NewDialer(s.smtpHost, s.smtpPort, s.authEmail, s.authPwd)
	m := gomail.NewMessage()
	m.SetHeader("From", s.authEmail)
	m.SetHeader("To", email)
	m.SetHeader("Subject", subject)
	m.SetBody("text/html", body)
	return d.DialAndSend(m)
}
