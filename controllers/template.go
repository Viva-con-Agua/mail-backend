package controllers

import (
	"mail-backend/models"
	"mail-backend/dao"
	"net/http"
	"github.com/Viva-con-Agua/vcago/verr"
	"github.com/labstack/echo/v4"
)


//InsertTemplate echo controller for handling email template insert 
func InsertTemplate(c echo.Context) (err error) {
	body := new(models.Template)
	if err = verr.JSONValidate(c, body); err != nil {
		return
	}
	resp, err := dao.InsertTemplate(c.Request().Context(), body)
	if err != nil {
		return 
	}
	return c.JSON(http.StatusCreated, resp)
}

//ListTemplate echo controller for handling email template list.
func ListTemplate(c echo.Context) (err error) {
	resp, err := dao.ListTemplate(c.Request().Context())
	if err != nil {
		return
	}
	return c.JSON(http.StatusOK, resp)
}
//UpdateTemplate echo controller for handling email templates updates
func UpdateTemplate(c echo.Context) (err error) {
	body := new(models.Template)
	if err = verr.JSONValidate(c, body); err != nil {
		return
	}
	err = dao.UpdateTemplate(c.Request().Context(), body)
	if err != nil {
		return
	}
	return c.JSON(http.StatusCreated, body)
}

