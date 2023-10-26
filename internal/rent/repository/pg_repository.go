package repository

import (
	"context"
	"errors"

	"github.com/google/uuid"
	"github.com/nikolaev-roman/simbir-go/internal/models"
	"github.com/nikolaev-roman/simbir-go/internal/rent"
	"gorm.io/gorm"
)

type rentRepo struct {
	db *gorm.DB
}

func NewRentRepository(db *gorm.DB) rent.Repository {
	return &rentRepo{db: db}
}

func (r *rentRepo) Create(ctx context.Context, rent *models.Rent) (*models.Rent, error) {
	result := r.db.Save(&rent)
	if result.Error != nil {
		return nil, result.Error
	}

	return rent, nil
}

func (r *rentRepo) Update(ctx context.Context, rent *models.Rent) (*models.Rent, error) {
	result := r.db.Save(&rent)
	if result.Error != nil {
		return nil, result.Error
	}

	return rent, nil
}

func (r *rentRepo) GetByID(ctx context.Context, ID uuid.UUID) (*models.Rent, error) {
	var rent models.Rent
	result := r.db.First(&rent, "id = ?", ID)
	if result.Error != nil {
		return nil, errors.New("No transport found")
	}

	return &rent, nil
}

func (r *rentRepo) Delete(ctx context.Context, ID uuid.UUID) error {
	result := r.db.Delete(&models.Transport{}, ID)
	return result.Error
}
