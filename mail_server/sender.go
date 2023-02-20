package main

import (
	"net/smtp"
)

type ISender interface {
	SendMail(from, to, subject string, id int) error
}

type Sender struct {
	Auth       string
	User       string
	Host       string
	SmtpServer string
	Template   *map[int]string
	BodyForm   map[string]interface{}
}

func NewSender(auth, user string, host string, smtp string, template *map[int]string, body map[string]interface{}) ISender {
	return &Sender{
		Auth:       auth,
		User:       user,
		Host:       host,
		SmtpServer: smtp,
		Template:   template,
		BodyForm:   body,
	}
}

func (s *Sender) SendMail(from, to, sub string, id int) error {
	mime := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n"
	pass := s.Auth
	user := s.User
	subject := "Subject: " + sub + "!\n"

	body, err := ParseTemplate(MsgTemplate[id], s.BodyForm)
	if err != nil {
		return err
	}

	msg := []byte(subject + mime + body)

	err = smtp.SendMail(s.SmtpServer,
		smtp.PlainAuth("", user, pass, s.Host),
		from, []string{to}, msg)

	return err

}
