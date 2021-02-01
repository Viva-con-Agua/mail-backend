package models

import (
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"	
	"github.com/google/uuid"
	"github.com/Viva-con-Agua/vcago/vmod"
)

type (
	//TemplateCreate represents a Email Template for storing in database
	TemplateCreate struct {
		Name        string   `bson:"name" json:"name" validate:"required"`
		Scope 	  string `bson:"scope" json:"scope" validate:"required"`
		Case string `bson:"case" json:"case" validate:"required"`
		Subject string `bson:"subject" json:"subject" validate:"required"`
		HTML        string   `bson:"html" json:"html" validate:"required,html"`
		Tags        []string `bson:"tags" json:"tags" validate:"required"`
		Description string   `bson:"description" json:"description" validate:"required"`
		Type string `bson:"type" json:"type" validate:"required"`

	}
	//Template represents a Email Template for storing in database.
	Template struct {
		ID          string   `bson:"_id" json:"id" validate:"required"`
		*TemplateCreate
		Modified vmod.Modified `bson:"modified" json:"modified"`
	}
)

var TemplateIndex = []mongo.IndexModel{
	{
		Keys: bson.D{{Key: "name", Value: 1}},
		Options: options.Index().SetUnique(true),
	},
}

func (t *TemplateCreate) Insert() *Template {
	return &Template{
		uuid.New().String(),
		t,
		*vmod.NewModified(),
	}
} 
