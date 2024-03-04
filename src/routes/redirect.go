package routes

import (
	"golang-url-shortener/src/services"

	"github.com/gin-gonic/gin"
)

func AddRedirectRoutes(router *gin.RouterGroup, service services.ShortenerServiceInterface) {
	r := router.Group("/r")

	r.GET(":key", func(ctx *gin.Context) {
		key := ctx.Param("key")

		shortLink, err := service.Get(key)

		if err != nil {
			ctx.JSON(500, NewErrorResponse("Failed while check short link in storage"))
			return
		}

		if shortLink == nil {
			ctx.JSON(400, NewErrorResponse("Not found"))
			return
		}

		ctx.Redirect(302, *shortLink.Value)
	})
}
