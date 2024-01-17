package message

import "go.mongodb.org/mongo-driver/mongo"

type Repository struct {
	db *mongo.Client
}

func NewRepository(db *mongo.Client) *Repository {
	return &Repository{db: db}
}

func Get() {
	return
}
