package nats

import (
	"log"
	"mail-backend/mail"
	"mail-backend/models"

	"github.com/Viva-con-Agua/echo-pool/nats"
)

func SubscribeSignUp() {
	_, err := nats.Nats.Subscribe("mail.signup", func(m *models.MailInfo) {
		mail.SignUp(m.To, m.Token)
	})
	if err != nil {
		log.Print("Nats Error: ", err)
	}

}
