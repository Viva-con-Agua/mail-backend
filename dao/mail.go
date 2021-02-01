package dao

import (
    "strings"
    "context"
	"mail-backend/models"

	"github.com/Viva-con-Agua/vcago/verr"
)

//InsertEmailAddress inserts an EmailAddressCreate as EmailAddress into the database
func InsertEmailAddress(ctx context.Context, create *models.EmailAddressCreate) (*models.EmailAddress, *verr.APIError) {
	insert := create.Insert()
	if _, err := DB.Collection("email_addresses").InsertOne(ctx, insert); err != nil {
		if strings.Contains(err.Error(), "duplicate key") {
			return nil, verr.NewAPIError(err).Conflict("email_address_duplicate")
		}
		return nil, verr.NewAPIError(err).InternalServerError()
	}
	return insert, nil
}

//ListEmailAddress return a list of EmailAddress structs from the database
func ListEmailAddress()([]models.EmailAddress, *verr.APIError) {
    return nil, nil
}

//GetEmailAddress return one EmailAddress struct from the database
func GetEmailAddress()(*models.EmailAddress, *verr.APIError) {
    return nil, nil
}

//UpdateEmailAddress updates a EmailAddrerss struct in the database
func UpdateEmailAddress()(*models.EmailAddress, *verr.APIError) {
    return nil, nil
}

//DeleteEmailAddress deletes an EmailAddress struct from the database
func DeleteEmailAddress()(*verr.APIError) {
    return nil
}

