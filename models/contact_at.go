package models

import (
	"github.com/Viva-con-Agua/vcago/vmod"
	"github.com/google/uuid"
)

type (
	ContactAtCreate struct {
		FirstName string `json:"first_name" validate:"required"`
		LastName  string `json:"last_name" validate:"required"`
		Message   string `json:"message" validate:"required"`
		Email     string `json:"email"`
		KnownFrom string `json:"known_from"`
	}

	ContactAt struct {
		ID        string        `bson:"_id" json:"id"`
		FirstName string        `bson:"first_name" json:"first_name" validate:"required"`
		LastName  string        `bson:"last_name" json:"last_name" validate:"required"`
		Message   string        `bson:"message" json:"message" validate:"required"`
		Email     string        `bson:"email" json:"email"`
		KnownFrom string        `bson:"known_from" json:"known_from"`
		Modified  vmod.Modified `bson:"modified" json:"modified"`
	}
)

func (cc *ContactAtCreate) Insert() *ContactAt {
	return &ContactAt{
		ID:        uuid.New().String(),
		FirstName: cc.FirstName,
		LastName:  cc.LastName,
		Message:   cc.Message,
		Email:     cc.Email,
		KnownFrom: cc.KnownFrom,
		Modified:  *vmod.NewModified(),
	}
}

func (c *ContactAt) Update() *ContactAt {
	c.Modified = *c.Modified.Update()
	return c
}
