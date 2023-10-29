package usecase

import (
	"context"
	"errors"

	"github.com/google/uuid"
	"github.com/nikolaev-roman/simbir-go/config"
	"github.com/nikolaev-roman/simbir-go/internal/account"
	"github.com/nikolaev-roman/simbir-go/internal/models"
	"github.com/nikolaev-roman/simbir-go/internal/payment"
)

type paymentUC struct {
	cfg         *config.Config
	accountRepo account.Repository
}

func NewPaymentUseCase(cfg *config.Config, accountRepo account.Repository) payment.UseCase {
	return &paymentUC{cfg: cfg, accountRepo: accountRepo}
}

func (u *paymentUC) EnrichBalance(context context.Context, accountID uuid.UUID) (*models.Account, error) {

	foundAccount, err := u.accountRepo.GetByID(context, accountID)
	if err != nil {
		return nil, errors.New("account not found")
	}

	foundAccount.Balance = foundAccount.Balance + 250000

	updatedAccount, err := u.accountRepo.Update(context, foundAccount)
	if err != nil {
		return nil, errors.New("Failed to update account")
	}

	return updatedAccount, nil
}
