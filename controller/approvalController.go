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

type approvalController struct {
	approvalService service.ApprovalServiceInterface
	db              *gorm.DB
}

func ApprovalController(approvalService service.ApprovalServiceInterface, db *gorm.DB) *approvalController {
	return &approvalController{approvalService, db}
}

func (controller *approvalController) CreateApproval(context *gin.Context) {

	var request *model.ApprovalRequest

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
		context.JSON(httpStatus, model.ApprovalResponse{
			StandardResponse: *standardResponse,
		})

	} else {

		approval, error := controller.approvalService.CreateApproval(request, controller.db)

		if error == nil {

			description = append(description, "Success")
			httpStatus = http.StatusOK
			standardResponse = &model.StandardResponse{
				HttpStatus:  httpStatus,
				Description: description,
			}
			context.JSON(httpStatus, model.ApprovalResponse{
				StandardResponse: *standardResponse,
				Result:           approval,
			})

		} else {

			description = append(description, error.Error())

			httpStatus = http.StatusBadRequest
			standardResponse = &model.StandardResponse{
				HttpStatus:  httpStatus,
				Description: description,
			}
			context.JSON(httpStatus, model.ApprovalResponse{
				StandardResponse: *standardResponse,
				Result:           approval,
			})
		}
	}
}

func (controller *approvalController) ReadApproval(context *gin.Context) {

	description := []string{}
	httpStatus := http.StatusOK
	var standardResponse *model.StandardResponse

	approval, error := controller.approvalService.ReadApproval()

	if error == nil {

		description = append(description, "Success")
		httpStatus = http.StatusOK
		standardResponse = &model.StandardResponse{
			HttpStatus:  httpStatus,
			Description: description,
		}
		context.JSON(httpStatus, model.GetApprovalReponse{
			StandardResponse: *standardResponse,
			Result:           approval,
		})

	} else {

		description = append(description, error.Error())
		httpStatus = http.StatusBadRequest

		standardResponse = &model.StandardResponse{
			HttpStatus:  httpStatus,
			Description: description,
		}
		context.JSON(httpStatus, model.GetApprovalReponse{
			StandardResponse: *standardResponse,
			Result:           approval,
		})

	}
}

func (controller *approvalController) ReadApprovalById(context *gin.Context) {

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
		context.JSON(httpStatus, model.ApprovalResponse{
			StandardResponse: *standardResponse,
		})

	}

	approval, error := controller.approvalService.ReadApprovalById(id)

	if error == nil {

		description = append(description, "Success")
		httpStatus = http.StatusOK
		standardResponse = &model.StandardResponse{
			HttpStatus:  httpStatus,
			Description: description,
		}
		context.JSON(httpStatus, model.ApprovalResponse{
			StandardResponse: *standardResponse,
			Result:           approval,
		})

	} else {

		description = append(description, error.Error())
		httpStatus = http.StatusBadRequest

		standardResponse = &model.StandardResponse{
			HttpStatus:  httpStatus,
			Description: description,
		}
		context.JSON(httpStatus, model.ApprovalResponse{
			StandardResponse: *standardResponse,
		})
	}

}
