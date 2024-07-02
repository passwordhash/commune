package entity

import (
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type ObjectID string

func NewObjectId() ObjectID {
	return ObjectID(uuid.NewString())
}

type MessageCreate struct {
	Text     string   `json:"text" binding:"required"`
	AuthorID ObjectID `bson:"author_id" json:"author_id" binding:"required"`
}

type Message struct {
	ID       ObjectID           `bson:"_id" json:"id"`
	Text     string             `bson:"text" json:"text"`
	Date     primitive.DateTime `bson:"date" json:"date"`
	AuthorID ObjectID           `bson:"author_id" json:"author_id"`
}

func NewMessage(text string, authorID ObjectID) Message {
	return Message{
		ID:       NewObjectId(),
		Text:     text,
		Date:     primitive.NewDateTimeFromTime(time.Now()),
		AuthorID: authorID,
	}
}
