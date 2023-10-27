package usecase

import (
	"context"
	"errors"
	"time"

	"github.com/google/uuid"
	"github.com/nikolaev-roman/simbir-go/config"
	"github.com/nikolaev-roman/simbir-go/internal/models"
	"github.com/nikolaev-roman/simbir-go/internal/rent"
	"github.com/nikolaev-roman/simbir-go/internal/transport"
	"github.com/nikolaev-roman/simbir-go/pkg/utils"
)

type rentUC struct {
	cfg         *config.Config
	rentRepo    rent.Repository
	transportUC transport.UseCase
}

func NewRentUseCase(
	cfg *config.Config,
	rentRepo rent.Repository,
	transportUC transport.UseCase,
) rent.UseCase {
	return &rentUC{cfg: cfg, rentRepo: rentRepo, transportUC: transportUC}
}

func (u *rentUC) Start(ctx context.Context, transportID uuid.UUID, rentType string) (*models.Rent, error) {

	account, err := utils.GetAccountFromCtx(ctx)
	if err != nil {
		return nil, err
	}

	transport, err := u.transportUC.RentingStart(ctx, transportID)
	if err != nil {
		return nil, err
	}

	priceOfUnit, err := transport.GetPriceByType(rentType)

	rent := &models.Rent{
		TransportID: transport.ID,
		UserID:      account.ID,
		TimeStart:   time.Now(),
		PriceOfUnit: priceOfUnit,
		PriceType:   rentType,
	}

	startedRent, err := u.rentRepo.Create(ctx, rent)
	if err != nil {

		u.transportUC.RentingClose(ctx, transport)

		return nil, err
	}

	return startedRent, nil
}

func (u *rentUC) End(ctx context.Context, rentID uuid.UUID, coordinates *models.Coordinates) (*models.Rent, error) {

	account, err := utils.GetAccountFromCtx(ctx)
	if err != nil {
		return nil, err
	}

	rent, err := u.GetByID(ctx, rentID)
	if err != nil {
		return nil, err
	}

	if rent.UserID != account.ID {
		return nil, errors.New("forbidden")
	}

	rent.TimeEnd = time.Now()

	endedRent, err := u.rentRepo.Update(ctx, rent)
	if err != nil {
		return nil, err
	}

	_, err = u.transportUC.RentingEnd(ctx, endedRent.TransportID, coordinates)
	if err != nil {
		return nil, err
	}

	return endedRent, nil
}

func (u *rentUC) GetByID(ctx context.Context, rentID uuid.UUID) (*models.Rent, error) {

	account, err := utils.GetAccountFromCtx(ctx)
	if err != nil {
		return nil, err
	}

	rent, err := u.rentRepo.GetByID(ctx, rentID)
	if err != nil {
		return nil, err
	}

	transport, err := u.transportUC.GetByID(ctx, rent.TransportID)
	if err != nil {
		return nil, err
	}

	if account.ID != rent.UserID && account.ID != transport.OwnerID {
		return nil, errors.New("forbidden")
	}

	return rent, nil
}

func (u *rentUC) SearchTransportToRent(ctx context.Context, searchParams *models.SearchToRent) ([]*models.Transport, error) {
	return u.transportUC.Search(ctx, searchParams)
}

func (u *rentUC) HistoryByAccount(ctx context.Context, accountID uuid.UUID) ([]*models.Rent, error) {
	return u.rentRepo.GetHystoryByAccount(ctx, accountID)
}

func (u *rentUC) HistoryByTransport(ctx context.Context, transportID uuid.UUID) ([]*models.Rent, error) {
	return u.rentRepo.GetHystoryByTransport(ctx, transportID)
}
