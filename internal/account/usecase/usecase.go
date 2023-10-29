package usecase

import (
	"context"
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
	"github.com/nikolaev-roman/simbir-go/config"
	"github.com/nikolaev-roman/simbir-go/internal/account"
	"github.com/nikolaev-roman/simbir-go/internal/models"
	"golang.org/x/crypto/bcrypt"
)

type accountUC struct {
	cfg         *config.Config
	accountRepo account.Repository
}

func NewAccountUseCase(cfg *config.Config, accountRepo account.Repository) account.UseCase {
	return &accountUC{cfg: cfg, accountRepo: accountRepo}
}

func (u *accountUC) Register(ctx context.Context, account *models.Account) (*models.Account, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(account.Password), 10)

	if err != nil {
		return nil, errors.New("Failed to hash password")
	}

	account.Password = string(hash)
	account.Balance = 0
	account.IsAdmin = false

	return u.accountRepo.Create(ctx, account)
}

func (u *accountUC) SignIn(ctx context.Context, account *models.Account) (string, error) {

	foundAccount, err := u.accountRepo.GetByUserName(ctx, account)
	if err != nil {
		return "", errors.New("Invalid username or password")
	}

	err = bcrypt.CompareHashAndPassword([]byte(foundAccount.Password), []byte(account.Password))
	if err != nil {
		return "", errors.New("Invalid username or password")
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": foundAccount.ID,
		"exp": time.Now().Add(time.Hour * 24).Unix(),
	})

	tokenString, err := token.SignedString([]byte(u.cfg.Server.JwtSecretKey))

	if err != nil {
		return "", errors.New("Failed to build token")
	}

	return tokenString, nil
}

func (u *accountUC) Create(ctx context.Context, account *models.Account) (*models.Account, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(account.Password), 10)
	if err != nil {
		return nil, errors.New("failed to hash password")
	}

	account.Password = string(hash)

	return u.accountRepo.Create(ctx, account)
}

func (u *accountUC) Update(ctx context.Context, account *models.Account) (*models.Account, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(account.Password), 10)
	if err != nil {
		return nil, errors.New("failed to hash password")
	}

	account.Password = string(hash)

	return u.accountRepo.Update(ctx, account)
}

func (u *accountUC) Delete(ctx context.Context, accountID uuid.UUID) error {
	return u.accountRepo.Delete(ctx, accountID)
}

func (u *accountUC) GetByID(ctx context.Context, ID uuid.UUID) (*models.Account, error) {
	return u.accountRepo.GetByID(ctx, ID)
}

func (u *accountUC) Search(ctx context.Context, params models.AccountSearchParams) ([]*models.Account, error) {
	return u.accountRepo.Search(ctx, params)
}
