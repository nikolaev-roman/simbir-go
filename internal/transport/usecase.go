package transport

import (
	"context"

	"github.com/nikolaev-roman/simbir-go/internal/models"
)

type UseCase interface {
	Create(ctx context.Context, transport *models.Transport) (*models.Transport, error)
}
