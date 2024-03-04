package routes

import (
	"errors"
	"golang-url-shortener/src/common"
	"golang-url-shortener/src/services"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

var validate = validator.New()

type CreateShortenerRequest struct {
	Value string `json:"value" validate:"required,url"`
}

type CreateShortenerResult struct {
	ID string `json:"id"`
}

func AddShortenerRoutes(router *gin.RouterGroup, service services.ShortenerServiceInterface) {
	r := router.Group("/shortener")

	r.POST("", func(ctx *gin.Context) {
		request := &CreateShortenerRequest{}
		if err := ctx.ShouldBindJSON(request); err != nil {
			ctx.JSON(400, "Invalid request")
			return
		}
		err := validate.Struct(request)

		if err != nil {
			ctx.JSON(400, NewErrorResponse("Invalid request "+err.Error()))
			return
		}

		res, err := service.Create(request.Value)

		if err != nil {
			switch errors.Unwrap(err) {
			case common.ErrDuplicate:
				ctx.JSON(422, NewErrorResponse("Duplicate"))
				return
			default:
				ctx.JSON(500, NewErrorResponse("Failed while creation of short link"))
				return
			}
		}

		ctx.JSON(
			200,
			NewResponse(false, &CreateShortenerResult{
				ID: res.ID,
			}),
		)
	})
}
