package model

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

type Endpoint struct {
	ID         primitive.ObjectID `bson:"_id,omitempty" json:"_id"`
	Name       string             `bson:"name" json:"name"`
	Kind       string             `bson:"kind" json:"kind"`
	Schedule   *EndpointSchedule  `bson:"schedule" json:"schedule"`
	Emqx       *EndpointEmqx      `bson:"emqx" json:"emqx"`
	CreateTime time.Time          `bson:"create_time" json:"create_time"`
	UpdateTime time.Time          `bson:"update_time" json:"update_time"`
}

type EndpointSchedule struct {
	Node string `bson:"node" json:"node"`
}

type EndpointEmqx struct {
	Host      string `bson:"host" json:"host"`
	ApiKey    string `bson:"api_key" json:"api_key"`
	SecretKey string `bson:"secret_key" json:"secret_key"`
}

func SetEndpoints(ctx context.Context, db *mongo.Database) (err error) {
	var ns []string
	if ns, err = db.ListCollectionNames(ctx, bson.M{"name": "endpoints"}); err != nil {
		return
	}
	var jsonSchema bson.D
	if err = LoadJsonSchema("endpoint", &jsonSchema); err != nil {
		return
	}
	if len(ns) == 0 {
		option := options.CreateCollection().SetValidator(jsonSchema)
		if err = db.CreateCollection(ctx, "endpoints", option); err != nil {
			return
		}
	} else {
		if err = db.RunCommand(ctx, bson.D{
			{"collMod", "endpoints"},
			{"validator", jsonSchema},
			{"validationLevel", "strict"},
		}).Err(); err != nil {
			return
		}
	}
	return
}
