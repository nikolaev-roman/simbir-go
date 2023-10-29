package http

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/nikolaev-roman/simbir-go/config"
	"github.com/nikolaev-roman/simbir-go/internal/payment"
	"github.com/nikolaev-roman/simbir-go/pkg/utils"
)

type paymentHandlers struct {
	cfg       *config.Config
	paymentUC payment.UseCase
}

func NewPaymentHandlers(cfg *config.Config, paymentUC payment.UseCase) payment.Handlers {
	return &paymentHandlers{cfg: cfg, paymentUC: paymentUC}
}

func (h *paymentHandlers) PaymentHesoyam() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := utils.GetRequestCtx(c)

		account, err := utils.GetAccountFromCtx(ctx)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			return
		}

		updatedAccount, err := h.paymentUC.EnrichBalance(ctx, account.ID)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, updatedAccount)
	}
}
