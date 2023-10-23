package account

import (
	"context"

	"github.com/google/uuid"
	"github.com/nikolaev-roman/simbir-go/internal/models"
)

type UseCase interface {
	Register(ctx context.Context, account *models.Account) (*models.Account, error)
	SignIn(ctx context.Context, account *models.Account) (string, error)
	Update(ctx context.Context, account *models.Account) (*models.Account, error)
	GetByID(ctx context.Context, ID uuid.UUID) (*models.Account, error)
}
