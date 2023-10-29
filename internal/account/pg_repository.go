package account

import (
	"context"

	"github.com/google/uuid"
	"github.com/nikolaev-roman/simbir-go/internal/models"
)

type Repository interface {
	Create(ctx context.Context, account *models.Account) (*models.Account, error)
	Update(ctx context.Context, account *models.Account) (*models.Account, error)
	Delete(ctx context.Context, accountID uuid.UUID) error
	GetByID(ctx context.Context, ID uuid.UUID) (*models.Account, error)
	GetByUserName(ctx context.Context, account *models.Account) (*models.Account, error)
	Search(ctx context.Context, params models.AccountSearchParams) ([]*models.Account, error)
}
