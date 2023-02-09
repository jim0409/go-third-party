package main

import (
	"fmt"
	"net/smtp"
)

// from/ to/ subject/ body
var basicmsg = func(from, to, subject, body string) string {
	return fmt.Sprintf("From: %v\nTo: %v\nSubject: %v\n\n%v", from, to, subject, body)
}

type ISender interface {
	SendMail(from, to, subject string, id int) error
}

type Sender struct {
	Auth       string
	User       string
	Host       string
	SmtpServer string
	Template   *map[int]string
}

func NewSender(auth, user string, host string, smtp string, template *map[int]string) ISender {
	return &Sender{
		Auth:       auth,
		User:       user,
		Host:       host,
		SmtpServer: smtp,
		Template:   template,
	}
}

func (s *Sender) SendMail(from, to, subject string, id int) error {
	pass := s.Auth
	user := s.User
	msg := basicmsg(from, to, subject, (*s.Template)[id])

	err := smtp.SendMail(s.SmtpServer,
		smtp.PlainAuth("", user, pass, s.Host),
		from, []string{to}, []byte(msg))

	return err

}
