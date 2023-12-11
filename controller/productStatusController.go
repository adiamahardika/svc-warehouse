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

type productStatusController struct {
	productStatusService service.ProductStatusServiceInterface
}

func ProductStatusController(productStatusService service.ProductStatusServiceInterface) *productStatusController {
	return &productStatusController{productStatusService}
}

func (controller *productStatusController) CreateProductStatus(context *gin.Context) {

	var request *model.ProductStatus

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
		context.JSON(httpStatus, model.ProductStatusResponse{
			StandardResponse: *standardResponse,
		})

	} else {

		productStatus, error := controller.productStatusService.CreateProductStatus(request)

		if error == nil {

			description = append(description, "Success")
			httpStatus = http.StatusOK
			standardResponse = &model.StandardResponse{
				HttpStatus:  httpStatus,
				Description: description,
			}
			context.JSON(httpStatus, model.ProductStatusResponse{
				StandardResponse: *standardResponse,
				Result:           productStatus,
			})

		} else {

			description = append(description, error.Error())
			httpStatus = http.StatusBadRequest

			standardResponse = &model.StandardResponse{
				HttpStatus:  httpStatus,
				Description: description,
			}
			context.JSON(httpStatus, model.ProductStatusResponse{
				StandardResponse: *standardResponse,
				Result:           productStatus,
			})
		}
	}
}

func (controller *productStatusController) ReadProductStatus(context *gin.Context) {

	description := []string{}
	httpStatus := http.StatusOK
	var standardResponse *model.StandardResponse

	productStatus, error := controller.productStatusService.ReadProductStatus()

	if error == nil {

		description = append(description, "Success")
		httpStatus = http.StatusOK
		standardResponse = &model.StandardResponse{
			HttpStatus:  httpStatus,
			Description: description,
		}
		context.JSON(httpStatus, model.ProductStatusResponse{
			StandardResponse: *standardResponse,
			Result:           productStatus,
		})

	} else {

		description = append(description, error.Error())
		httpStatus = http.StatusBadRequest

		standardResponse = &model.StandardResponse{
			HttpStatus:  httpStatus,
			Description: description,
		}
		context.JSON(httpStatus, model.ProductStatusResponse{
			StandardResponse: *standardResponse,
			Result:           productStatus,
		})

	}
}

func (controller *productStatusController) UpdateProductStatus(context *gin.Context) {

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
		context.JSON(httpStatus, model.ProductStatusResponse{
			StandardResponse: *standardResponse,
		})

	}

	var request *model.ProductStatus
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
		context.JSON(httpStatus, model.ProductStatusResponse{
			StandardResponse: *standardResponse,
		})

	} else {

		productStatus, error := controller.productStatusService.UpdateProductStatus(id, request)

		if error == nil {

			description = append(description, "Success")
			httpStatus = http.StatusOK
			standardResponse = &model.StandardResponse{
				HttpStatus:  httpStatus,
				Description: description,
			}
			context.JSON(httpStatus, model.ProductStatusResponse{
				StandardResponse: *standardResponse,
				Result:           productStatus,
			})

		} else {

			description = append(description, error.Error())
			httpStatus = http.StatusBadRequest

			standardResponse = &model.StandardResponse{
				HttpStatus:  httpStatus,
				Description: description,
			}
			context.JSON(httpStatus, model.ProductStatusResponse{
				StandardResponse: *standardResponse,
				Result:           productStatus,
			})
		}
	}
}

func (controller *productStatusController) DeleteProductStatus(context *gin.Context) {

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
		context.JSON(httpStatus, model.ProductStatusResponse{
			StandardResponse: *standardResponse,
		})

	}

	error = controller.productStatusService.DeleteProductStatus(id)

	if error == nil {

		description = append(description, "Success")
		httpStatus = http.StatusOK
		standardResponse = &model.StandardResponse{
			HttpStatus:  httpStatus,
			Description: description,
		}
		context.JSON(httpStatus, model.ProductStatusResponse{
			StandardResponse: *standardResponse,
		})

	} else {

		description = append(description, error.Error())
		httpStatus = http.StatusBadRequest

		standardResponse = &model.StandardResponse{
			HttpStatus:  httpStatus,
			Description: description,
		}
		context.JSON(httpStatus, model.ProductStatusResponse{
			StandardResponse: *standardResponse,
		})
	}

}
