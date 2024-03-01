package routes

import "github.com/gin-gonic/gin"

func AddRedirectRoutes(router *gin.RouterGroup) {
	r := router.Group("/r/")

	r.GET("", func(ctx *gin.Context) {
		// path := ctx.Request.URL.Path
	})
}
