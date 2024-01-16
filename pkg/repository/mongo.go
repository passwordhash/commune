package repository

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// TODO: move to config
const connectURI = "mongodb://passwordhash:root@localhost:27017/"

func NewMongoDB() (*mongo.Client, error) {
	//ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	ctx := context.TODO()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(connectURI))
	if err != nil {
		return nil, err
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		return nil, err
	}

	return client, nil
}
