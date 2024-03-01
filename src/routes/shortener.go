package routes

import (
	"golang-url-shortener/src/services"

	"github.com/gin-gonic/gin"
)

func AddShortenerRoutes(router *gin.RouterGroup, service services.ShortenerServiceInterface) {
	r := router.Group("/shortener")

	r.POST("", func(ctx *gin.Context) {
		// body, err := ctx.Request.Body.Read()
		// service.Create()
	})
}
