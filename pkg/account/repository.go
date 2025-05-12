package account

import (
	"github.com/Ratchaphon1412/assistant-llm/pkg/entities"
	"gorm.io/gorm"
)

type Repository interface {
	Create(account *entities.Account) (*entities.Account, error)
	Update(account *entities.Account) (*entities.Account, error)
	Delete(id uint) error
	GetAccountByID(id uint) (*entities.Account, error)
	GetAccountByEmail(email string) (*entities.Account, error)
	GetAllAccounts() ([]entities.Account, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return &repository{
		db: db,
	}
}

func (r *repository) Create(account *entities.Account) (*entities.Account, error) {
	if err := r.db.Create(account).Error; err != nil {
		return nil, err
	}
	return account, nil
}

func (r *repository) Update(account *entities.Account) (*entities.Account, error) {
	if err := r.db.Save(account).Error; err != nil {
		return nil, err
	}
	return account, nil
}

func (r *repository) Delete(id uint) error {
	if err := r.db.Delete(&entities.Account{}, id).Error; err != nil {
		return err
	}
	return nil
}
func (r *repository) GetAccountByID(id uint) (*entities.Account, error) {
	var account entities.Account
	if err := r.db.First(&account, id).Error; err != nil {
		return nil, err
	}
	return &account, nil
}
func (r *repository) GetAccountByEmail(email string) (*entities.Account, error) {
	var account entities.Account
	if err := r.db.Where("email = ?", email).First(&account).Error; err != nil {
		return nil, err
	}
	return &account, nil
}
func (r *repository) GetAllAccounts() ([]entities.Account, error) {
	var accounts []entities.Account
	if err := r.db.Find(&accounts).Error; err != nil {
		return nil, err
	}
	return accounts, nil
}
