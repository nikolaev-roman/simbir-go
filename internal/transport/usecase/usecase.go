package usecase

import (
	"context"

	"github.com/nikolaev-roman/simbir-go/config"
	"github.com/nikolaev-roman/simbir-go/internal/account"
	"github.com/nikolaev-roman/simbir-go/internal/models"
	"github.com/nikolaev-roman/simbir-go/internal/transport"
	"github.com/nikolaev-roman/simbir-go/pkg/utils"
)

type transportUC struct {
	cfg           *config.Config
	transportRepo transport.Repository
	accountRepo   account.Repository
}

func NewTransportUseCase(cfg *config.Config, transportRepo transport.Repository, accountRepo account.Repository) transport.UseCase {
	return &transportUC{cfg: cfg, transportRepo: transportRepo, accountRepo: accountRepo}
}

func (u *transportUC) Create(ctx context.Context, transport *models.Transport) (*models.Transport, error) {

	account, err := utils.GetAccountFromCtx(ctx)
	if err != nil {
		return nil, err
	}

	transport.OwnerID = account.ID

	return u.transportRepo.Create(ctx, transport)
}
