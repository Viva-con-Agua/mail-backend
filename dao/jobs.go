package dao

import (
	"context"
	"mail-backend/models"

	"github.com/Viva-con-Agua/vcago/verr"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

const coll = models.JobsColl

//InsertJob dao
func InsertJob(ctx context.Context, insert *models.Job) (err error) {
	if _, err := DB.Collection(coll).InsertOne(ctx, insert); err != nil {
		return verr.MongoInsertOneError(ctx, err, coll)
	}
	return nil
}
		
	
//ListJobs returns all Jobs from database
func ListJobs(ctx context.Context, filter bson.M) ( result []models.Job, err error) {
	var cursor *mongo.Cursor
	if cursor, err = DB.Collection(coll).Find(ctx, filter); err != nil {
		return nil, verr.InternalServerError(ctx, err)
	}
	if err = cursor.All(ctx, &result); err != nil {
		return nil, verr.InternalServerError(ctx, err)
	}
	return
}

//GetJob return Job by filter
func GetJob(ctx context.Context, filter bson.M) (result *models.Job, err error) {
	if err = DB.Collection(coll).FindOne(ctx, filter).Decode(&result); err != nil {
		return nil, verr.MongoFindOneError(ctx, err, coll)
	}
	return
}

//UpdateJob updates an Job in the mail database.
func UpdateJob(ctx context.Context, update *models.Job) (err error) {
	update.Modified = *update.Modified.Update()
	result, err := DB.Collection("jobs").UpdateOne(
		ctx,
		bson.M{"_id": update.ID},
		bson.M{"$set": update},
	)
	if err = verr.MongoUpdateOneError(ctx, err, "jobs", result); err != nil {
		return
	}
	return nil
}

//DeleteJob deletes Job in the mail database
func DeleteJob(ctx context.Context, delete *models.Job) (err error) {
	result, err := DB.Collection("jobs").DeleteOne(ctx, bson.M{"_id": delete.ID})
	if err = verr.MongoDeleteOneError(ctx, err, "jobs", result); err != nil {
		return err
	}
	return nil
}

//GetJobWithSubs return JobWithSubs
func GetJobWithSubs(ctx context.Context, filter bson.M) (*models.JobWithSubs, error) {
	job, err := GetJob(ctx, filter)
	if err != nil {
		return nil, err
	}
	email, err := GetEmailAddress(ctx, bson.M{"_id": job.EmailID})
	if err != nil {
		return nil, err
	}/*
	var templates = make(map[string]models.Template)
	for v := range job.Templates {
		var template *models.Template
		template, err = GetTemplate(ctx, bson.M{"_id": job.Templates[v]})
		if err != nil {
			return nil, err
		}
		templates[v] = *template
	}*/
		
	return job.JobWithSubs(email), nil
}
