package controllers

import (
	"mail-backend/dao"
	"mail-backend/models"

	"github.com/Viva-con-Agua/vcago/verr"
	"github.com/Viva-con-Agua/vcago/vmod"
	"github.com/labstack/echo/v4"
)

//InsertEmail handler for inserting job into jobs collection.
func InsertEmail(c echo.Context) (err error) {
    var ctx = c.Request().Context()
    body := new(models.EmailAddress)
    if err = verr.JSONValidate(c, body); err != nil {
        return
    }
    //var address = body.Insert()
    if err = dao.InsertEmailAddress(ctx, body); err != nil {
        return
    }
    return c.JSON(vmod.RespCreated(body, "mail_address"))
}

