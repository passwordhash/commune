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

func (s *MessageService) Create(m entity.Message) (entity.ObjectID, error) {
	//message := entity.NewMessage(text)
	if m.ID == "" {
		m.ID = entity.NewObjectId()
	}
	return s.msgRepo.Create(m)
}
