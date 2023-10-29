package transport

import (
	"context"

	"github.com/google/uuid"
	"github.com/nikolaev-roman/simbir-go/internal/models"
)

type UseCase interface {
	Create(ctx context.Context, transport *models.Transport) (*models.Transport, error)
	GetByID(ctx context.Context, ID uuid.UUID) (*models.Transport, error)
	Update(ctx context.Context, transport *models.Transport) (*models.Transport, error)
	Delete(ctx context.Context, ID uuid.UUID) error
	SearchToRent(ctx context.Context, searchParams *models.SearchToRent) ([]*models.Transport, error)
	Search(ctx context.Context, searchParams models.TransportSearchParams) ([]*models.Transport, error)

	RentingStart(ctx context.Context, ID uuid.UUID) (*models.Transport, error)
	RentingClose(ctx context.Context, transport *models.Transport) (*models.Transport, error)
	RentingEnd(ctx context.Context, ID uuid.UUID, coordinates *models.Coordinates) (*models.Transport, error)
}
