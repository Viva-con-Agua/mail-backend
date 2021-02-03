package dao

import (
	"context"
	"mail-backend/models"

	"github.com/Viva-con-Agua/vcago/verr"
	"go.mongodb.org/mongo-driver/bson"
)

//InsertEmailAddress inserts an EmailAddressCreate as EmailAddress into the database
func InsertEmailAddress(ctx context.Context, create *models.EmailAddressCreate) (*models.EmailAddress, *verr.APIError) {
	insert := create.Insert()
	if _, err := DB.Collection("email_addresses").InsertOne(ctx, insert); err != nil {
		return nil, verr.MongoHandleError(err)
	}
	return insert, nil
}

//ListEmailAddress return a list of EmailAddress structs from the database
func ListEmailAddress()([]models.EmailAddress, *verr.APIError) {
    return nil, nil
}

//GetEmailAddress return one EmailAddress struct from the database
func GetEmailAddress(ctx context.Context, filter bson.M)(*models.EmailAddress, *verr.APIError) {
	result := new(models.EmailAddress)
	err := DB.Collection("email_addresses").FindOne(ctx, filter).Decode(&result)
	if err != nil {
		return nil, verr.MongoHandleError(verr.NewMongoCollError(err, "email_addresses"))
	}
    return result, nil
}

//UpdateEmailAddress updates a EmailAddrerss struct in the database
func UpdateEmailAddress()(*models.EmailAddress, *verr.APIError) {
    return nil, nil
}

//DeleteEmailAddress deletes an EmailAddress struct from the database
func DeleteEmailAddress()(*verr.APIError) {
    return nil
}

