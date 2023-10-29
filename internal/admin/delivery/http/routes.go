package http

import (
	"github.com/gin-gonic/gin"
	"github.com/nikolaev-roman/simbir-go/config"
	"github.com/nikolaev-roman/simbir-go/internal/account"
	"github.com/nikolaev-roman/simbir-go/internal/admin"
	"github.com/nikolaev-roman/simbir-go/internal/middleware"
)

func MapAdminRoutes(
	adminGroup *gin.RouterGroup,
	h admin.Handlers,
	us account.UseCase,
	cfg *config.Config,
	mw *middleware.MiddlewareManager,
) {
	adminGroup.Use(mw.CheckAuth(us, cfg))
	adminGroup.Use(mw.CheckAdmin(us, cfg))

	accounts := adminGroup.Group("/Account")

	accounts.GET("/", h.GetAccountList())
	accounts.GET("/:account_id", h.GetAccount())
	accounts.POST("/", h.CreateAccount())
	accounts.PUT("/:account_id", h.UpdateAccount())
	accounts.DELETE("/:account_id", h.DeleteAccount())

	transports := adminGroup.Group("/Transport")

	transports.GET("/", h.GetTransportList())
	transports.GET("/:transport_id", h.GetTransport())
	transports.POST("/", h.CreateTransport())
	transports.PUT("/:transport_id", h.UpdateTransport())
	transports.DELETE("/:transport_id", h.DeleteTransport())

	rents := adminGroup.Group("/Rent")

	rents.GET("/:rent_id", h.GetRent())
	rents.GET("/UserHistory/:account_id", h.GetRentUserHistory())
	rents.GET("/TransportHistory/:transport_id", h.GetRentTransportHistory())
	rents.POST("/", h.CreateRent())
	rents.POST("/End/:rent_id", h.EndRent())
	rents.PUT("/:rent_id", h.UpdateRent())
	rents.DELETE("/:rent_id", h.DeleteRent())
}
