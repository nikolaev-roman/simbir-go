package repository

import (
	"context"
	"errors"
	"strings"

	"github.com/google/uuid"
	"github.com/nikolaev-roman/simbir-go/internal/account"
	"github.com/nikolaev-roman/simbir-go/internal/models"
	"gorm.io/gorm"
)

type accountRepo struct {
	db *gorm.DB
}

func NewAccountRepository(db *gorm.DB) account.Repository {
	return &accountRepo{db: db}
}

func (r *accountRepo) Create(ctx context.Context, account *models.Account) (*models.Account, error) {

	a := account

	result := r.db.Create(&a)

	if result.Error != nil {
		if strings.Contains(result.Error.Error(), "duplicate key") {
			return nil, errors.New("Account with that username already exists")
		}
		return nil, errors.New("Account with that username already exists")
	}

	return a, nil
}

func (r *accountRepo) GetByID(ctx context.Context, ID uuid.UUID) (*models.Account, error) {
	var account models.Account
	result := r.db.First(&account, "id = ?", ID)
	if result.Error != nil {
		return nil, errors.New("No account found")
	}

	return &account, nil
}

func (r *accountRepo) Update(ctx context.Context, account *models.Account) (*models.Account, error) {
	result := r.db.Save(&account)
	if result.Error != nil {
		return nil, errors.New("No account found")
	}

	return account, nil
}

func (r *accountRepo) GetByUserName(ctx context.Context, account *models.Account) (*models.Account, error) {
	foundAccount := &models.Account{}

	result := r.db.First(&foundAccount, "username = ?", &account.Username)
	if result.Error != nil {
		return nil, errors.New("No account found")
	}

	return foundAccount, nil
}
