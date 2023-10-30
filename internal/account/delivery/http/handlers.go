package http

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/nikolaev-roman/simbir-go/config"
	"github.com/nikolaev-roman/simbir-go/internal/account"
	"github.com/nikolaev-roman/simbir-go/internal/models"
	"github.com/nikolaev-roman/simbir-go/pkg/utils"
)

type accountHandlers struct {
	cfg       *config.Config
	accountUC account.UseCase
}

func NewAccountHandlers(cfg *config.Config, accountUC account.UseCase) account.Handlers {
	return &accountHandlers{cfg: cfg, accountUC: accountUC}
}

// SignUp
// @Summary		Create new account
// @Schemes
// @Description	create new account
// @Tags		Account
// @Accept		json
// @Produce		json
// @Param   	request body models.AccountSign true "query params"
// @Success		200	{object} models.Account
// @Failure		500
// @Router		/Account/SignUp [post]
func (h *accountHandlers) SignUp() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		account := &models.Account{}
		accountSign := &models.AccountSign{}

		if ctx.Bind(&accountSign) != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": "Failed to read body",
			})
			return
		}

		account.Username = accountSign.Username
		account.Password = accountSign.Password

		createdAccount, err := h.accountUC.Register(nil, account)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		ctx.JSON(http.StatusOK, createdAccount)
	}
}

// SignIn
// @Summary		SignIn
// @Schemes
// @Description	jwt token getting
// @Tags		Account
// @Accept		json
// @Produce		json
// @Param   	request body models.AccountSign true "query params"
// @Success		200	{object} string
// @Failure		500
// @Router		/Account/SignIn [post]
func (h *accountHandlers) SignIn() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		login := &models.AccountSign{}

		if ctx.Bind(&login) != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": "Failed to read body",
			})
			return
		}

		token, err := h.accountUC.SignIn(nil, &models.Account{
			Username: login.Username,
			Password: login.Password,
		})
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		ctx.JSON(http.StatusOK, token)
	}
}

// Me
// @Summary		Me
// @Schemes
// @Description	getting account info
// @Tags		Account
// @Accept		json
// @Produce		json
// @Security 	Authorization
// @Success		200	{object} models.Account
// @Failure		500
// @Router		/Account/Me [get]
func (h *accountHandlers) GetMe() gin.HandlerFunc {
	return func(c *gin.Context) {

		ctx := utils.GetRequestCtx(c)

		account, err := utils.GetAccountFromCtx(ctx)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		account, err = h.accountUC.GetByID(nil, account.ID)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		c.JSON(http.StatusOK, account)
	}
}

// Update
// @Summary
// @Schemes
// @Description	update account info
// @Tags		Account
// @Accept		json
// @Produce		json
// @Security 	Authorization
// @Param   	request body models.AccountSign true "query params"
// @Success		200	{object} models.Account
// @Failure		500
// @Router		/Account/Update [put]
func (h *accountHandlers) Update() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := utils.GetRequestCtx(c)

		currentAccount, err := utils.GetAccountFromCtx(ctx)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		accountSign := &models.AccountSign{}
		if c.Bind(&accountSign) != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "Failed to read body",
			})
			return
		}

		currentAccount.Username = accountSign.Username
		currentAccount.Password = accountSign.Password

		updatedAccount, err := h.accountUC.Update(ctx, currentAccount)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, updatedAccount)
	}
}
