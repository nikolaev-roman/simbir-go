package rent

import "github.com/gin-gonic/gin"

type Handlers interface {
	New() gin.HandlerFunc
	Get() gin.HandlerFunc
	End() gin.HandlerFunc
	SearchTransport() gin.HandlerFunc
	MyHistory() gin.HandlerFunc
	TransportHistory() gin.HandlerFunc
}
