package service

import (
	"commune/internal/entity"
	"commune/internal/repository"
	"time"
)

type User interface {
	// TODO: return jwt
	SignUp(u entity.UserCreate) (entity.JWTToken, error)
	GetById(id entity.ObjectID) (entity.User, error)
	GetAll() ([]entity.User, error)

	Authenticate(credentials entity.UserAuth) (entity.JWTToken, error)
	GenerateToken(u entity.User) (entity.JWTToken, error)
	ParseToken(token entity.JWTToken) (entity.ObjectID, error)
}

type Message interface {
	Get(ID entity.ObjectID) (entity.Message, error)
	GetList() []entity.Message
	Create(m entity.Message) (entity.ObjectID, error)
}

type Service struct {
	User
	Message
}

type Deps struct {
	PassphraseSalt string
	AccessTokenTTL time.Duration
	SigingKey      string
}

func NewService(repos *repository.Repository, deps Deps) *Service {
	return &Service{
		User:    NewUserService(repos.User, deps),
		Message: NewMessageService(repos.Message)}
}
