package main

import (
	"mail-backend/nats"
	"mail-backend/utils"

	"github.com/go-playground/validator"
	"github.com/labstack/echo"
)

type (
	CustomValidator struct {
		validator *validator.Validate
	}
)

func (cv *CustomValidator) Validate(i interface{}) error {
	return cv.validator.Struct(i)
}

func main() {

	// intial loading function
	utils.LoadConfig()
	nats.Connect()
	nats.SubscribeToken()
	//create echo server
	e := echo.New()
	e.Logger.Fatal(e.Start(":1333"))
}
