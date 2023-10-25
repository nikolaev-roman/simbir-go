package http

import (
	"github.com/gin-gonic/gin"
	"github.com/nikolaev-roman/simbir-go/config"
	"github.com/nikolaev-roman/simbir-go/internal/account"
	"github.com/nikolaev-roman/simbir-go/internal/middleware"
	"github.com/nikolaev-roman/simbir-go/internal/transport"
)

func MapTransportRoutes(transportGroup *gin.RouterGroup, h transport.Handlers, us account.UseCase, cfg *config.Config, mv *middleware.MiddlewareManager) {
	transportGroup.GET("/:transport_id", h.Get())

	transportGroup.Use(mv.CheckAuth(us, cfg))
	transportGroup.POST("", h.Post())
	transportGroup.PUT("/:transport_id", h.Put())
	transportGroup.DELETE("/:transport_id", h.Delete())
}
