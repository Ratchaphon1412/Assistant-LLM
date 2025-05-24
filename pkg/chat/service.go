package chat

import (
	"context"

	"github.com/Ratchaphon1412/assistant-llm/cmd/driver/temporal"
	"github.com/Ratchaphon1412/assistant-llm/configs"
	"github.com/Ratchaphon1412/assistant-llm/pkg/entities"
)

type Service interface {
	CreateChat(chat *entities.Chat) (*entities.Chat, error)
	TriggerAIWorkflow(ctx context.Context, cfg *configs.Config, chat *entities.Chat, redisChanel string, question string) (*entities.Chat, error)
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

func (s *service) TriggerAIWorkflow(ctx context.Context, cfg *configs.Config, chat *entities.Chat, redisChanel string, question string) (*entities.Chat, error) {
	err := temporal.AIWorkflow(ctx, cfg, chat.ID, redisChanel, chat.WorkflowID, question)
	if err != nil {
		return nil, err
	}
	return chat, nil
}
