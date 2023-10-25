package repository

import (
	"context"
	"errors"

	"github.com/google/uuid"
	"github.com/nikolaev-roman/simbir-go/internal/models"
	"github.com/nikolaev-roman/simbir-go/internal/transport"
	"gorm.io/gorm"
)

type transportRepo struct {
	db *gorm.DB
}

func NewTransportRepository(db *gorm.DB) transport.Repository {
	return &transportRepo{db: db}
}

func (r *transportRepo) Create(ctx context.Context, transport *models.Transport) (*models.Transport, error) {

	result := r.db.Save(&transport)
	if result.Error != nil {
		return nil, result.Error
	}

	return transport, nil
}

func (r *transportRepo) GetByID(ctx context.Context, ID uuid.UUID) (*models.Transport, error) {
	var transport models.Transport
	result := r.db.First(&transport, "id = ?", ID)
	if result.Error != nil {
		return nil, errors.New("No transport found")
	}

	return &transport, nil
}

func (r *transportRepo) Update(ctx context.Context, transport *models.Transport) (*models.Transport, error) {
	result := r.db.Save(&transport)
	if result.Error != nil {
		return nil, result.Error
	}

	return transport, nil
}

func (r *transportRepo) Delete(ctx context.Context, ID uuid.UUID) error {
	result := r.db.Delete(&models.Transport{}, ID)
	return result.Error
}
