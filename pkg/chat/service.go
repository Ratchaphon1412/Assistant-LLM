package chat

import (
	"github.com/Ratchaphon1412/assistant-llm/pkg/entities"
)

type Service interface {
	CreateChat(chat *entities.Chat) (*entities.Chat, error)
}

type service struct {
	repository Repository
}

func NewService(repo Repository) Service {
	return &service{
		repository: repo,
	}
}
func (s *service) CreateChat(chat *entities.Chat) (*entities.Chat, error) {
	chat, err := s.repository.Create(chat)
	if err != nil {
		return nil, err
	}
	return chat, nil
}
