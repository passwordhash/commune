package repository

import "go.mongodb.org/mongo-driver/mongo"

type MessageRepository struct {
	db *mongo.Client
}

func NewMessageMongo(db *mongo.Client) *MessageRepository {
	return &MessageRepository{db: db}
}
