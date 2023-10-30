package http

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/nikolaev-roman/simbir-go/config"
	"github.com/nikolaev-roman/simbir-go/internal/account"
	"github.com/nikolaev-roman/simbir-go/internal/admin"
	"github.com/nikolaev-roman/simbir-go/internal/models"
	"github.com/nikolaev-roman/simbir-go/internal/rent"
	"github.com/nikolaev-roman/simbir-go/internal/transport"
	"github.com/nikolaev-roman/simbir-go/pkg/utils"
)

type adminHandlers struct {
	cfg         *config.Config
	accountUC   account.UseCase
	transportUC transport.UseCase
	rentUC      rent.UseCase
}

func NewAdminHandlers(cfg *config.Config,
	accountUC account.UseCase,
	transportUC transport.UseCase,
	rentUC rent.UseCase,
) admin.Handlers {
	return &adminHandlers{
		cfg:         cfg,
		accountUC:   accountUC,
		transportUC: transportUC,
		rentUC:      rentUC,
	}
}

// AccountList
// @Summary     AccountList
// @Schemes
// @Description	AccountList
// @Tags		Admin
// @Accept		json
// @Produce		json
// @Security 	Authorization
// @Param   	start query string true "start"
// @Param   	count query string true "count"
// @Success		200	{array} models.Account
// @Failure		500
// @Router		/Admin/Account [get]
func (h *adminHandlers) GetAccountList() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := utils.GetRequestCtx(c)

		count, err := strconv.Atoi(c.Query("count"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalide lat format: " + err.Error()})
		}

		start, err := strconv.Atoi(c.Query("start"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalide long format: " + err.Error()})
		}

		searchParams := &models.AccountSearchParams{
			Count: count,
			Start: start,
		}

		accountList, err := h.accountUC.Search(ctx, *searchParams)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, accountList)

	}
}

// Account
// @Summary     Account
// @Schemes
// @Description	Account
// @Tags		Admin
// @Accept		json
// @Produce		json
// @Security 	Authorization
// @Param   	account_id path string true "account ID"
// @Success		200	{object} models.Account
// @Failure		500
// @Router		/Admin/Account/{account_id} [get]
func (h *adminHandlers) GetAccount() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := utils.GetRequestCtx(c)

		accountID, err := uuid.Parse(c.Param("account_id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		account, err := h.accountUC.GetByID(ctx, accountID)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, account)
	}
}

// Create Account
// @Summary     Create Account
// @Schemes
// @Description	Create Account
// @Tags		Admin
// @Accept		json
// @Produce		json
// @Security 	Authorization
// @Param   	request body models.Account true "query params"
// @Success		200	{object} models.Account
// @Failure		500
// @Router		/Admin/Account [post]
func (h *adminHandlers) CreateAccount() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := utils.GetRequestCtx(c)

		account := &models.Account{}

		if c.Bind(&account) != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "Failed to read body",
			})
			return
		}

		createdAccount, err := h.accountUC.Create(ctx, account)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, createdAccount)

	}
}

// Update Account
// @Summary     Update Account
// @Schemes
// @Description	Update Account
// @Tags		Admin
// @Accept		json
// @Produce		json
// @Security 	Authorization
// @Param   	request body models.Account true "query params"
// @Success		200	{object} models.Account
// @Failure		500
// @Router		/Admin/Account [put]
func (h *adminHandlers) UpdateAccount() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := utils.GetRequestCtx(c)

		account := &models.Account{}

		if c.Bind(&account) != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "Failed to read body",
			})
			return
		}

		updatedAccount, err := h.accountUC.Update(ctx, account)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, updatedAccount)
	}
}

