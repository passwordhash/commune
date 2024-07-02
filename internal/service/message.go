package service

import (
	"commune/internal/entity"
	"commune/internal/repository"
)

type MessageService struct {
	msgRepo repository.Message
}

func NewMessageService(msgRepo repository.Message) *MessageService {
	return &MessageService{msgRepo: msgRepo}
}

func (s *MessageService) GetList() []entity.Message {
	return s.msgRepo.GetList()
}

func (s *MessageService) Get(ID entity.ObjectID) (entity.Message, error) {
	return s.msgRepo.Get(ID)
}

func (s *MessageService) Create(m entity.MessageCreate) (entity.ObjectID, error) {
	msg := entity.NewMessage(m.Text, m.AuthorID)
	return s.msgRepo.Create(msg)
}
