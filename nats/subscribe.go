package nats

import (
	"log"
	"mail-backend/dao"
	"mail-backend/mail"

	"github.com/Viva-con-Agua/vcago/verr"
	"github.com/Viva-con-Agua/vcago/vmod"
	"go.mongodb.org/mongo-driver/bson"
	"golang.org/x/net/context"
)

//Subscribe starts all subscribtions
func Subscribe() {
	SubscribeCode()
}

//SubscribeCode used for email code model
func SubscribeCode() {
	Nats.Subscribe("mail.code", func(m *vmod.MailCode) {
		ctx := context.Background()
		log.Print(m.JobCase)
		job, err := dao.GetJobWithSubs(ctx, bson.M{"case": m.JobCase, "scope": m.JobScope})
		if job == nil {
			job, err = dao.GetJobWithSubs(ctx, bson.M{"case": m.JobCase, "scope": "default"})
			if err == nil {
				sendMail := mail.NewSendMail(
					job.Email.Email,
					m.To,
					job.Template.Subject,
					m,
					job.Template.Name,
					job.Template.HTML,
				)
				sendMail, err := sendMail.CreateBody()
				if err != nil {
					log.Print(verr.ErrorWithColor, err)
				}
				err = sendMail.Send()
				if err != nil {
					log.Print(verr.ErrorWithColor, err)
				}
			} else {
				log.Print(verr.ErrorWithColor, err)
			}
		}
	})
}
