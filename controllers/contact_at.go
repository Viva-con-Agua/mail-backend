package controllers

import (
	"log"
	"mail-backend/dao"
	"mail-backend/mail"
	"mail-backend/models"

	"github.com/Viva-con-Agua/vcago/verror"
	"github.com/Viva-con-Agua/vcago/vlog"
	"github.com/Viva-con-Agua/vcago/vmod"
	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson"
)

func InsertContactAt(c echo.Context) (err error) {
	var ctx = c.Request().Context()
	body := new(models.ContactAtCreate)
	if verr := verror.JSONValidate(c, body); verr != nil {
		return c.JSON(verr.Response())
	}
	create := body.Insert()
	d := dao.NewContactAtDAO(ctx)
	result, verr := d.Insert(create)
	if verr != nil {
		vlog.Print(c.Request(), verr)
		return c.JSON(verr.Response())
	}
	job, err := dao.GetJobWithSubs(ctx, bson.M{"case": "contact_at", "scope": "default"})
	if err != nil {
		log.Print(err)
		return
	}
	template := models.GetTemplate(job.Templates, "de")
	if template == nil {
		log.Print(err)
		return
	}
	sendMail := mail.NewSendMail(
		&job.Email,
		"kontakt@vivaconagua.org",
		template.Subject,
		result,
		template.HTML,
	)
	sendMail, err = sendMail.CreateBody()
	if err != nil {
		log.Print(err)
		return
	}
	err = sendMail.Send()
	if err != nil {
		log.Print(err)
		return
	}

	return c.JSON(vmod.RespCreated(result, d.Coll()))
}

//ListCampaigns handler for list jobs by query filter
func ListContactAt(c echo.Context) (err error) {
	var ctx = c.Request().Context()
	d := dao.NewContactAtDAO(ctx)
	result, verr := d.List()
	if verr != nil {
		vlog.Print(c.Request(), verr)
		return c.JSON(verr.Response())
	}
	return c.JSON(vmod.RespSelected(result, d.Coll()))
}
