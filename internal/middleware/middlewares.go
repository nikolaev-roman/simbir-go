package middleware

import (
	"github.com/nikolaev-roman/simbir-go/config"
	"github.com/nikolaev-roman/simbir-go/internal/account"
)

type MiddlewareManager struct {
	accountUC account.UseCase
	cfg       *config.Config
	origins   []string
}

// Middleware manager constructor
func NewMiddlewareManager(accountUC account.UseCase, cfg *config.Config, origins []string) *MiddlewareManager {
	return &MiddlewareManager{accountUC: accountUC, cfg: cfg, origins: origins}
}
