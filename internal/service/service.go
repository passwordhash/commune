package service

import (
	"commune/internal/entity"
	"commune/internal/repository"
	"time"
)

type User interface {
	SignUp(u entity.UserCreate) (entity.JWTToken, entity.Passcode, error)
	GeneratePasscode() (entity.Passcode, error)
	UpdatePasscode(u entity.UserCreate) (entity.Passcode, error)
	GetById(id entity.ObjectID) (entity.User, error)
	GetAll() ([]entity.User, error)

	Authenticate(credentials entity.UserAuth) (entity.AuthData, error)
	GenerateToken(u entity.User) (entity.JWTToken, error)
	ParseToken(token entity.JWTToken) (entity.ObjectID, error)
}

type Message interface {
	Get(ID entity.ObjectID) (entity.Message, error)
	GetList() []entity.Message
	Create(m entity.MessageCreate) (entity.ObjectID, error)
}

type Email interface {
	SendCode(to string, passcode entity.Passcode) error
}

type Service struct {
	User
	Message
	Email
}

type Deps struct {
	PasscodeSalt   string
	AccessTokenTTL time.Duration
	SigingKey      string

	EmailDeps
}

type EmailDeps struct {
	SmtpHost string
	SmtpPort int

	From     string
	Password string
}

func NewService(repos *repository.Repository, deps Deps) *Service {
	return &Service{
		User:    NewUserService(repos.User, deps),
		Message: NewMessageService(repos.Message),
		Email:   NewEmailService(deps.EmailDeps),
	}
}
