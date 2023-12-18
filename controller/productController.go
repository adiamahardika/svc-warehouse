package controller

import (
	"net/http"
	"svc-warehouse/model"
	"svc-warehouse/service"

	"github.com/gin-gonic/gin"
)

type productController struct {
	productService service.ProductServiceInterface
}

func ProductController(productService service.ProductServiceInterface) *productController {
	return &productController{productService}
}

func (controller *productController) ReadProduct(context *gin.Context) {

	description := []string{}
	httpStatus := http.StatusOK
	var standardResponse *model.StandardResponse

	product, error := controller.productService.ReadProduct()

	if error == nil {

		description = append(description, "Success")
		httpStatus = http.StatusOK
		standardResponse = &model.StandardResponse{
			HttpStatus:  httpStatus,
			Description: description,
		}
		context.JSON(httpStatus, model.ProductResponse{
			StandardResponse: *standardResponse,
			Result:           product,
		})

	} else {

		description = append(description, error.Error())
		httpStatus = http.StatusBadRequest

		standardResponse = &model.StandardResponse{
			HttpStatus:  httpStatus,
			Description: description,
		}
		context.JSON(httpStatus, model.ProductResponse{
			StandardResponse: *standardResponse,
			Result:           product,
		})

	}
}
