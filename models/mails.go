package models

import (
	"github.com/Viva-con-Agua/vcago/vmod"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type (
	//EmailAddressCreate represents the create model for EmailAdress
	EmailAddressCreate struct {
		Email string   `json:"email" bson:"email" validate:"required,email"`
		Tags  []string `json:"cases" bson:"cases" validate:"required"`
	}
	//EmailAddress represents the email address in database
	EmailAddress struct {
		ID       string        `json:"id" bson:"_id" validate:"required"`
		Email    string        `json:"email" bson:"email" validate:"required,email"`
		Tags     []string      `json:"scope" bson:"scope" validate:"required"`
		Modified vmod.Modified `json:"modified" bson:"modified" validate:"required"`
	}
)
//EmailAddressesIndex contains all database indexes for email_addresses collection
var EmailAddressesIndex = []mongo.IndexModel{
	{
		Keys:    bson.D{{Key: "scope", Value: 1}, {Key: "case", Value: 1}},
		Options: options.Index().SetUnique(true),
	},
}

//Insert creates EmailAddress from EmailAddressCreate for insert into the database
func (cr *EmailAddressCreate) Insert() *EmailAddress {
	return &EmailAddress{
		ID:       uuid.New().String(),
		Email:    cr.Email,
		Tags:     cr.Tags,
		Modified: *vmod.NewModified(),
	}
}
