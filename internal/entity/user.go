package entity

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type User struct {
	ID   ObjectID `bson:"_id"`
	Name string   `bson:"name"`
	// TODO validation
	Passphrase string             `bson:"passphrase"`
	CreatedAt  primitive.DateTime `bson:"createdAt"`
}

func (u *User) IsEmpty() bool {
	return u.ID == "" || u.Passphrase == ""
}

type UserCreate struct {
	Name       string `json:"name" binding:"required"`
	Passphrase string `json:"passphrase" binding:"required"`
}

type UserAuth struct {
	Name       string `json:"name" binding:"required"`
	Passphrase string `json:"passphrase" binding:"required"`
}

type UserResponse struct {
	ID         ObjectID  `json:"id"`
	Name       string    `json:"name"`
	Passphrase string    `json:"passphrase"`
	CreatedAt  time.Time `json:"createdAt"`
}

type JWTToken string
