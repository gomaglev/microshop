package injector

import (
	"context"

	imongo "github.com/gomaglev/microshop/internal/app/model/mongo"

	"go.mongodb.org/mongo-driver/mongo"
)

// InitMongo initiate Mongo storage
func InitMongo() (*mongo.Client, func(), error) {
	client, cleanFunc, err := imongo.NewClient()
	if err != nil {
		return nil, cleanFunc, err
	}

	err = imongo.CreateIndexes(context.Background(), client)
	if err != nil {
		return nil, cleanFunc, err
	}
	return client, cleanFunc, nil
}
