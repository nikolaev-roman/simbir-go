package http

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/nikolaev-roman/simbir-go/config"
	"github.com/nikolaev-roman/simbir-go/internal/models"
	"github.com/nikolaev-roman/simbir-go/internal/transport"
	"github.com/nikolaev-roman/simbir-go/pkg/utils"
)

type transportHandlers struct {
	cfg         *config.Config
	transportUC transport.UseCase
}

func NewAccountHandlers(cfg *config.Config, transportUC transport.UseCase) transport.Handlers {
	return &transportHandlers{cfg: cfg, transportUC: transportUC}
}

func (h *transportHandlers) Post() gin.HandlerFunc {
	return func(c *gin.Context) {
		transport := &models.Transport{}

		ctx := utils.GetRequestCtx(c)

		if c.Bind(&transport) != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "Failed to read body",
			})

			return
		}

		createdTransport, err := h.transportUC.Create(ctx, transport)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
		}

		c.JSON(http.StatusOK, createdTransport)
	}
}
