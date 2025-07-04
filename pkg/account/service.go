package account

import (
	"context"
	"fmt"
	"time"

	"github.com/Ratchaphon1412/assistant-llm/cmd/driver/auth"
	"github.com/Ratchaphon1412/assistant-llm/configs"
	"github.com/Ratchaphon1412/assistant-llm/pkg/entities"
	"github.com/golang-jwt/jwt/v5"
)

type Service interface {
	CreateAccount(account *entities.Account) (*entities.Account, error)
	UpdateAccount(account *entities.Account) (*entities.Account, error)
	DeleteAccount(id uint) error
	GetAccountByID(id uint) (*entities.Account, error)
	GetAccountByEmail(email string) (*entities.Account, error)
	GetAllAccounts() ([]entities.Account, error)
	SignIn(account *entities.Account, cfg configs.Config) (string, error)
	GoogleSignIn(cfg *configs.Config) (string, error)
	GoogleCallback(ctx context.Context, code string, cfg *configs.Config) (*auth.GoogleResponse, error)
}

type service struct {
	repository Repository
}

func NewService(repo Repository) Service {
	return &service{
		repository: repo,
	}
}

func (s *service) SignIn(account *entities.Account, cfg configs.Config) (string, error) {
	tokenJwt := jwt.New(jwt.SigningMethodHS256)
	claims := tokenJwt.Claims.(jwt.MapClaims)
	claims["email"] = account.Email
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()

	t, err := tokenJwt.SignedString([]byte(cfg.JWT_SECRET))
	if err != nil {
		return "", fmt.Errorf("failed to sign token: %v", err)
	}
	return t, nil
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

func (s *service) GoogleSignIn(cfg *configs.Config) (string, error) {
	path := auth.ConfigGoogle(cfg)
	url := path.AuthCodeURL("state")
	return url, nil
}

func (s *service) GoogleCallback(ctx context.Context, code string, cfg *configs.Config) (*auth.GoogleResponse, error) {
	token, err := auth.ConfigGoogle(cfg).Exchange(ctx, code)
	if err != nil {
		return nil, fmt.Errorf("failed to exchange token: %v", err)
	}
	userinfo := auth.GetUserInfo(token.AccessToken)
	if userinfo.Email == "" {
		return nil, fmt.Errorf("failed to get user info")
	}
	return &userinfo, nil
}
