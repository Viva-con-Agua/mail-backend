package main

import (
	"log"
	"mail-backend/controllers"
	"mail-backend/dao"
	"os"
	"strings"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
)

func main() {

	// intial loading function
	godotenv.Load()
	log.Print(strings.Split(os.Getenv("ALLOW_ORIGINS"), ","))
	dao.Connect()
	dao.Init()
	//nats.SubscribeAddModel()
	//create echo server
	e := echo.New()
	admin := e.Group("/admin")
	admin.GET("/email/template", controllers.InitTemplate)
	admin.GET("/email/email", controllers.InitAddress)
	//e.GET("/", controllers.Mail)
	apiV1 := e.Group("/v1")
	email := apiV1.Group("/email")
	template := email.Group("/template")
	template.POST("", controllers.InsertTemplate)
	template.GET("", controllers.ListTemplate)
	template.PUT("", controllers.UpdateTemplate)


	
	if port, ok := os.LookupEnv("REPO_PORT"); ok {
		e.Logger.Fatal(e.Start(":" + port))
	} else {
		e.Logger.Fatal(e.Start(":1324"))
	}
}
