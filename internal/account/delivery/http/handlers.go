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

func (h *accountHandlers) SignUp() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		account := &models.Account{}

		if ctx.Bind(&account) != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": "Failed to read body",
			})

			return
		}

		createdAccount, err := h.accountUC.Register(nil, account)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
		}

		ctx.JSON(http.StatusOK, createdAccount)
	}
}

func (h *accountHandlers) SignIn() gin.HandlerFunc {
	type Login struct {
		Username string `json:"username"`
		Password string `json:"password,omitempty"`
	}
	return func(ctx *gin.Context) {
		login := &Login{}

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
		}

		ctx.JSON(http.StatusOK, token)
	}
}

func (h *accountHandlers) GetMe() gin.HandlerFunc {
	return func(c *gin.Context) {

		ctx := utils.GetRequestCtx(c)

		account, err := utils.GetAccountFromCtx(ctx)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
		}

		account, err = h.accountUC.GetByID(nil, account.ID)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
		}

		c.JSON(http.StatusOK, account)
	}
}

func (h *accountHandlers) Update() gin.HandlerFunc {
	return nil
}
