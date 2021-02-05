package dao

import (
	"context"
	"mail-backend/models"

	"github.com/Viva-con-Agua/vcago/verr"
	"go.mongodb.org/mongo-driver/bson"
)

//InsertTemplate dao
func InsertTemplate(ctx context.Context, create *models.TemplateCreate) (*models.Template, error) {
	insert := create.Insert()
	if _, err := DB.Collection("templates").InsertOne(ctx, insert); err != nil {
		return nil, verr.MongoInsertOneError(ctx, err, "templates")
	}
	return insert, nil
}

//ListTemplate returns all application from database
func ListTemplate(ctx context.Context) (list []models.Template, apiErr error) {
	cursor, err := DB.Collection("templates").Find(ctx, bson.M{})
	if err != nil {
		return nil, verr.InternalServerError(ctx, err)
	}
	if err = cursor.All(ctx, &list); err != nil {
		return nil, verr.InternalServerError(ctx, err)
	}
	return list, nil
}

//GetTemplate return application by app_name
func GetTemplate(ctx context.Context, filter bson.M) (*models.Template, error) {
	result := new(models.Template)
	err := DB.Collection("templates").FindOne(ctx, filter).Decode(&result)
	if err != nil {
		return nil, verr.MongoFindOneError(ctx, err, "templates")
	}
	return result, nil
}

//UpdateTemplate updates an given app in database.
func UpdateTemplate(ctx context.Context, update *models.Template) (err error) {
	update.Modified = *update.Modified.Update()
	result, err := DB.Collection("templates").UpdateOne(
		ctx,
		bson.M{"_id": update.ID},
		bson.M{"$set": update},
	)
	if err = verr.MongoUpdateOneError(ctx, err, "templates", result); err != nil {
		return err
	}
	return nil
}

//DeleteTemplate deletes a given application
func DeleteTemplate(ctx context.Context, delete *models.Template) (err error){
	result, err := DB.Collection("templates").DeleteOne(ctx, bson.M{"_id": delete.ID})
	if err = verr.MongoDeleteOneError(ctx, err, "templates", result); err != nil {
		return err
	}
	return nil
}
