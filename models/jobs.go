package models

import (
	//"mail-backend/mail"

	"github.com/Viva-con-Agua/vcago/vmod"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)
type (

	//JobCreate used for create Job
	JobCreate struct {
		Name      string `json:"name" bson:"name" validate:"required"`
		Scope     string `json:"scope" bson:"scope" validate:"required"`
		Case      string `json:"case" bson:"case" validate:"required"`
		TemplateID string `json:"template_id" bson:"template_id" validate:"required"`
		Country   string `json:"country"`
		EmailID   string `json:"email_id" bson:"email_id" validate:"required"`
	}

	//Job represent a mailing job into database
	Job struct {
		ID        string        `json:"id" bson:"_id" validate:"required"`
		Name      string `json:"name" bson:"name" validate:"required"`
		Scope     string `json:"scope" bson:"scope" validate:"required"`
		Case      string `json:"case" bson:"case" validate:"required"`
		Templates map[string]string `json:"templates" bson:"templates" validate:"required"`
		EmailID   string `json:"email_id" bson:"email_id" validate:"required"`
		Modified  vmod.Modified `json:"modified" bson:"modified" validate:"required"`
	}
	//JobWithSubs Job model with Template and EmailAdress model instead of _id's
	JobWithSubs struct {
		ID        string        `json:"id" bson:"_id" validate:"required"`
		Name      string        `json:"name" bson:"name" validate:"required"`
        Scope     string        `json:"scope" bson:"scope" validate:"required"`
		Case      string        `json:"case" bson:"case" validate:"required"`
		Template map[string]Template `json:"template"`
		Email EmailAddress `json:"email_address"`
	}
)

//JobsColl ist the database collection name for the models.Job model.
const JobsColl = "jobs"

//JobsIndex contains all database indexes for jobs collection
var JobsIndex = []mongo.IndexModel{
	{
		Keys:    bson.D{{Key: "scope", Value: 1}, {Key: "case", Value: 1}},
		Options: options.Index().SetUnique(true),
	},
}

//Insert creates Job for inserting into database
func (cr *JobCreate) Insert() *Job {
	templates := make(map[string]string)
	templates["default"] = cr.TemplateID
	templates[cr.Country] = cr.TemplateID
	return &Job{
		uuid.New().String(),
		cr.Name,
		cr.Scope,
		cr.Case,
		templates,
		cr.EmailID,
		*vmod.NewModified(),
	}
}
//JobWithSubs converts Job into JobWithSubs. Need Template and EmailAddress for provide.
func (j *Job) JobWithSubs(template map[string]Template, email *EmailAddress) *JobWithSubs{
	return &JobWithSubs{
		ID: j.ID,
		Name: j.Name,
		Scope: j.Scope,
		Case: j.Case,
		Template: template,
		Email: *email,
	}
}












