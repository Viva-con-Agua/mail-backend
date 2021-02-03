package mail

import (
	"bytes"
	"crypto/tls"
	"html/template"
	"mail-backend/env"

	"gopkg.in/gomail.v2"
)

type (
	//SendMail represents a Email type used for sending an email to the smtp relay.
	SendMail struct {
		From      string
		To        string
		Subject   string
		Body      string
		Interface interface{}
		Name      string
		HTML      string
	}
)

func NewSendMail(from string, to string, subject string, inter interface{}, name string, html string) *SendMail {
	return &SendMail{
		From: from,
		To: to,
		Subject: subject,
		Interface: inter,
		Name: name,
		HTML: html,
	}
}

//CreateBody creates an html/template type from s.HTML using s.Interface and store it in the s.Body param as string
func (s *SendMail) CreateBody() (*SendMail, error) {
	t := template.New(s.Name)
	t, err := t.Parse(s.HTML)
	if err != nil {
		return nil, err
	}
	var tpl bytes.Buffer
	if err := t.Execute(&tpl, s.Interface); err != nil {
		return nil, err
	}
	s.Body = tpl.String()
	return s, nil
}

//Send connects to the smtp relay and sends the email.
func (s *SendMail) Send() error {
	m := gomail.NewMessage()
	m.SetHeader("From", s.From)
	m.SetHeader("To", s.To)
	m.SetHeader("Subject", s.Subject)
	m.SetBody("text/html", s.Body)
	d := gomail.Dialer{Host: env.MailSMTPHost, Port: env.MailSMTPPort, TLSConfig: &tls.Config{InsecureSkipVerify: true}}

	// Send the email to Bob, Cora and Dan.
	err := d.DialAndSend(m)
	if err != nil {
		return err
	}
	return err

}
