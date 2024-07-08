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

func (s *MessageService) GetListWithAdditions() ([]entity.Message, error) {
	return s.msgRepo.GetListWithAuthors()
}

func (s *MessageService) Get(ID entity.ObjectID) (entity.Message, error) {
	return s.msgRepo.Get(ID)
}

func (s *MessageService) Create(m entity.MessageCreate, authorID entity.ObjectID) (entity.ObjectID, error) {
	msg := entity.NewMessage(m.Text, authorID)
	return s.msgRepo.Create(msg)
}
