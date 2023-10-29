package admin

import "github.com/gin-gonic/gin"

type Handlers interface {

	// Account handlers
	GetAccountList() gin.HandlerFunc
	GetAccount() gin.HandlerFunc
	CreateAccount() gin.HandlerFunc
	UpdateAccount() gin.HandlerFunc
	DeleteAccount() gin.HandlerFunc

	// Transport Handlers
	GetTransportList() gin.HandlerFunc
	GetTransport() gin.HandlerFunc
	CreateTransport() gin.HandlerFunc
	UpdateTransport() gin.HandlerFunc
	DeleteTransport() gin.HandlerFunc

	// Rent Handlers
	GetRent() gin.HandlerFunc
	GetRentUserHistory() gin.HandlerFunc
	GetRentTransportHistory() gin.HandlerFunc
	CreateRent() gin.HandlerFunc
	EndRent() gin.HandlerFunc
	UpdateRent() gin.HandlerFunc
	DeleteRent() gin.HandlerFunc
}
