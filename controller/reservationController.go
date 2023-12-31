package controller

import (
	"fmt"
	"net/http"
	"strconv"
	"svc-warehouse/model"
	"svc-warehouse/service"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

type reservationController struct {
	reservationService service.ReservationServiceInterface
	db                 *gorm.DB
}

func ReservationController(reservationService service.ReservationServiceInterface, db *gorm.DB) *reservationController {
	return &reservationController{reservationService, db}
}

func (controller *reservationController) CreateReservation(context *gin.Context) {

	var request *model.ReservationRequest

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
		context.JSON(httpStatus, model.ReservationResponse{
			StandardResponse: *standardResponse,
		})

	} else {

		reservation, error := controller.reservationService.CreateReservation(request, controller.db)

		if error == nil {

			description = append(description, "Success")
			httpStatus = http.StatusOK
			standardResponse = &model.StandardResponse{
				HttpStatus:  httpStatus,
				Description: description,
			}
			context.JSON(httpStatus, model.ReservationResponse{
				StandardResponse: *standardResponse,
				Result:           reservation,
			})

		} else {

			description = append(description, error.Error())

			httpStatus = http.StatusBadRequest
			standardResponse = &model.StandardResponse{
				HttpStatus:  httpStatus,
				Description: description,
			}
			context.JSON(httpStatus, model.ReservationResponse{
				StandardResponse: *standardResponse,
				Result:           reservation,
			})
		}
	}
}

func (controller *reservationController) ReadReservation(context *gin.Context) {

	description := []string{}
	httpStatus := http.StatusOK
	var standardResponse *model.StandardResponse

	reservation, error := controller.reservationService.ReadReservation()

	if error == nil {

		description = append(description, "Success")
		httpStatus = http.StatusOK
		standardResponse = &model.StandardResponse{
			HttpStatus:  httpStatus,
			Description: description,
		}
		context.JSON(httpStatus, model.GetReservationReponse{
			StandardResponse: *standardResponse,
			Result:           reservation,
		})

	} else {

		description = append(description, error.Error())
		httpStatus = http.StatusBadRequest

		standardResponse = &model.StandardResponse{
			HttpStatus:  httpStatus,
			Description: description,
		}
		context.JSON(httpStatus, model.GetReservationReponse{
			StandardResponse: *standardResponse,
			Result:           reservation,
		})

	}
}

func (controller *reservationController) ReadReservationById(context *gin.Context) {

	description := []string{}
	httpStatus := http.StatusOK
	var standardResponse *model.StandardResponse

	ids := context.Param("id")
	id, error := strconv.Atoi(ids)

	if error != nil {

		errorMessage := "Id parameter must be an integer"
		description = append(description, errorMessage)

		httpStatus = http.StatusBadRequest
		standardResponse = &model.StandardResponse{
			HttpStatus:  httpStatus,
			Description: description,
		}
		context.JSON(httpStatus, model.ReservationResponse{
			StandardResponse: *standardResponse,
		})

	}

	reservation, error := controller.reservationService.ReadReservationById(id)

	if error == nil {

		description = append(description, "Success")
		httpStatus = http.StatusOK
		standardResponse = &model.StandardResponse{
			HttpStatus:  httpStatus,
			Description: description,
		}
		context.JSON(httpStatus, model.ReservationResponse{
			StandardResponse: *standardResponse,
			Result:           reservation,
		})

	} else {

		description = append(description, error.Error())
		httpStatus = http.StatusBadRequest

		standardResponse = &model.StandardResponse{
			HttpStatus:  httpStatus,
			Description: description,
		}
		context.JSON(httpStatus, model.ReservationResponse{
			StandardResponse: *standardResponse,
		})
	}

}
