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

// TODO: разбить на разные DTO
type Message struct {
	ID   ObjectID           `bson:"_id" json:"id"`
	Text string             `bson:"text" json:"text"`
	Date primitive.DateTime `bson:"date" json:"date"`
}

func NewMessage(text string) Message {
	return Message{
		ID:   NewObjectId(),
		Text: text,
		Date: primitive.NewDateTimeFromTime(time.Now()),
	}
}
