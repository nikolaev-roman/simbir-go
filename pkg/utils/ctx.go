package utils

import (
	"context"
	"errors"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/nikolaev-roman/simbir-go/internal/models"
)

type AccountCtxKey struct{}

func GetAccountFromCtx(ctx context.Context) (*models.Account, error) {
	fmt.Println(ctx)

	account, ok := ctx.Value(AccountCtxKey{}).(*models.Account)
	if !ok {
		return nil, errors.New("no ctx account")
	}

	return account, nil
}

func GetRequestCtx(c *gin.Context) context.Context {
	return c.Request.Context()
}
