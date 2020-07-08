package mail

import (
	"crypto/tls"
	"fmt"
	"log"
	"mail-backend/utils"
	"net"
	"net/mail"
	"net/smtp"
)

type (
	Msg struct {
		Subj string
		Body string
	}
)

func SignUp(to string, token string) {
	msg := new(Msg)
	msg.Subj = "Last SignUp step"
	msg.Body = token
	SendMail(to, msg)
}

func SendMail(t string, msg *Msg) {
	from := mail.Address{"", utils.Config.Mail.From}
	to := mail.Address{"", t}

	//setup headers
	headers := make(map[string]string)
	headers["From"] = from.String()
	headers["To"] = to.String()
	headers["Subject"] = msg.Subj

	// Setup message
	message := ""
	for k, v := range headers {
		message += fmt.Sprintf("%s: %s\r\n", k, v)
	}
	message += "\r\n" + msg.Body

	// Connect to the SMTP Server
	host, _, _ := net.SplitHostPort(utils.Config.Mail.Host)

	auth := smtp.PlainAuth("", "", "", host)

	// TLS config
	tlsconfig := &tls.Config{
		InsecureSkipVerify: true,
		ServerName:         host,
	}

	c, err := smtp.Dial(utils.Config.Mail.Host)
	if err != nil {
		log.Print("utils.SendMail.Dial: ", err)
	}
	log.Print("Email : \n", msg)
	c.StartTLS(tlsconfig)

	// Auth
	if err = c.Auth(auth); err != nil {
		log.Print("utils.SendMail: ", err)
	}

	// To && From
	if err = c.Mail(from.Address); err != nil {
		log.Print("utils.SendMail: ", err)
	}

	if err = c.Rcpt(to.Address); err != nil {
		log.Print("utils.SendMail: ", err)
	}

	// Data
	w, err := c.Data()
	if err != nil {
		log.Print("utils.SendMail: ", err)
	}

	_, err = w.Write([]byte(message))
	if err != nil {
		log.Print("utils.SendMail: ", err)
	}

	err = w.Close()
	if err != nil {
		log.Print("utils.SendMail: ", err)
	}

	c.Quit()

}

func SendSignUpMail(address string, token string) error {

	auth := smtp.PlainAuth("", "", "", "172.3.200.3")
	to := []string{address}
	msg := []byte("To: " + address + "\r\n" +
		"Subject: Last SignUp step\r\n" +
		"\r\n" +
		token)

	err := smtp.SendMail(utils.Config.Mail.Host, auth, "noreply@vivaconagua.org", to, msg)
	if err != nil {
		log.Print("utils.SendSignUpMail: ", err)
	}
	return err
}
