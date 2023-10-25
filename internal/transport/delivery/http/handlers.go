package http

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
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

func (h transportHandlers) Get() gin.HandlerFunc {
	return func(c *gin.Context) {

		ctx := utils.GetRequestCtx(c)

		ID, err := uuid.Parse(c.Param("transport_id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		transport, err := h.transportUC.GetByID(ctx, ID)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, transport)
	}
}

func (h *transportHandlers) Put() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := utils.GetRequestCtx(c)

		ID, err := uuid.Parse(c.Param("transport_id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		transport := &models.Transport{}
		if err = c.Bind(transport); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		transport.ID = ID

		updatedTransport, err := h.transportUC.Update(ctx, transport)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, updatedTransport)
	}
}

func (h *transportHandlers) Delete() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := utils.GetRequestCtx(c)

		ID, err := uuid.Parse(c.Param("transport_id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		err = h.transportUC.Delete(ctx, ID)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"status": "deleted"})
	}
}
