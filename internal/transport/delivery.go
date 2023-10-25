package transport

import "github.com/gin-gonic/gin"

type Handlers interface {
	Post() gin.HandlerFunc
	Get() gin.HandlerFunc
	Put() gin.HandlerFunc
	Delete() gin.HandlerFunc
}
