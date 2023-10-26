package transport

import (
	"context"

	"github.com/google/uuid"
	"github.com/nikolaev-roman/simbir-go/internal/models"
)

type Repository interface {
	Create(ctx context.Context, transport *models.Transport) (*models.Transport, error)
	Update(ctx context.Context, transport *models.Transport) (*models.Transport, error)
	GetByID(ctx context.Context, ID uuid.UUID) (*models.Transport, error)
	Delete(ctx context.Context, ID uuid.UUID) error
	Search(ctx context.Context, searchParams *models.SearchToRent) ([]*models.Transport, error)
}
