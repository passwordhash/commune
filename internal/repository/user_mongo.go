package repository

import (
	"commune/internal/entity"
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserRepository struct {
	db    *mongo.Database
	userC *mongo.Collection
}

func NewUserMongo(db *mongo.Database) *UserRepository {
	return &UserRepository{db: db, userC: userCollection(db)}
}

func (r *UserRepository) Create(u entity.User) (entity.ObjectID, error) {
	if _, err := r.userC.InsertOne(context.TODO(), u); err != nil {
		return "", err
	}

	return u.ID, nil
}

func (r *UserRepository) GetById(id entity.ObjectID) (entity.User, error) {
	var user entity.User

	cur := r.userC.FindOne(context.TODO(), bson.M{})
	if cur.Err() != nil {
		return user, cur.Err()
	}

	err := cur.Decode(&user)

	return user, err
}

func (r *UserRepository) GetByPassphrase(p string) (entity.User, error) {
	var user entity.User

	cur := r.userC.FindOne(context.TODO(), bson.M{"passphrase": p})
	if cur.Err() != nil {
		return user, cur.Err()
	}

	err := cur.Decode(&user)

	return user, err
}

// TODO протестировать декодирование без курсора
func (r *UserRepository) GetAll() ([]entity.User, error) {
	var list []entity.User

	cur, err := r.userC.Find(context.TODO(), bson.M{})
	if err != nil {
		return list, cur.Err()
	}

	for cur.Next(context.TODO()) {
		var user entity.User

		if cur.Err() != nil {
			return list, cur.Err()
		}

		if err := cur.Decode(&user); err != nil {
			return list, err
		}

		list = append(list, user)
	}

	return list, err
}
