package controllers

import (
	"mail-backend/dao"
	"mail-backend/models"
	"net/http"

	"github.com/labstack/echo/v4"
)


//InitTemplate inserts initial EmailTemplates into the database
func InitTemplate(c echo.Context) error {
	htmlRegister := `
<!DOCTYPE html>
<html>

</head>

<body>
  <p>
    <strong>Hello {{.Name}}</strong> <br>
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
	result, apiErr := dao.InsertTemplate(c.Request().Context(), defaultRegister)
	if apiErr != nil {
        apiErr.Log(c)
		return c.JSON(apiErr.Code, apiErr.Body)
	}
	return c.JSON(http.StatusCreated, result)
}
//InitAddress inserts initial EmailAdress into database
func InitAddress(c echo.Context) error {
   defaultEmail := &models.EmailAddressCreate{
      Email: "dennis_kleber@mailbox.org",
      Tags: []string{"register", "password_reset", "confim"},
   }
   result, apiErr := dao.InsertEmailAddress(c.Request().Context(), defaultEmail)
   if apiErr != nil {
      apiErr.Log(c)
      return c.JSON(apiErr.Code, apiErr.Body)
   }
   return c.JSON(http.StatusCreated, result)

}
