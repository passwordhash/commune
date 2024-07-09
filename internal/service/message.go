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

func (s *MessageService) GetWithAuthor(ID entity.ObjectID) (entity.Message, error) {
	return s.msgRepo.GetWithAuthor(ID)
}

func (s *MessageService) Get(ID entity.ObjectID) (entity.Message, error) {
	return s.msgRepo.Get(ID)
}

// TODO: shit code
func (s *MessageService) Create(m entity.MessageCreate, authorID entity.ObjectID) (entity.Message, error) {
	msg := entity.NewMessage(m.Text, authorID)
	_, err := s.msgRepo.Create(msg)
	if err != nil {
		return msg, err
	}

	// добавляем автора к новому сообщению
	msgTmp, err := s.msgRepo.GetWithAuthor(msg.ID)
	if err != nil {
		return msg, err
	}
	msg.Author = msgTmp.Author
	msg.AuthorID = ""

	return msg, nil
}
