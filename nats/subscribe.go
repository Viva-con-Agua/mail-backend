package nats

import (
	"log"
	"mail-backend/mail"
	"mail-backend/models"
)

func SubscribeToken() {
	_, err := Nats.Subscribe("mail.token", func(m *models.MailInfo) {
		mail.SendSignUpMail(m.To, m.Token)
	})
	if err != nil {
		log.Print("Nats Error: ", err)
	}
}
