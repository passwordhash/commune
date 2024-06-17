package entity

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type User struct {
	ID       ObjectID `bson:"_id"`
	Nickname string   `bson:"nickname"`
	// TODO validation
	Passcode  string             `bson:"passcode"`
	CreatedAt primitive.DateTime `bson:"createdAt"`
}

func (u *User) IsEmpty() bool {
	return u.ID == "" || u.Passcode == ""
}

type UserCreate struct {
	Nickname string `json:"nickname" binding:"required"`
	Email    string `json:"email" binding:"required"`
}

type UserAuth struct {
	Nickname string `json:"nickname" binding:"required"`
	Passcode string `json:"passcode" binding:"required"`
}

type UserResponse struct {
	ID        ObjectID  `json:"id"`
	Nickname  string    `json:"nickname"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"createdAt"`
}

type JWTToken string
