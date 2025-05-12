package account

import (
	"github.com/Ratchaphon1412/assistant-llm/pkg/entities"
)

type Service interface {
	CreateAccount(account *entities.Account) (*entities.Account, error)
	UpdateAccount(account *entities.Account) (*entities.Account, error)
	DeleteAccount(id uint) error
	GetAccountByID(id uint) (*entities.Account, error)
	GetAccountByEmail(email string) (*entities.Account, error)
	GetAllAccounts() ([]entities.Account, error)
}

type service struct {
	repository Repository
}

func NewService(repo Repository) Service {
	return &service{
		repository: repo,
	}
}

func (s *service) CreateAccount(account *entities.Account) (*entities.Account, error) {
	return s.repository.Create(account)
}
func (s *service) UpdateAccount(account *entities.Account) (*entities.Account, error) {
	return s.repository.Update(account)
}
func (s *service) DeleteAccount(id uint) error {
	return s.repository.Delete(id)
}
func (s *service) GetAccountByID(id uint) (*entities.Account, error) {
	return s.repository.GetAccountByID(id)
}
func (s *service) GetAccountByEmail(email string) (*entities.Account, error) {
	return s.repository.GetAccountByEmail(email)
}
func (s *service) GetAllAccounts() ([]entities.Account, error) {
	return s.repository.GetAllAccounts()
}
