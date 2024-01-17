package service

import (
	"commune/internal/repository"
)

type Message interface {
}

type Service struct {
	Message
}

func NewService(repos *repository.Repository) *Service {
	return &Service{Message: NewMessageService(repos.Message)}
}
