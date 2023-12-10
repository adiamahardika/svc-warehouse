package controller

import (
	"fmt"
	"net/http"
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
	http_status := http.StatusOK
	var standardResponse *model.StandardResponse

	if error != nil {

		for _, value := range error.(validator.ValidationErrors) {
			errorMessage := fmt.Sprintf("Error on field %s, condition: %s", value.Field(), value.ActualTag())
			description = append(description, errorMessage)
		}

		http_status = http.StatusBadRequest
		standardResponse = &model.StandardResponse{
			HttpStatus:  http_status,
			Description: description,
		}
		context.JSON(http_status, model.MasterCategoryResponse{
			StandardResponse: *standardResponse,
		})

	} else {

		master_category, error := controller.masterCategoryService.CreateMasterCategory(request)

		if error == nil {

			description = append(description, "Success")
			http_status = http.StatusOK
			standardResponse = &model.StandardResponse{
				HttpStatus:  http_status,
				Description: description,
			}
			context.JSON(http_status, model.MasterCategoryResponse{
				StandardResponse: *standardResponse,
				Result:           master_category,
			})

		} else {

			description = append(description, error.Error())
			http_status = http.StatusBadRequest

			standardResponse = &model.StandardResponse{
				HttpStatus:  http_status,
				Description: description,
			}
			context.JSON(http_status, model.MasterCategoryResponse{
				StandardResponse: *standardResponse,
				Result:           master_category,
			})
		}
	}
}
