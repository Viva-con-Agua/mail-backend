package main

import (
	"log"
	"mail-backend/controllers"
	"mail-backend/dao"
	"mail-backend/env"
	"os"
	"strings"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {

	// intial loading function
	env.LoadConfig()
	log.Print(strings.Split(os.Getenv("ALLOW_ORIGINS"), ","))
	dao.Connect()
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
	admin := e.Group("/admin")
	admin.GET("/email/init", controllers.InitModels)
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
