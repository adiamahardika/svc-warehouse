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

type masterCategoryController struct {
	masterCategoryService service.MasterCategoryServiceInterface
}

func MasterCategoryController(masterCategoryService service.MasterCategoryServiceInterface) *masterCategoryController {
	return &masterCategoryController{masterCategoryService}
}

func (controller *masterCategoryController) CreateMasterCategory(context *gin.Context) {

	var request *model.MasterCategory

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
		context.JSON(httpStatus, model.MasterCategoryResponse{
			StandardResponse: *standardResponse,
		})

	} else {

		masterCategory, error := controller.masterCategoryService.CreateMasterCategory(request)

		if error == nil {

			description = append(description, "Success")
			httpStatus = http.StatusOK
			standardResponse = &model.StandardResponse{
				HttpStatus:  httpStatus,
				Description: description,
			}
			context.JSON(httpStatus, model.MasterCategoryResponse{
				StandardResponse: *standardResponse,
				Result:           masterCategory,
			})

		} else {

			description = append(description, error.Error())
			httpStatus = http.StatusBadRequest

			standardResponse = &model.StandardResponse{
				HttpStatus:  httpStatus,
				Description: description,
			}
			context.JSON(httpStatus, model.MasterCategoryResponse{
				StandardResponse: *standardResponse,
				Result:           masterCategory,
			})
		}
	}
}

func (controller *masterCategoryController) ReadMasterCategory(context *gin.Context) {

	description := []string{}
	httpStatus := http.StatusOK
	var standardResponse *model.StandardResponse

	masterCategory, error := controller.masterCategoryService.ReadMasterCategory()

	if error == nil {

		description = append(description, "Success")
		httpStatus = http.StatusOK
		standardResponse = &model.StandardResponse{
			HttpStatus:  httpStatus,
			Description: description,
		}
		context.JSON(httpStatus, model.MasterCategoryResponse{
			StandardResponse: *standardResponse,
			Result:           masterCategory,
		})

	} else {

		description = append(description, error.Error())
		httpStatus = http.StatusBadRequest

		standardResponse = &model.StandardResponse{
			HttpStatus:  httpStatus,
			Description: description,
		}
		context.JSON(httpStatus, model.MasterCategoryResponse{
			StandardResponse: *standardResponse,
			Result:           masterCategory,
		})

	}
}

func (controller *masterCategoryController) UpdateMasterCategory(context *gin.Context) {

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
		context.JSON(httpStatus, model.MasterCategoryResponse{
			StandardResponse: *standardResponse,
		})

	}

	var request *model.MasterCategory
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
		context.JSON(httpStatus, model.MasterCategoryResponse{
			StandardResponse: *standardResponse,
		})

	} else {

		masterCategory, error := controller.masterCategoryService.UpdateMasterCategory(id, request)

		if error == nil {

			description = append(description, "Success")
			httpStatus = http.StatusOK
			standardResponse = &model.StandardResponse{
				HttpStatus:  httpStatus,
				Description: description,
			}
			context.JSON(httpStatus, model.MasterCategoryResponse{
				StandardResponse: *standardResponse,
				Result:           masterCategory,
			})

		} else {

			description = append(description, error.Error())
			httpStatus = http.StatusBadRequest

			standardResponse = &model.StandardResponse{
				HttpStatus:  httpStatus,
				Description: description,
			}
			context.JSON(httpStatus, model.MasterCategoryResponse{
				StandardResponse: *standardResponse,
				Result:           masterCategory,
			})
		}
	}
}

func (controller *masterCategoryController) DeleteMasterCategory(context *gin.Context) {

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
		context.JSON(httpStatus, model.MasterCategoryResponse{
			StandardResponse: *standardResponse,
		})

	}

	error = controller.masterCategoryService.DeleteMasterCategory(id)

	if error == nil {

		description = append(description, "Success")
		httpStatus = http.StatusOK
		standardResponse = &model.StandardResponse{
			HttpStatus:  httpStatus,
			Description: description,
		}
		context.JSON(httpStatus, model.MasterCategoryResponse{
			StandardResponse: *standardResponse,
		})

	} else {

		description = append(description, error.Error())
		httpStatus = http.StatusBadRequest

		standardResponse = &model.StandardResponse{
			HttpStatus:  httpStatus,
			Description: description,
		}
		context.JSON(httpStatus, model.MasterCategoryResponse{
			StandardResponse: *standardResponse,
		})
	}

}
