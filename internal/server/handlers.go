package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/nikolaev-roman/simbir-go/docs"

	accountHttp "github.com/nikolaev-roman/simbir-go/internal/account/delivery/http"
	accountRepository "github.com/nikolaev-roman/simbir-go/internal/account/repository"
	accountUseCase "github.com/nikolaev-roman/simbir-go/internal/account/usecase"

	transportHttp "github.com/nikolaev-roman/simbir-go/internal/transport/delivery/http"
	transportRepository "github.com/nikolaev-roman/simbir-go/internal/transport/repository"
	transportUseCase "github.com/nikolaev-roman/simbir-go/internal/transport/usecase"

	rentHttp "github.com/nikolaev-roman/simbir-go/internal/rent/delivery/http"
	rentRepository "github.com/nikolaev-roman/simbir-go/internal/rent/repository"
	rentUseCase "github.com/nikolaev-roman/simbir-go/internal/rent/usecase"

	"github.com/nikolaev-roman/simbir-go/internal/middleware"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func (s *Server) MapHandlers(server *gin.Engine) error {

	// Init Repositories
	accountRepo := accountRepository.NewAccountRepository(s.db)
	transportRepo := transportRepository.NewTransportRepository(s.db)
	rentRepo := rentRepository.NewRentRepository(s.db)

	// Init UseCases
	accountUC := accountUseCase.NewAccountUseCase(s.cfg, accountRepo)
	transportUC := transportUseCase.NewTransportUseCase(s.cfg, transportRepo, accountRepo)
	rentUC := rentUseCase.NewRentUseCase(s.cfg, rentRepo, transportUC)

	// Init handlers
	accountHandlers := accountHttp.NewAccountHandlers(s.cfg, accountUC)
	transportHandlers := transportHttp.NewAccountHandlers(s.cfg, transportUC)
	rentHandlers := rentHttp.NewRentHandlers(s.cfg, rentUC)

	mw := middleware.NewMiddlewareManager(accountUC, s.cfg, []string{"*"})

	docs.SwaggerInfo.BasePath = "/api"
	server.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	api := server.Group("/api")

	accountGroup := api.Group("/Account")
	accountHttp.MapAccountRoutes(accountGroup, accountHandlers, accountUC, s.cfg, mw)

	transportGroup := api.Group("/Transport")
	transportHttp.MapTransportRoutes(transportGroup, transportHandlers, accountUC, s.cfg, mw)

	rentGroup := api.Group("/Rent")
	rentHttp.MapRentRoutes(rentGroup, rentHandlers, accountUC, s.cfg, mw)

	api.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "ok",
		})
	})
	return nil
}