// Delete Account
// @Summary     Delete Account
// @Schemes
// @Description	Delete Account
// @Tags		Admin
// @Accept		json
// @Produce		json
// @Security 	Authorization
// @Param   	account_id path string true "account ID"
// @Success		200	{object} models.Account
// @Failure		500
// @Router		/Admin/Account/{account_id} [delete]
func (h *adminHandlers) DeleteAccount() gin.HandlerFunc {
	return func(c *gin.Context) {

		ctx := utils.GetRequestCtx(c)

		accountID, err := uuid.Parse(c.Param("account_id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		err = h.accountUC.Delete(ctx, accountID)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"success": true})
	}
}

// Transport List
// @Summary     Transport List
// @Schemes
// @Description	Transport List
// @Tags		Admin
// @Accept		json
// @Produce		json
// @Security 	Authorization
// @Param   	start query string true "start"
// @Param   	count query string true "count"
// @Success		200	{array} models.Transport
// @Failure		500
// @Router		/Admin/Transport [get]
func (h *adminHandlers) GetTransportList() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := utils.GetRequestCtx(c)

		count, err := strconv.Atoi(c.Query("count"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalide lat format: " + err.Error()})
		}

		start, err := strconv.Atoi(c.Query("start"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalide long format: " + err.Error()})
		}

		searchParams := &models.TransportSearchParams{
			Count: count,
			Start: start,
		}

		transportList, err := h.transportUC.Search(ctx, *searchParams)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, transportList)
	}
}

// Transport
// @Summary     Transport
// @Schemes
// @Description	Transport
// @Tags		Admin
// @Accept		json
// @Produce		json
// @Security 	Authorization
// @Param  		transport_id path string true "Trasnport ID"
// @Success		200	{object} models.Transport
// @Failure		500
// @Router		/Admin/Transport/{transport_id} [get]
func (h *adminHandlers) GetTransport() gin.HandlerFunc {
	return func(c *gin.Context) {

		ctx := utils.GetRequestCtx(c)

		transportID, err := uuid.Parse(c.Param("transport_id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		transport, err := h.transportUC.GetByID(ctx, transportID)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, transport)

	}
}

// Create Transport
// @Summary 	Create Transport
// @Schemes
// @Description	Create Transport
// @Tags		Admin
// @Accept		json
// @Produce		json
// @Security 	Authorization
// @Param   	request body models.Transport true "query params"
// @Success		200	{object} models.Transport
// @Failure		500
// @Router		/Admin/Transport/ [post]
func (h *adminHandlers) CreateTransport() gin.HandlerFunc {
	return func(c *gin.Context) {

		ctx := utils.GetRequestCtx(c)

		transport := &models.Transport{}

		if c.Bind(&transport) != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "Failed to read body",
			})
			return
		}

		createdTransport, err := h.transportUC.Create(ctx, transport)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, createdTransport)

	}
}

// Update
// @Summary 	update transport
// @Schemes
// @Description	update transport
// @Tags		Admin
// @Accept		json
// @Produce		json
// @Security 	Authorization
// @Param  		transport_id path string true "Trasnport ID"
// @Param   	request body models.Transport true "query params"
// @Success		200	{object} models.Transport
// @Failure		500
// @Router		/Admin/Transport/{transport_id} [put]
func (h *adminHandlers) UpdateTransport() gin.HandlerFunc {
	return func(c *gin.Context) {

		ctx := utils.GetRequestCtx(c)

		transportID, err := uuid.Parse(c.Param("transport_id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		transport := &models.Transport{}

		if c.Bind(&transport) != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "Failed to read body",
			})
			return
		}

		transport.ID = transportID

		updatedTransport, err := h.transportUC.Update(ctx, transport)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, updatedTransport)

	}
}

