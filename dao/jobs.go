package dao

import (
	"context"
	"mail-backend/models"

	"github.com/Viva-con-Agua/vcago/verr"
	"go.mongodb.org/mongo-driver/bson"
	
)
//InsertJob dao
func InsertJob(ctx context.Context, create *models.JobCreate) (*models.Job, *verr.APIError) {
	insert := create.Insert()
	if _, err := DB.Collection("jobs").InsertOne(ctx, insert); err != nil {
		return nil, verr.MongoHandleError(err)
	}
	return insert, nil
}
		
	
//ListJobs returns all Jobs from database
func ListJobs(ctx context.Context) ( list []models.Job, apiErr *verr.APIError) {
	cursor, err := DB.Collection("jobs").Find(ctx, bson.M{})
	if err != nil {
		return nil, verr.MongoHandleError(err)
	}
	if err = cursor.All(ctx, &list); err != nil {
		return nil, verr.MongoHandleError(err)
	}
	return list, nil
}

//GetJob return Job by filter
func GetJob(ctx context.Context, filter bson.M) (*models.Job, *verr.APIError) {
	result := new(models.Job)
	err := DB.Collection("jobs").FindOne(ctx, filter).Decode(&result)
	if err != nil {
		return nil, verr.MongoHandleError(verr.NewMongoCollError(err, "jobs"))
	}
	return result, nil
}

//UpdateJob updates an Job in the mail database.
func UpdateJob(ctx context.Context, update *models.Job) (*models.Job, *verr.APIError) {
	update.Modified = *update.Modified.Update()
	_, err := DB.Collection("jobs").UpdateOne(
		ctx,
		bson.M{"_id": update.ID},
		bson.M{"$set": update},
	)
	if err != nil {
		return nil, verr.MongoHandleError(err)
	}
	return update, nil
}

//DeleteJob deletes Job in the mail database
func DeleteJob(ctx context.Context, delete *models.Job) *verr.APIError {
	_, err := DB.Collection("jobs").DeleteOne(ctx, bson.M{"_id": delete.ID})
if err != nil {
	return verr.MongoHandleError(err)
	}
	return nil
}

//GetJobWithSubs return JobWithSubs
func GetJobWithSubs(ctx context.Context, filter bson.M) (*models.JobWithSubs, *verr.APIError) {
	job, apiErr := GetJob(ctx, filter)
	if apiErr != nil {
		return nil, apiErr
	}
	email, apiErr := GetEmailAddress(ctx, bson.M{"_id": job.EmailID})
	if apiErr != nil {
		return nil, apiErr
	}
	template, apiErr := GetTemplate(ctx, bson.M{"_id": job.TemplateID})
	if apiErr != nil {
		return nil, apiErr
	}
	return job.JobWithSubs(template, email), nil
}
