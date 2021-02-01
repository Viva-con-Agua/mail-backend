package dao

import (
	"context"
	"mail-backend/models"

	"github.com/Viva-con-Agua/vcago/verr"
	"go.mongodb.org/mongo-driver/bson"
)

//InsertTemplate dao
func InsertTemplate(ctx context.Context, create *models.TemplateCreate) (*models.Template, *verr.APIError) {
	insert := create.Insert()
	if _, err := DB.Collection("templates").InsertOne(ctx, insert); err != nil {
		return nil, verr.MongoHandleError(err)
	}
	return insert, nil
}
		
	
//ListTemplate returns all application from database
func ListTemplate(ctx context.Context) ( list []models.Template, apiErr *verr.APIError) {
	cursor, err := DB.Collection("templates").Find(ctx, bson.M{})
	if err != nil {
		return nil, verr.MongoHandleError(err)
	}
	if err = cursor.All(ctx, &list); err != nil {
		return nil, verr.MongoHandleError(err)
	}
	return list, nil
}

//GetTemplate return application by app_name
func GetTemplate(ctx context.Context, filter bson.M) (*models.Template, *verr.APIError) {
	result := new(models.Template)
	err := DB.Collection("templates").FindOne(ctx, filter).Decode(&result)
	if err != nil {
		return nil, verr.MongoHandleError(verr.NewMongoCollError(err, "email_templates"))
	}
	return result, nil
}

//UpdateTemplate updates an given app in database.
func UpdateTemplate(ctx context.Context, update *models.Template) (*models.Template, *verr.APIError) {
	update.Modified = *update.Modified.Update()
	_, err := DB.Collection("templates").UpdateOne(
		ctx,
		bson.M{"_id": update.ID},
		bson.M{"$set": update},
	)
	if err != nil {
		return nil, verr.MongoHandleError(err)
	}
	return update, nil
}

//DeleteTemplate deletes a given application
func DeleteTemplate(ctx context.Context, delete *models.Template) *verr.APIError {
	_, err := DB.Collection("templates").DeleteOne(ctx, bson.M{"_id": delete.ID})
if err != nil {
	return verr.MongoHandleError(err)
	}
	return nil
}	
