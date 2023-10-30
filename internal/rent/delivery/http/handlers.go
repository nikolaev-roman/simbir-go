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

// Create new rent
// @Summary		New rent
// @Schemes
// @Description	New rent
// @Tags		Rent
// @Accept		json
// @Produce		json
// @Security	Authorization
// @Param   	transport_id path string true "Transport ID"
// @Param   	rentType query string true "rent type"
// @Success		200	{object} models.Rent
// @Failure		500
// @Router		/Rent/New/{transport_id} [post]
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

// Get rent
// @Summary		Get rent info
// @Schemes
// @Description	Get rent info
// @Tags		Rent
// @Accept		json
// @Produce		json
// @Security	Authorization
// @Param   	rent_id path string true "Rent ID"
// @Success		200	{object} models.Transport
// @Failure		500
// @Router		/Rent/{rent_id} [get]
func (h *rentHandlers) Get() gin.HandlerFunc {
	return func(c *gin.Context) {

		ctx := utils.GetRequestCtx(c)

		rentID, err := uuid.Parse(c.Param("rent_id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		rent, err := h.rentUC.GetByIDForUser(ctx, rentID)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, rent)

	}
}

// End rent
// @Summary		End rent
// @Schemes
// @Description	End rent
// @Tags		Rent
// @Accept		json
// @Produce		json
// @Security	Authorization
// @Param   	rent_id path string true "Rent ID"
// @Success		200	{object} models.Rent
// @Failure		500
// @Router		/Rent/End/{rent_id} [post]
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

// SearchTransport
// @Summary		Get Transports to rent
// @Schemes
// @Description	Get Transports to rent
// @Tags		Rent
// @Accept		json
// @Produce		json
// @Param   	lat query string true "latitude"
// @Param   	long query string true "longitude"
// @Param   	radius query string true "radius"
// @Param   	type query string true "transport type"
// @Success		200	{array} models.Transport
// @Failure		500
// @Router		/Rent/Transport [get]
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

// My History
// @Summary		Get user history
// @Schemes
// @Description	Get user history
// @Tags		Rent
// @Accept		json
// @Produce		json
// @Security	Authorization
// @Success		200	{array} models.Rent
// @Failure		500
// @Router		/Rent/MyHistory [get]
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

// Transport History
// @Summary		Transport History
// @Schemes
// @Description	Transport History
// @Tags		Rent
// @Accept		json
// @Produce		json
// @Security	Authorization
// @Param  		transport_id path string true "Trasnport ID"
// @Success		200	{array} models.Rent
// @Failure		500
// @Router		/Rent/TransportHistory/{transport_id} [get]
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
