package usecase

import (
	"context"
	"errors"

	"github.com/google/uuid"
	"github.com/nikolaev-roman/simbir-go/config"
	"github.com/nikolaev-roman/simbir-go/internal/account"
	"github.com/nikolaev-roman/simbir-go/internal/models"
	"github.com/nikolaev-roman/simbir-go/internal/transport"
	"github.com/nikolaev-roman/simbir-go/pkg/utils"
)

type transportUC struct {
	cfg           *config.Config
	transportRepo transport.Repository
}

func NewTransportUseCase(cfg *config.Config, transportRepo transport.Repository, accountRepo account.Repository) transport.UseCase {
	return &transportUC{cfg: cfg, transportRepo: transportRepo}
}

func (u *transportUC) Create(ctx context.Context, transport *models.Transport) (*models.Transport, error) {

	account, err := utils.GetAccountFromCtx(ctx)
	if err != nil {
		return nil, err
	}

	transport.OwnerID = account.ID

	if err = utils.ValidateStruct(ctx, transport); err != nil {
		return nil, err
	}

	return u.transportRepo.Create(ctx, transport)
}

func (u *transportUC) GetByID(ctx context.Context, ID uuid.UUID) (*models.Transport, error) {
	return u.transportRepo.GetByID(ctx, ID)
}

func (u *transportUC) Update(ctx context.Context, transport *models.Transport) (*models.Transport, error) {

	transportByID, err := u.transportRepo.GetByID(ctx, transport.ID)
	if err != nil {
		return nil, err
	}

	account, err := utils.GetAccountFromCtx(ctx)
	if err != nil {
		return nil, err
	}

	if transportByID.OwnerID != account.ID {
		return nil, errors.New("forbidden")
	}

	transport.OwnerID = account.ID

	return u.transportRepo.Update(ctx, transport)
}

func (u *transportUC) Delete(ctx context.Context, ID uuid.UUID) error {
	transportByID, err := u.transportRepo.GetByID(ctx, ID)
	if err != nil {
		return err
	}

	account, err := utils.GetAccountFromCtx(ctx)
	if err != nil {
		return err
	}

	if transportByID.OwnerID != account.ID {
		return errors.New("forbidden")
	}

	return u.transportRepo.Delete(ctx, ID)
}

func (u *transportUC) RentingStart(ctx context.Context, transportID uuid.UUID) (*models.Transport, error) {

	account, err := utils.GetAccountFromCtx(ctx)
	if err != nil {
		return nil, err
	}

	transportByID, err := u.GetByID(ctx, transportID)
	if err != nil {
		return nil, err
	}

	if transportByID.OwnerID == account.ID {
		return nil, errors.New("You cannot rent your own transport")
	}

	if transportByID.CanBeRented != true {
		return nil, errors.New("Transport already rented")
	}

	transportByID.CanBeRented = false

	return u.transportRepo.Update(ctx, transportByID)
}

func (u *transportUC) RentingClose(ctx context.Context, transport *models.Transport) (*models.Transport, error) {

	transport.CanBeRented = true

	return u.transportRepo.Update(ctx, transport)
}

func (u *transportUC) RentingEnd(ctx context.Context, ID uuid.UUID, coordinates *models.Coordinates) (*models.Transport, error) {
	transport, err := u.GetByID(ctx, ID)
	if err != nil {
		return nil, err
	}

	transport.Latitude = coordinates.Latitude
	transport.Longitude = coordinates.Longitude
	transport.CanBeRented = true
	updated, err := u.transportRepo.Update(ctx, transport)
	if err != nil {
		return nil, err
	}

	return updated, nil
}

func (u *transportUC) SearchToRent(ctx context.Context, searchParams *models.SearchToRent) ([]*models.Transport, error) {
	return u.transportRepo.SearchToRent(ctx, searchParams)
}

func (u *transportUC) Search(ctx context.Context, searchParams models.TransportSearchParams) ([]*models.Transport, error) {
	return u.transportRepo.Search(ctx, searchParams)
}
