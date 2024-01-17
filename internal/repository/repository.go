package repository

import (
	"commune/internal/domain/message"
	"go.mongodb.org/mongo-driver/mongo"
)

type Message interface {
	Get()
}

type Repository struct {
	Message *message.Repository
}

func NewRepository(db *mongo.Client) *Repository {
	return &Repository{Message: message.NewRepository(db)}
}
