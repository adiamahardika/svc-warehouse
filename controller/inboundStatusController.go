package controller

import (
	"fmt"
	"net/http"
	"strconv"
	"svc-warehouse/model"
	"svc-warehouse/service"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type reservationStatusController struct {
	reservationStatusService service.ReservationStatusServiceInterface
}

func ReservationStatusController(reservationStatusService service.ReservationStatusServiceInterface) *reservationStatusController {
	return &reservationStatusController{reservationStatusService}
}

func (controller *reservationStatusController) CreateReservationStatus(context *gin.Context) {

	var request *model.ReservationStatus

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
		context.JSON(httpStatus, model.ReservationStatusResponse{
			StandardResponse: *standardResponse,
		})

	} else {

		reservationStatus, error := controller.reservationStatusService.CreateReservationStatus(request)

		if error == nil {

			description = append(description, "Success")
			httpStatus = http.StatusOK
			standardResponse = &model.StandardResponse{
				HttpStatus:  httpStatus,
				Description: description,
			}
			context.JSON(httpStatus, model.ReservationStatusResponse{
				StandardResponse: *standardResponse,
				Result:           reservationStatus,
			})

		} else {

			description = append(description, error.Error())
			httpStatus = http.StatusBadRequest

			standardResponse = &model.StandardResponse{
				HttpStatus:  httpStatus,
				Description: description,
			}
			context.JSON(httpStatus, model.ReservationStatusResponse{
				StandardResponse: *standardResponse,
				Result:           reservationStatus,
			})
		}
	}
}

func (controller *reservationStatusController) ReadReservationStatus(context *gin.Context) {

	description := []string{}
	httpStatus := http.StatusOK
	var standardResponse *model.StandardResponse

	reservationStatus, error := controller.reservationStatusService.ReadReservationStatus()

	if error == nil {

		description = append(description, "Success")
		httpStatus = http.StatusOK
		standardResponse = &model.StandardResponse{
			HttpStatus:  httpStatus,
			Description: description,
		}
		context.JSON(httpStatus, model.ReservationStatusResponse{
			StandardResponse: *standardResponse,
			Result:           reservationStatus,
		})

	} else {

		description = append(description, error.Error())
		httpStatus = http.StatusBadRequest

		standardResponse = &model.StandardResponse{
			HttpStatus:  httpStatus,
			Description: description,
		}
		context.JSON(httpStatus, model.ReservationStatusResponse{
			StandardResponse: *standardResponse,
			Result:           reservationStatus,
		})

	}
}

func (controller *reservationStatusController) UpdateReservationStatus(context *gin.Context) {

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
		context.JSON(httpStatus, model.ReservationStatusResponse{
			StandardResponse: *standardResponse,
		})

	}

	var request *model.ReservationStatus
	error = context.ShouldBind(&request)

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
		context.JSON(httpStatus, model.ReservationStatusResponse{
			StandardResponse: *standardResponse,
		})

	} else {

		reservationStatus, error := controller.reservationStatusService.UpdateReservationStatus(id, request)

		if error == nil {

			description = append(description, "Success")
			httpStatus = http.StatusOK
			standardResponse = &model.StandardResponse{
				HttpStatus:  httpStatus,
				Description: description,
			}
			context.JSON(httpStatus, model.ReservationStatusResponse{
				StandardResponse: *standardResponse,
				Result:           reservationStatus,
			})

		} else {

			description = append(description, error.Error())
			httpStatus = http.StatusBadRequest

			standardResponse = &model.StandardResponse{
				HttpStatus:  httpStatus,
				Description: description,
			}
			context.JSON(httpStatus, model.ReservationStatusResponse{
				StandardResponse: *standardResponse,
				Result:           reservationStatus,
			})
		}
	}
}

func (controller *reservationStatusController) DeleteReservationStatus(context *gin.Context) {

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
		context.JSON(httpStatus, model.ReservationStatusResponse{
			StandardResponse: *standardResponse,
		})

	}

	error = controller.reservationStatusService.DeleteReservationStatus(id)

	if error == nil {

		description = append(description, "Success")
		httpStatus = http.StatusOK
		standardResponse = &model.StandardResponse{
			HttpStatus:  httpStatus,
			Description: description,
		}
		context.JSON(httpStatus, model.ReservationStatusResponse{
			StandardResponse: *standardResponse,
		})

	} else {

		description = append(description, error.Error())
		httpStatus = http.StatusBadRequest

		standardResponse = &model.StandardResponse{
			HttpStatus:  httpStatus,
			Description: description,
		}
		context.JSON(httpStatus, model.ReservationStatusResponse{
			StandardResponse: *standardResponse,
		})
	}

}
