package dao

import (
	//	"html"
	"mail-backend/models"

	"github.com/labstack/gommon/log"
	"golang.org/x/net/context"
)

//Init initial mongo db
func Init() {
    //initial Email Template
    if _, err := DB.Collection("email_templates").Indexes().CreateMany(
        context.Background(),
        models.TemplateIndex,
    ); err != nil {
        log.Fatal("faild initial template database", err)
    }
}
