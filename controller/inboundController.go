package controller

import (
	"fmt"
	"net/http"
	"svc-warehouse/model"
	"svc-warehouse/service"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

type inboundController struct {
	inboundService service.InboundServiceInterface
	db             *gorm.DB
}

func InboundController(inboundService service.InboundServiceInterface, db *gorm.DB) *inboundController {
	return &inboundController{inboundService, db}
}

func (controller *inboundController) CreateInbound(context *gin.Context) {

	var request *model.InboundRequest

	error := context.ShouldBind(&request)
	description := []string{}
	httpStatus := http.StatusOK
	var standardResponse *model.StandardResponse

	if error != nil {

		for _, value := range error.(validator.ValidationErrors) {
			errorMessage := fmt.Sprintf("Error on field %s, condition: %s", value.Field(), value.ActualTag())
			description = append(description, errorMessage)
		}

		httpStatus = http.StatusBadRequest
		standardResponse = &model.StandardResponse{
			HttpStatus:  httpStatus,
			Description: description,
		}
		context.JSON(httpStatus, model.InboundResponse{
			StandardResponse: *standardResponse,
		})

	} else {

		inbound, error := controller.inboundService.CreateInbound(request, controller.db)

		if error == nil {

			description = append(description, "Success")
			httpStatus = http.StatusOK
			standardResponse = &model.StandardResponse{
				HttpStatus:  httpStatus,
				Description: description,
			}
			context.JSON(httpStatus, model.InboundResponse{
				StandardResponse: *standardResponse,
				Result:           inbound,
			})

		} else {

			description = append(description, error.Error())
			httpStatus = http.StatusBadRequest

			standardResponse = &model.StandardResponse{
				HttpStatus:  httpStatus,
				Description: description,
			}
			context.JSON(httpStatus, model.InboundResponse{
				StandardResponse: *standardResponse,
				Result:           inbound,
			})
		}
	}
}
