package http

import (
	"github.com/gin-gonic/gin"
	"github.com/nikolaev-roman/simbir-go/config"
	"github.com/nikolaev-roman/simbir-go/internal/account"
	"github.com/nikolaev-roman/simbir-go/internal/middleware"
)

func MapAccountRoutes(accountGroup *gin.RouterGroup, h account.Handlers, uc account.UseCase, cfg *config.Config, mv *middleware.MiddlewareManager) {
	accountGroup.POST("/SignUp", h.SignUp())
	accountGroup.POST("/SignIn", h.SignIn())

	accountGroup.Use(mv.CheckAuth(uc, cfg))
	accountGroup.GET("/Me", h.GetMe())
	accountGroup.PUT("/Update", h.Update())
}
