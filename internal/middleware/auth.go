package middleware

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
	"github.com/nikolaev-roman/simbir-go/config"
	"github.com/nikolaev-roman/simbir-go/internal/account"
	"github.com/nikolaev-roman/simbir-go/pkg/utils"
)

func extractBearerToken(header string) (string, error) {
	if header == "" {
		return "", errors.New("bad header value given")
	}

	jwtToken := strings.Split(header, " ")
	if len(jwtToken) != 2 {
		return "", errors.New("incorrectly formatted authorization header")
	}

	return jwtToken[1], nil
}

func (mw *MiddlewareManager) CheckAuth(accountUC account.UseCase, cfg *config.Config) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		tokenString, err := extractBearerToken(ctx.GetHeader("Authorization"))

		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
			}

			return []byte(cfg.Server.JwtSecretKey), nil
		})

		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {

			if float64(time.Now().Unix()) > claims["exp"].(float64) {
				ctx.AbortWithStatus(http.StatusUnauthorized)
			}

			userID, ok := claims["sub"].(string)
			if !ok {
				ctx.AbortWithStatus(http.StatusUnauthorized)
			}

			userUUID, err := uuid.Parse(userID)
			if err != nil {
				ctx.AbortWithStatus(http.StatusUnauthorized)
			}

			ac, err := accountUC.GetByID(ctx.Request.Context(), userUUID)
			if err != nil {
				ctx.AbortWithStatus(http.StatusUnauthorized)
			}

			ctx.Set("account", ac)

			context := context.WithValue(ctx.Request.Context(), utils.AccountCtxKey{}, ac)

			ctx.Request = ctx.Request.WithContext(context)

			ctx.Next()

		} else {
			ctx.AbortWithStatus(http.StatusUnauthorized)
		}
	}
}
