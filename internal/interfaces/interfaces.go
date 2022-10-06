package interfaces

import "github.com/gin-gonic/gin"

type Handler interface {
	GetMethod() string
	GetPath() string
	Handle(ctx *gin.Context)
}
