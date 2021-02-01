package mail

import (
	"log"
	"bytes"
	"html/template"
	//"mail-backend/utils"
	"gopkg.in/gomail.v2"
)

type (
	Msg struct {
		Subj string
		Body string
	}
)
/*
//ProcessMail
func ProcessMail(m vmod.Mail, t interface{}) {
	msg := new(Msg)
	msg.Subj = "Last SignUp step"
	msg.Body = token
	SendMail(to, msg)
}*/

type TokenMail struct {
	Name string
	Token string
}
type Template struct {
	Subject string
	From string
	TO string
}

func SendMail() error {
	tem := Template{
		Subject: "SignUp", 
		From: "dennis_kleber@mailbox.org", 
		TO: "d.kleber@vivaconagua.org",
	}
	i := TokenMail{Name: "Dennis", Token: "asdasdasdasd"}
	t := template.New("login.html")
	var err error
	t, err = t.ParseFiles("templates/login.html")
	if err != nil {
		log.Println(err)
	}

	log.Print(t)
	var tpl bytes.Buffer
	if err := t.Execute(&tpl, i); err != nil {
		log.Println(err)
	}

	result := tpl.String()

	m := gomail.NewMessage()
	m.SetHeader("From", tem.From)
	m.SetHeader("To", tem.TO)
	m.SetHeader("Subject", tem.Subject)
	m.SetBody("text/html", result)
	m.Attach("template.html")// attach whatever you want

	//auth := smtp.PlainAuth("", "", "", utils.Config.Mail.Host)
	//to := []string{i.TO}
	//msg := []byte("To: " + i.TO + "\r\n" +
	//	"Subject: Last SignUp step\r\n" +
	//	"\r\n" +
	//	token)
	d := gomail.NewDialer("127.0.0.1", 25, "", "")

	// Send the email to Bob, Cora and Dan.
	if err := d.DialAndSend(m); err != nil {
		return err
	}
	return err
}