// Delete
// @Summary 	delete transport
// @Schemes
// @Description	delete transport
// @Tags		Admin
// @Accept		json
// @Produce		json
// @Security 	Authorization
// @Param  		transport_id path string true "Trasnport ID"
// @Success		200
// @Failure		500
// @Router		/Admin/Transport/{transport_id} [delete]
func (h *adminHandlers) DeleteTransport() gin.HandlerFunc {
	return func(c *gin.Context) {

		ctx := utils.GetRequestCtx(c)

		transportID, err := uuid.Parse(c.Param("transport_id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		err = h.transportUC.Delete(ctx, transportID)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"success": true})

	}
}

// Get Rent
// @Summary 	Get Rent
// @Schemes
// @Description	Get Rent
// @Tags		Admin
// @Accept		json
// @Produce		json
// @Param  		rent_id path string true "Rent ID"
// @Success		200	{object} models.Transport
// @Failure		500
// @Router		/Admin/Rent/{rent_id} [get]
func (h *adminHandlers) GetRent() gin.HandlerFunc {
	return func(c *gin.Context) {

		ctx := utils.GetRequestCtx(c)

		rentID, err := uuid.Parse(c.Param("rent_id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		rent, err := h.rentUC.GetByID(ctx, rentID)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, rent)

	}
}

// User History
// @Summary		Get user history
// @Schemes
// @Description	Get user history
// @Tags		Admin
// @Accept		json
// @Produce		json
// @Security	Authorization
// @Param  		account_id path string true "Account ID"
// @Success		200	{array} models.Rent
// @Failure		500
// @Router		/Admin/Rent/UserHistory/{account_id} [get]
func (h *adminHandlers) GetRentUserHistory() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := utils.GetRequestCtx(c)

		accountID, err := uuid.Parse(c.Param("account_id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		rentHistory, err := h.rentUC.HistoryByAccount(ctx, accountID)
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
// @Tags		Admin
// @Accept		json
// @Produce		json
// @Security	Authorization
// @Param  		transport_id path string true "Trasnport ID"
// @Success		200	{array} models.Rent
// @Failure		500
// @Router		/Admin/Rent/TransportHistory/{transport_id} [get]
func (h *adminHandlers) GetRentTransportHistory() gin.HandlerFunc {
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

// Create Rent
// @Summary 	Create Rent
// @Schemes
// @Description	Create Rent
// @Tags		Admin
// @Accept		json
// @Produce		json
// @Security 	Authorization
// @Param   	request body models.Rent true "query params"
// @Success		200	{object} models.Rent
// @Failure		500
// @Router		/Admin/Rent/ [post]
func (h *adminHandlers) CreateRent() gin.HandlerFunc {
	return func(c *gin.Context) {

		ctx := utils.GetRequestCtx(c)

		rent := &models.Rent{}

		if c.Bind(&rent) != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "Failed to read body",
			})
			return
		}

		createdRent, err := h.rentUC.Create(ctx, rent)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, createdRent)

	}
}

// End rent
// @Summary		End rent
// @Schemes
// @Description	End rent
// @Tags		Admin
// @Accept		json
// @Produce		json
// @Security	Authorization
// @Param   	rent_id path string true "Rent ID"
// @Success		200	{object} models.Rent
// @Failure		500
// @Router		/Admin/Rent/End/{rent_id} [post]
func (h *adminHandlers) EndRent() gin.HandlerFunc {
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

// Update Rent
// @Summary  Update Rent
// @Schemes
// @Description Update Rent
// @Tags		Admin
// @Accept		json
// @Produce		json
// @Security 	Authorization
// @Param   	rent_id path string true "Rent ID"
// @Param   	request body models.Rent true "query params"
// @Success		200	{object} models.Rent
// @Failure		500
// @Router		/Admin/Rent/{rent_id} [put]
func (h *adminHandlers) UpdateRent() gin.HandlerFunc {
	return func(c *gin.Context) {

		ctx := utils.GetRequestCtx(c)

		rent := &models.Rent{}

		rentID, err := uuid.Parse(c.Param("rent_id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		if c.Bind(&rent) != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "Failed to read body",
			})
			return
		}

		rent.ID = rentID

		updatedRent, err := h.rentUC.Update(ctx, rent)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, updatedRent)

	}
}

// Delete Rent
// @Summary  Delete Rent
// @Schemes
// @Description Delete Rent
// @Tags		Admin
// @Accept		json
// @Produce		json
// @Security 	Authorization
// @Param   	rent_id path string true "Rent ID"
// @Success		200
// @Failure		500
// @Router		/Admin/Rent/{rent_id} [delete]
func (h *adminHandlers) DeleteRent() gin.HandlerFunc {
	return func(c *gin.Context) {

		ctx := utils.GetRequestCtx(c)

		rentID, err := uuid.Parse(c.Param("rent_id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		err = h.rentUC.Delete(ctx, rentID)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"success": true})

	}
}
