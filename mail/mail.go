package mail

import (
	"bytes"
	"crypto/tls"
	"encoding/base64"
	"html/template"
	"mail-backend/env"
	"mail-backend/models"

	"gopkg.in/gomail.v2"
)

type (
	//SendMail represents a Email type used for sending an email to the smtp relay.
	SendMail struct {
		Host      string
		Port      int
		Password  string
		From      string
		To        string
		Subject   string
		Body      string
		Interface interface{}
		Name      string
		HTML      string
	}
)

func NewSendMail(email *models.EmailAddress, to string, subject string, inter interface{}, html string) *SendMail {
	return &SendMail{
		Host: email.Host,
		Port: email.Port,
		Password: email.Password,
		From:      email.Email,
		To:        to,
		Subject:   subject,
		Interface: inter,
		HTML:      html,
	}
}

//CreateBody creates an html/template type from s.HTML using s.Interface and store it in the s.Body param as string
func (s *SendMail) CreateBody() (*SendMail, error) {
	t := template.New(s.Name)
	html, _ := base64.StdEncoding.DecodeString(s.HTML)
	t, err := t.Parse(string(html))
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
	if env.LogLevel == "debug" {
		d := gomail.Dialer{Host: env.MailSMTPHost, Port: env.MailSMTPPort, TLSConfig: &tls.Config{InsecureSkipVerify: true}}
		// Send the email to Bob, Cora and Dan.
		err := d.DialAndSend(m)
		if err != nil {
			return err
		}
	} else {
		dprod := gomail.NewDialer(s.Host, s.Port, s.From, s.Password)
		err := dprod.DialAndSend(m)
		if err != nil {
			return err
		}
	}
	return nil

}
