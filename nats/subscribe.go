package nats

import (
	"log"
	"mail-backend/dao"
	"mail-backend/mail"
	"mail-backend/models"

	"github.com/Viva-con-Agua/vcago/vmod"
	"go.mongodb.org/mongo-driver/bson"
	"golang.org/x/net/context"
)

func SubscribeCode() {
	Nats.Subscribe("mail.code", func(m *vmod.MailCode) {
		ctx := context.Background()
		job, err := dao.GetJobWithSubs(bson.M{"case": m.Case, "scope": m.Scope})
		if err != nil {
			err.Log(nil)
		}
		
		if job == nil {
			job, err := dao.GetJobWithSubs(bson.M{"case": m.Case, "scope": "default"})
			if err != nil {
				err.Log(nil)
			}
		}
		//Process Mail
	})
}
