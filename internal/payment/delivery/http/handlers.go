package http

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
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

// Payment
// @Summary		Enrich balance
// @Schemes
// @Description	Enrich balance
// @Tags		Payment
// @Accept		json
// @Produce		json
// @Security	Authorization
// @Param   	account_id path string true "account ID"
// @Success		200	{object} models.Account
// @Failure		500
// @Router		/Payment/Hesoyam/{account_id} [post]
func (h *paymentHandlers) PaymentHesoyam() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := utils.GetRequestCtx(c)

		accountID, err := uuid.Parse(c.Param("account_id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		account, err := utils.GetAccountFromCtx(ctx)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			return
		}

		if account.IsAdmin != true && accountID != account.ID {
			c.JSON(http.StatusForbidden, gin.H{"error": "forbidden"})
			return
		}

		updatedAccount, err := h.paymentUC.EnrichBalance(ctx, accountID)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, updatedAccount)
	}
}
