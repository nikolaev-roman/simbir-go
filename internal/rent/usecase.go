package rent

import (
	"context"

	"github.com/google/uuid"
	"github.com/nikolaev-roman/simbir-go/internal/models"
)

type UseCase interface {
	Create(ctx context.Context, rent *models.Rent) (*models.Rent, error)
	Update(ctx context.Context, rent *models.Rent) (*models.Rent, error)
	Delete(ctx context.Context, rentID uuid.UUID) error
	Start(ctx context.Context, transportID uuid.UUID, rentType string) (*models.Rent, error)
	End(ctx context.Context, rentID uuid.UUID, coordinates *models.Coordinates) (*models.Rent, error)
	GetByID(ctx context.Context, rentID uuid.UUID) (*models.Rent, error)
	GetByIDForUser(ctx context.Context, rentID uuid.UUID) (*models.Rent, error)
	SearchTransportToRent(ctx context.Context, searchParams *models.SearchToRent) ([]*models.Transport, error)
	HistoryByAccount(ctx context.Context, accountID uuid.UUID) ([]*models.Rent, error)
	HistoryByTransport(ctx context.Context, transportID uuid.UUID) ([]*models.Rent, error)
}
