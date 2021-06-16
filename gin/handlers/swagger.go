package handlers

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/swaggo/gin-swagger/swaggerFiles"

	ginSwagger "github.com/swaggo/gin-swagger"
)

func RegisterSwagger(router *gin.RouterGroup) {
	router.GET("/swagger/*any",
		ginSwagger.WrapHandler(swaggerFiles.Handler, ginSwagger.URL(fmt.Sprintf("%v/swagger/doc.json", router.BasePath()))),
	)
}
