package message

type Service struct {
	msgRepo Repository
}

func NewService(msgRepo Repository) *Service {
	return &Service{msgRepo: msgRepo}
}
