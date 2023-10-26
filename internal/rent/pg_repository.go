package rent

import (
	"context"

	"github.com/google/uuid"
	"github.com/nikolaev-roman/simbir-go/internal/models"
)

type Repository interface {
	Create(ctx context.Context, rent *models.Rent) (*models.Rent, error)
	Update(ctx context.Context, rent *models.Rent) (*models.Rent, error)
	GetByID(ctx context.Context, ID uuid.UUID) (*models.Rent, error)
	Delete(ctx context.Context, ID uuid.UUID) error
}
