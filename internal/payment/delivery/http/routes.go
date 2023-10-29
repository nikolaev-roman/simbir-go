package http

import (
	"github.com/gin-gonic/gin"
	"github.com/nikolaev-roman/simbir-go/config"
	"github.com/nikolaev-roman/simbir-go/internal/account"
	"github.com/nikolaev-roman/simbir-go/internal/middleware"
	"github.com/nikolaev-roman/simbir-go/internal/payment"
)

func MapPaymentRoutes(paymentGroup *gin.RouterGroup, h payment.Handlers,
	uc account.UseCase, cfg *config.Config, mv *middleware.MiddlewareManager,
) {
	paymentGroup.Use(mv.CheckAuth(uc, cfg))
	paymentGroup.POST("/Hesoyam/:account_id", h.PaymentHesoyam())
}
