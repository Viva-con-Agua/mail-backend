package models

import (
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"	
)

type (
	//Template represents a Email Template for storing in database
	Template struct {
		Case 	  string `bson:"case" json:"case" validate:"required"`
		Subject string `bson:"subject" json:"subject" validate:"required"`
		HTML        string   `bson:"html" json:"html" validate:"required,html"`
		Country string `bson:"country" json:"country" validate:"required"`
	}
	//Template represents a Email Template for storing in database.
	//Template struct {
	//	ID          string   `bson:"_id" json:"id" validate:"required"`
	//	*TemplateCreate
	//	Modified vmod.Modified `bson:"modified" json:"modified"`
	//}
)

var TemplateIndex = []mongo.IndexModel{
	{
		Keys: bson.D{{Key: "name", Value: 1}},
		Options: options.Index().SetUnique(true),
	},
}

//GetTemplate return template for case
func GetTemplate(list []Template, country string) *Template {
	for i := range list {
		if list[i].Country == country {
			return &list[i]
		}
	}
	for y := range list {
		if list[y].Case == "default" {
			return &list[y]
		}
	}
	return nil
}
/*
func (t *TemplateCreate) Insert() *Template {
	return &Template{
		uuid.New().String(),
		t,
		*vmod.NewModified(),
	}
} */
