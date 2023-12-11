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

type masterProductController struct {
	masterProductService service.MasterProductServiceInterface
}

func MasterProductController(masterProductService service.MasterProductServiceInterface) *masterProductController {
	return &masterProductController{masterProductService}
}

func (controller *masterProductController) CreateMasterProduct(context *gin.Context) {

	var request *model.MasterProduct

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
		context.JSON(httpStatus, model.MasterProductResponse{
			StandardResponse: *standardResponse,
		})

	} else {

		masterProduct, error := controller.masterProductService.CreateMasterProduct(request)

		if error == nil {

			description = append(description, "Success")
			httpStatus = http.StatusOK
			standardResponse = &model.StandardResponse{
				HttpStatus:  httpStatus,
				Description: description,
			}
			context.JSON(httpStatus, model.MasterProductResponse{
				StandardResponse: *standardResponse,
				Result:           masterProduct,
			})

		} else {

			description = append(description, error.Error())
			httpStatus = http.StatusBadRequest

			standardResponse = &model.StandardResponse{
				HttpStatus:  httpStatus,
				Description: description,
			}
			context.JSON(httpStatus, model.MasterProductResponse{
				StandardResponse: *standardResponse,
				Result:           masterProduct,
			})
		}
	}
}

func (controller *masterProductController) ReadMasterProduct(context *gin.Context) {

	description := []string{}
	httpStatus := http.StatusOK
	var standardResponse *model.StandardResponse

	masterProduct, error := controller.masterProductService.ReadMasterProduct()

	if error == nil {

		description = append(description, "Success")
		httpStatus = http.StatusOK
		standardResponse = &model.StandardResponse{
			HttpStatus:  httpStatus,
			Description: description,
		}
		context.JSON(httpStatus, model.MasterProductResponse{
			StandardResponse: *standardResponse,
			Result:           masterProduct,
		})

	} else {

		description = append(description, error.Error())
		httpStatus = http.StatusBadRequest

		standardResponse = &model.StandardResponse{
			HttpStatus:  httpStatus,
			Description: description,
		}
		context.JSON(httpStatus, model.MasterProductResponse{
			StandardResponse: *standardResponse,
			Result:           masterProduct,
		})

	}
}

func (controller *masterProductController) UpdateMasterProduct(context *gin.Context) {

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
		context.JSON(httpStatus, model.MasterProductResponse{
			StandardResponse: *standardResponse,
		})

	}

	var request *model.MasterProduct
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
		context.JSON(httpStatus, model.MasterProductResponse{
			StandardResponse: *standardResponse,
		})

	} else {

		masterProduct, error := controller.masterProductService.UpdateMasterProduct(id, request)

		if error == nil {

			description = append(description, "Success")
			httpStatus = http.StatusOK
			standardResponse = &model.StandardResponse{
				HttpStatus:  httpStatus,
				Description: description,
			}
			context.JSON(httpStatus, model.MasterProductResponse{
				StandardResponse: *standardResponse,
				Result:           masterProduct,
			})

		} else {

			description = append(description, error.Error())
			httpStatus = http.StatusBadRequest

			standardResponse = &model.StandardResponse{
				HttpStatus:  httpStatus,
				Description: description,
			}
			context.JSON(httpStatus, model.MasterProductResponse{
				StandardResponse: *standardResponse,
				Result:           masterProduct,
			})
		}
	}
}

func (controller *masterProductController) DeleteMasterProduct(context *gin.Context) {

	description := []string{}
	http_status := http.StatusOK
	var standardResponse *model.StandardResponse

	ids := context.Param("id")
	id, error := strconv.Atoi(ids)

	if error != nil {

		errorMessage := "Id parameter must be an integer"
		description = append(description, errorMessage)

		http_status = http.StatusBadRequest
		standardResponse = &model.StandardResponse{
			HttpStatus:  http_status,
			Description: description,
		}
		context.JSON(http_status, model.MasterProductResponse{
			StandardResponse: *standardResponse,
		})

	}

	error = controller.masterProductService.DeleteMasterProduct(id)

	if error == nil {

		description = append(description, "Success")
		http_status = http.StatusOK
		standardResponse = &model.StandardResponse{
			HttpStatus:  http_status,
			Description: description,
		}
		context.JSON(http_status, model.MasterProductResponse{
			StandardResponse: *standardResponse,
		})

	} else {

		description = append(description, error.Error())
		http_status = http.StatusBadRequest

		standardResponse = &model.StandardResponse{
			HttpStatus:  http_status,
			Description: description,
		}
		context.JSON(http_status, model.MasterProductResponse{
			StandardResponse: *standardResponse,
		})
	}

}
