package repository

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func NewMongoDB(uri string) (*mongo.Database, error) {
	//ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	ctx := context.TODO()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
	if err != nil {
		return nil, err
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		return nil, err
	}

	// TODO db name in config
	db := client.Database("commune")

	return db, nil
}
