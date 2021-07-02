package dao

import (
	"context"
	"mail-backend/models"

	"github.com/Viva-con-Agua/vcago/verror"
	"go.mongodb.org/mongo-driver/bson"
)

type ContactAtDAO struct {
	ctx    context.Context
	filter bson.M
	query  bson.D
	coll   string
}

func NewContactAtDAO(ctx context.Context) *ContactAtDAO {
	return &ContactAtDAO{
		ctx:  ctx,
		coll: "contract_at",
	}
}

func (d *ContactAtDAO) Coll() string {
	return d.coll
}

func (d *ContactAtDAO) Filter(f bson.M) *ContactAtDAO {
	d.filter = f
	return d
}

func (d *ContactAtDAO) Insert(m *models.ContactAt) (r *models.ContactAt, verr *verror.Error) {
	_, err := DB.Collection(d.coll).InsertOne(d.ctx, m)
	if err != nil {
		return nil, verror.New(err).Mongo(d.coll)
	}
	return m, nil
}

func (d *ContactAtDAO) Get() (r *models.ContactAt, ve *verror.Error) {
	err := DB.Collection(d.coll).FindOne(d.ctx, d.filter).Decode(&r)
	if err != nil {
		return nil, verror.New(err).Mongo(d.coll)
	}
	return
}

func (d *ContactAtDAO) List() (r []models.ContactAt, verr *verror.Error) {
	cursor, err := DB.Collection(d.coll).Aggregate(d.ctx, d.query)
	if err != nil {
		return nil, verror.New(err).Mongo(d.coll)
	}
	if err := cursor.All(d.ctx, &r); err != nil {
		return nil, verror.New(err).Mongo(d.coll)
	}
	return
}

func (d *ContactAtDAO) Update(m *models.ContactAt) (verr *verror.Error) {
	m = m.Update()
	result, err := DB.Collection(d.coll).UpdateOne(
		d.ctx,
		bson.M{"_id": m.ID},
		bson.M{"$set": m},
	)
	if err != nil {
		return verror.New(verror.MongoUpdateErr(err, result)).Mongo(d.coll)
	}
	return
}

func (d *ContactAtDAO) Delete(m *models.ContactAt) (verr *verror.Error) {
	result, err := DB.Collection(d.coll).DeleteOne(d.ctx, bson.M{"_id": m.ID})
	if err != nil {
		return verror.New(verror.MongoDeleteErr(err, result)).Mongo(d.coll)
	}
	return
}
