package handlers

import (
	"fmt"

	"github.com/gin-gonic/gin"
	swagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

func registerHandler(router *gin.RouterGroup, method string, path string, handler gin.HandlerFunc) {
	switch method {
	case "GET":
		router.GET(path, handler)
	case "POST":
		router.POST(path, handler)
	}
}

func RegisterSwagger(router *gin.RouterGroup) {
	registerHandler(router, "GET", "/swagger/*any", swagger.CustomWrapHandler(
		&swagger.Config{
			URL: fmt.Sprintf("%v/swagger/doc.json", router.BasePath()),
		},
		swaggerFiles.Handler,
	))
}
