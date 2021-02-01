package models

import (
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"	
    "github.com/google/uuid"
    "github.com/Viva-con-Agua/vcago/vmod"
)
type (
    //JobCreate used for create Job
    JobCreate struct {
        Name      string        `json:"name" bson:"name" validate:"required"`
        Scope string            `json:"scope" bson:"scope" validate:"required"`
        Case      string        `json:"case" bson:"case" validate:"required"`
		TempateID string        `json:"template_id" bson:"template_id" validate:"required"`
		EmailID   string        `json:"email_id bson:"email_id" validate:"required"`
    }

	//Job represent a mailing job into database
	Job struct {
		ID        string        `json:"id" bson:"_id" validate:"required"`
		Name      string        `json:"name" bson:"name" validate:"required"`
        Case      string        `json:"case" bson:"case" validate:"required"`
		TempateID string        `json:"template_id" bson:"template_id" validate:"required"`
		EmailID   string        `json:"email_id bson:"email_id" validate:"required"`
		Modified  vmod.Modified `json:"modified" bson:"modified" validate:"required"`
	}
)

//JobIndex contains all database indexes
var JobIndex = []mongo.IndexModel{
	{
		Keys: bson.D{{Key: "scope", Value: 1},{Key: "case", Value: 1}},
		Options: options.Index().SetUnique(true),
	},
}

//Insert creates Job for inserting into database
func (cr *JobCreate) Insert() *Job {
    return &Job{
        ID: uuid.New().String(),
        Name: cr.Name,
        Case: cr.Case,
        TempateID: cr.TempateID,
        EmailID: cr.EmailID,
        Modified: *vmod.NewModified(),
    }
}
