package http

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/nikolaev-roman/simbir-go/config"
	"github.com/nikolaev-roman/simbir-go/internal/models"
	"github.com/nikolaev-roman/simbir-go/internal/rent"
	"github.com/nikolaev-roman/simbir-go/pkg/utils"
)

type rentHandlers struct {
	cfg    *config.Config
	rentUC rent.UseCase
}

func NewRentHandlers(cfg *config.Config, rentUC rent.UseCase) rent.Handlers {
	return &rentHandlers{cfg: cfg, rentUC: rentUC}
}

func (h *rentHandlers) New() gin.HandlerFunc {
	return func(c *gin.Context) {

		ctx := utils.GetRequestCtx(c)

		transportID, err := uuid.Parse(c.Param("transport_id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		rentType := c.Query("rentType")

		createdRent, err := h.rentUC.Start(ctx, transportID, rentType)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
		}

		c.JSON(http.StatusOK, createdRent)
	}
}

func (h *rentHandlers) Get() gin.HandlerFunc {
	return func(c *gin.Context) {

		ctx := utils.GetRequestCtx(c)

		rentID, err := uuid.Parse(c.Param("rent_id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		rent, err := h.rentUC.GetByID(ctx, rentID)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, rent)

	}
}

func (h *rentHandlers) End() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := utils.GetRequestCtx(c)

		rentID, err := uuid.Parse(c.Param("rent_id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		latitude, err := strconv.ParseFloat(c.Query("lat"), 64)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalide lat format: " + err.Error()})
		}

		longitude, err := strconv.ParseFloat(c.Query("long"), 64)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalide long format: " + err.Error()})
		}

		coordinates := &models.Coordinates{
			Latitude:  latitude,
			Longitude: longitude,
		}

		rent, err := h.rentUC.End(ctx, rentID, coordinates)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, rent)
	}
}

func (h *rentHandlers) SearchTransport() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := utils.GetRequestCtx(c)

		latitude, err := strconv.ParseFloat(c.Query("lat"), 64)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalide lat format: " + err.Error()})
			return
		}

		longitude, err := strconv.ParseFloat(c.Query("long"), 64)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalide long format: " + err.Error()})
			return
		}

		radius, err := strconv.ParseFloat(c.Query("radius"), 64)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalide radius format: " + err.Error()})
			return
		}

		Ttype := c.Query("type")

		searchParams := &models.SearchToRent{
			Lat:    latitude,
			Long:   longitude,
			Radius: radius,
			Type:   Ttype,
		}

		transports, err := h.rentUC.SearchTransportToRent(ctx, searchParams)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		fmt.Println(transports)

		c.JSON(http.StatusOK, transports)
	}
}

func (h *rentHandlers) MyHistory() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := utils.GetRequestCtx(c)

		account, err := utils.GetAccountFromCtx(ctx)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			return
		}

		rentHistory, err := h.rentUC.HistoryByAccount(ctx, account.ID)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, rentHistory)
	}
}

func (h *rentHandlers) TransportHistory() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := utils.GetRequestCtx(c)

		transportID, err := uuid.Parse(c.Param("transport_id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		rentHistory, err := h.rentUC.HistoryByTransport(ctx, transportID)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, rentHistory)
	}
}
