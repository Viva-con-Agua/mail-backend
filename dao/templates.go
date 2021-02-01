package dao

import (
	"context"
	"log"
	"mail-backend/models"
	"strings"

	"github.com/Viva-con-Agua/vcago/verr"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)


func handleError(err error) {
	if mongo.IsDuplicateKeyError(err) {
		collection := strings.Split(err.Error(), "collection")
		log.Print(collection)
	}
}

//InsertTemplate dao
func InsertTemplate(ctx context.Context, create *models.TemplateCreate) (*models.Template, *verr.APIError) {
	insert := create.Insert()
	if _, err := DB.Collection("email_templates").InsertOne(ctx, insert); err != nil {
		handleError(err)
		if strings.Contains(err.Error(), "duplicate key") {
			return nil, verr.NewAPIError(err).Conflict("email_template_duplicate")
		}
		return nil, verr.NewAPIError(err).InternalServerError()
	}
	return insert, nil
}
		
	
//ListTemplate returns all application from database
func ListTemplate(ctx context.Context) ( list []models.Template, apiErr *verr.APIError) {
	cursor, err := DB.Collection("email_templates").Find(ctx, bson.M{})
	if err != nil {
		return nil, verr.NewAPIError(err).InternalServerError()
	}
	if err = cursor.All(ctx, &list); err != nil {
		return nil, verr.NewAPIError(err).InternalServerError()
	}
	return list, nil
}

//GetTemplate return application by app_name
func GetTemplate(ctx context.Context, name string) (*models.Template, *verr.APIError) {
	result := new(models.Template)
	err := DB.Collection("email_templates").FindOne(ctx, bson.M{"app_name": name}).Decode(&result)
	if err != nil {
		if strings.Contains(err.Error(), "no documents in result") {
			return nil, verr.NewAPIError(err).NotFound("application_not_found")
		}
		return nil, verr.NewAPIError(err).InternalServerError()
	}
	return result, nil
}

//UpdateTemplate updates an given app in database.
func UpdateTemplate(ctx context.Context, update *models.Template) (*models.Template, *verr.APIError) {
	update.Modified = *update.Modified.Update()
	_, err := DB.Collection("email_templates").UpdateOne(
		ctx,
		bson.M{"_id": update.ID},
		bson.M{"$set": update},
	)
	if err != nil {
		return nil, verr.NewAPIError(err).InternalServerError()
	}
	return update, nil
}

//DeleteTemplate deletes a given application
func DeleteTemplate(ctx context.Context, delete *models.Template) *verr.APIError {
	_, err := DB.Collection("email_templates").DeleteOne(ctx, bson.M{"_id": delete.ID})
	if err != nil {
		return verr.NewAPIError(err).InternalServerError()
	}
	return nil
}	
