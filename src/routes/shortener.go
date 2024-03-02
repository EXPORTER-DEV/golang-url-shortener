package routes

import (
	"golang-url-shortener/src/services"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

var validate = validator.New()

type CreateShortenerRequest struct {
	Value string `json:"value" validate:"required"`
}

func AddShortenerRoutes(router *gin.RouterGroup, service services.ShortenerServiceInterface) {
	r := router.Group("/shortener")

	r.POST("", func(ctx *gin.Context) {
		request := &CreateShortenerRequest{}
		if err := ctx.BindJSON(request); err != nil {
			ctx.JSON(400, "Invalid request")
			return
		}
		err := validate.Struct(request)

		if err != nil {
			ctx.JSON(400, "Invalid request: "+err.Error())
			return
		}

		log.Println("HERE")
		// service.Create()
	})
}
