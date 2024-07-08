package repository

import (
	"commune/internal/entity"
	"context"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type MessageRepository struct {
	db *mongo.Database
}

func NewMessageMongo(db *mongo.Database) *MessageRepository {
	return &MessageRepository{db: db}
}

func (r *MessageRepository) Get(ID entity.ObjectID) (entity.Message, error) {
	var message entity.Message

	cur, err := messageCollection(r.db).Find(context.TODO(), bson.M{"_id": ID})

	if err != nil {
		logrus.Error(err)
		return message, err
	}

	cur.Next(context.TODO())
	if err := cur.Decode(&message); err != nil {
		logrus.Error(err)
		return message, nil
	}

	return message, nil
}

func (r *MessageRepository) GetList() []entity.Message {
	var list []entity.Message

	cur, err := messageCollection(r.db).Find(context.TODO(), bson.D{})
	if err != nil {
		logrus.Error(err)
		return nil
	}
	defer cur.Close(context.TODO())

	for cur.Next(context.TODO()) {
		var msg entity.Message
		if err := cur.Decode(&msg); err != nil {
			logrus.Error(err)
			return list
		}

		list = append(list, msg)
	}

	return list
}

func (r *MessageRepository) GetListWithAuthors() ([]entity.Message, error) {
	var list []entity.Message

	pipeline := mongo.Pipeline{
		{{"$lookup", bson.D{
			{"from", "user"},
			{"localField", "author_id"},
			{"foreignField", "_id"},
			{"as", "author"},
		}}},
		{{"$unwind", bson.D{
			{"path", "$author"},
			{"preserveNullAndEmptyArrays", true},
		}}},
	}

	cur, err := messageCollection(r.db).Aggregate(context.TODO(), pipeline)
	if err != nil {
		return list, err
	}
	defer cur.Close(context.TODO())

	for cur.Next(context.TODO()) {
		var msg entity.Message
		if err := cur.Decode(&msg); err != nil {
			logrus.Error(err)
			return list, err
		}

		list = append(list, msg)
	}

	return list, cur.Err()
}

func (r *MessageRepository) Create(input entity.Message) (entity.ObjectID, error) {
	if _, err := messageCollection(r.db).InsertOne(context.TODO(), input); err != nil {
		return "", err
	}

	return input.ID, nil
}
