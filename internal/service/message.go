package service

import "commune/internal/repository"

type MessageService struct {
	msgRepo repository.Message
}

func NewMessageService(msgRepo repository.Message) *MessageService {
	return &MessageService{msgRepo: msgRepo}
}
