package transport

import "github.com/gin-gonic/gin"

type Handlers interface {
	Post() gin.HandlerFunc
}
