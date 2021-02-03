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
	body := new(models.TemplateCreate)
	if err = verr.JSONValidate(c, body); err != nil {
		return
	}

	resp, apiErr := dao.InsertTemplate(c.Request().Context(), body)
	if apiErr != nil {
		apiErr.Log(c)
		return c.JSON(apiErr.Code, apiErr.Body)
	}
	return c.JSON(http.StatusCreated, resp)
}

//ListTemplate echo controller for handling email template list.
func ListTemplate(c echo.Context) (err error) {
	resp, apiErr := dao.ListTemplate(c.Request().Context())
	if apiErr != nil {
		apiErr.Log(c)
		return c.JSON(apiErr.Code, apiErr.Body)
	}
	return c.JSON(http.StatusOK, resp)
}
//UpdateTemplate echo controller for handling email templates updates
func UpdateTemplate(c echo.Context) (err error) {
	body := new(models.Template)
	if err = verr.JSONValidate(c, body); err != nil {
		return
	}
	resp, apiErr := dao.UpdateTemplate(c.Request().Context(), body)
	if apiErr != nil {
		apiErr.Log(c)
		return c.JSON(apiErr.Code, apiErr.Body)
	}
	return c.JSON(http.StatusCreated, resp)

}

