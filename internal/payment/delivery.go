package payment

import "github.com/gin-gonic/gin"

type Handlers interface {
	PaymentHesoyam() gin.HandlerFunc
}
