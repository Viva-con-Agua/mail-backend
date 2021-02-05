package dao

import (
	"context"
	"mail-backend/models"

	"github.com/Viva-con-Agua/vcago/verr"
	"go.mongodb.org/mongo-driver/bson"
)

//InsertJob dao
func InsertJob(ctx context.Context, create *models.JobCreate) (*models.Job, error) {
	insert := create.Insert()
	if _, err := DB.Collection("jobs").InsertOne(ctx, insert); err != nil {
		return nil, verr.MongoInsertOneError(ctx, err, "jobs")
	}
	return insert, nil
}
		
	
//ListJobs returns all Jobs from database
func ListJobs(ctx context.Context) ( list []models.Job, err error) {
	cursor, err := DB.Collection("jobs").Find(ctx, bson.M{})
	if err != nil {
		return nil, verr.InternalServerError(ctx, err)
	}
	if err = cursor.All(ctx, &list); err != nil {
		return nil, verr.InternalServerError(ctx, err)
	}
	return list, nil
}

//GetJob return Job by filter
func GetJob(ctx context.Context, filter bson.M) (*models.Job, error) {
	result := new(models.Job)
	err := DB.Collection("jobs").FindOne(ctx, filter).Decode(&result)
	if err != nil {
		return nil, verr.MongoFindOneError(ctx, err, "jobs")
	}
	return result, nil
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
	}
	template, err := GetTemplate(ctx, bson.M{"_id": job.TemplateID})
	if err != nil {
		return nil, err
	}
	return job.JobWithSubs(template, email), nil
}
