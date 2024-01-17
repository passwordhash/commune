package repository

import (
	"go.mongodb.org/mongo-driver/mongo"
)

type Message interface {
}

type Repository struct {
	Message *MessageRepository
}

func NewRepository(db *mongo.Client) *Repository {
	return &Repository{Message: NewMessageMongo(db)}
}
