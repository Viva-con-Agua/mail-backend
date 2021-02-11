package main

import (
	"mail-backend/controllers"
	"mail-backend/dao"
	"mail-backend/env"
	"mail-backend/nats"
	"os"

	"github.com/Viva-con-Agua/vcago"
	"github.com/Viva-con-Agua/vcago/verr"
	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {

	// intial loading function
	env.LoadConfig()
	dao.Connect()
	nats.Connect()
	nats.Subscribe()
	dao.Init()
	cors := middleware.CORSConfig{
		AllowOrigins:     env.AllowOrigins,
		AllowHeaders:     []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept, echo.HeaderAuthorization},
		AllowCredentials: true,
	}
	//nats.SubscribeAddModel()
	//create echo server
	e := echo.New()
	e.Use(middleware.CORSWithConfig(cors))
// Middleware
	e.Use(middleware.CORSWithConfig(vcago.NewCORSConfig()))
	e.Validator = &verr.JSONValidator{Validator: validator.New()}
	admin := e.Group("/admin")
	admin.GET("/email/init", controllers.InitModels)
	admin.POST("/email/email", controllers.InsertEmail)
	admin.POST("/email/job", controllers.InsertJob)
	//e.GET("/", controllers.Mail)
	apiV1 := e.Group("/v1")
	email := apiV1.Group("/email")
	template := email.Group("/template")
	template.POST("", controllers.InsertTemplate)
	template.GET("", controllers.ListTemplate)
	template.PUT("", controllers.UpdateTemplate)

	jobs := email.Group("/jobs")
	jobs.GET("", controllers.ListJob)

	if port, ok := os.LookupEnv("REPO_PORT"); ok {
		e.Logger.Fatal(e.Start(":" + port))
	} else {
		e.Logger.Fatal(e.Start(":1323"))
	}
}
