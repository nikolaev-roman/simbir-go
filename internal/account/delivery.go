package account

import "github.com/gin-gonic/gin"

type Handlers interface {
	SignUp() gin.HandlerFunc
	SignIn() gin.HandlerFunc
	GetMe() gin.HandlerFunc
	Update() gin.HandlerFunc
}
