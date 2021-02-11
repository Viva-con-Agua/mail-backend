package controllers

import (
	"mail-backend/dao"
	"mail-backend/models"
	"net/http"

	"github.com/labstack/echo/v4"
)

//InitModels inserts initial EmailTemplates into the database
func InitModels(c echo.Context) (err error) {
	var htmlRegister = `
<!DOCTYPE html> 
<html>
</head>
<body>
  <p>
    <strong>Hello {{.UserName}}</strong> <br>
    <a href="https://auth.vivaconagua.org/#/confirm/{{.Code}}">Link</a>
  </p>
</body>
</html>
    `
	defaultRegister := &models.Template{
		Subject: "Finish Register",
		HTML:        htmlRegister,
		Country:     "en_EN",
	}
	defaultEmail := &models.EmailAddressCreate{
		Email: "register@vivaconagua.org",
		Tags:  []string{"register", "password_reset", "confim"},
	}
	email := defaultEmail.Insert()
	err = dao.InsertEmailAddress(c.Request().Context(), email)
	if err != nil {
		return
	}
	defaultJob := &models.JobCreate{
		Name:       "register_default",
		Scope:      "default",
		Case:       "register",
		Templates: []models.Template{*defaultRegister},
		EmailID:    email.ID,
	}
	var job = defaultJob.Insert()
	err = dao.InsertJob(c.Request().Context(), job)
	if err != nil {
		return
	}

	type result struct {
		Email    models.EmailAddress
		Job      models.Job
	}
	return c.JSON(http.StatusCreated, result{Email: *email, Job: *job})
}
