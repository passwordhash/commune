package repository

import (
	"commune/internal/entity"
	"go.mongodb.org/mongo-driver/mongo"
)

const (
	collectionUser    = "user"
	collectionMessage = "message"
)

type User interface {
	Create(u entity.User) (entity.ObjectID, error)
	GetById(id entity.ObjectID) (entity.User, error)
	GetByEmail(email string) (entity.User, error)
	UpdatePasscode(email string, newPasscodeHash string) error
	GetAll() ([]entity.User, error)
}

type Message interface {
	Get(ID entity.ObjectID) (entity.Message, error)
	GetWithAuthor(ID entity.ObjectID) (entity.Message, error)
	GetList() []entity.Message
	GetListWithAuthors() ([]entity.Message, error)
	Create(m entity.Message) (entity.ObjectID, error)
}

type Repository struct {
	User    *UserRepository
	Message *MessageRepository
}

func NewRepository(db *mongo.Database) *Repository {
	return &Repository{
		User:    NewUserMongo(db),
		Message: NewMessageMongo(db)}
}

func userCollection(db *mongo.Database) *mongo.Collection {
	return db.Collection(collectionUser)
}

func messageCollection(db *mongo.Database) *mongo.Collection {
	return db.Collection(collectionMessage)
}
