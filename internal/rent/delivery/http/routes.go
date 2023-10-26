package http

import (
	"github.com/gin-gonic/gin"
	"github.com/nikolaev-roman/simbir-go/config"
	"github.com/nikolaev-roman/simbir-go/internal/account"
	"github.com/nikolaev-roman/simbir-go/internal/middleware"
	"github.com/nikolaev-roman/simbir-go/internal/rent"
)

func MapRentRoutes(
	rentGroup *gin.RouterGroup,
	h rent.Handlers, us account.UseCase,
	cfg *config.Config,
	mv *middleware.MiddlewareManager,
) {
	rentGroup.GET("/Transport", h.SearchTransport())

	rentGroup.Use(mv.CheckAuth(us, cfg))
	rentGroup.GET("/:rent_id", h.Get())
	rentGroup.POST("/New/:transport_id", h.New())
	rentGroup.POST("/End/:rent_id", h.End())
}
