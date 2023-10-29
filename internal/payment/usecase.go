package payment

import (
	"context"

	"github.com/google/uuid"
	"github.com/nikolaev-roman/simbir-go/internal/models"
)

type UseCase interface {
	EnrichBalance(ctx context.Context, accountID uuid.UUID) (*models.Account, error)
}
