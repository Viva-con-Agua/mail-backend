package controllers

import (
	"mail-backend/dao"
	"mail-backend/models"
	"net/http"

	"github.com/labstack/echo/v4"
)


//InitModels inserts initial EmailTemplates into the database
func InitModels(c echo.Context) error {
	htmlRegister := `
<!DOCTYPE html>
<html>

</head>

<body>
  <p>
    <strong>Hello {{.UserName}}</strong> <br>
    <a href="http://localhost:8080/confirm/{{.Code}}">Link</a>

  </p>

</body>

</html>
    `
	defaultRegister := &models.TemplateCreate{
		Name:        "register_default",
		Scope:       "default",
		Case:        "register",
		HTML:        htmlRegister,
		Tags:        []string{"default", "register"},
		Description: "default register template to provide register process",
		Type:        "code",
	}
	template, apiErr := dao.InsertTemplate(c.Request().Context(), defaultRegister)
	if apiErr != nil {
        apiErr.Log(c)
		return c.JSON(apiErr.Code, apiErr.Body)
	}
    defaultEmail := &models.EmailAddressCreate{
      Email: "dennis_kleber@mailbox.org",
      Tags: []string{"register", "password_reset", "confim"},
   }
   email, apiErr := dao.InsertEmailAddress(c.Request().Context(), defaultEmail)
   if apiErr != nil {
      apiErr.Log(c)
      return c.JSON(apiErr.Code, apiErr.Body)
   }
   defaultJob := &models.JobCreate{
   	Name: "register_default",     
		Scope: "default",
		Case: "register",  
		TemplateID: template.ID,
		EmailID: email.ID,
   }
	job, apiErr := dao.InsertJob(c.Request().Context(), defaultJob)
   if apiErr != nil {
      apiErr.Log(c)
      return c.JSON(apiErr.Code, apiErr.Body)
   }

	type result struct {
		Template models.Template
		Email models.EmailAddress
		Job models.Job
	}
	return c.JSON(http.StatusCreated, result{Template: *template, Email: *email, Job: *job})
}
