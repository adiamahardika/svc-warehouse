package service

import (
	"svc-warehouse/model"
	"svc-warehouse/repository"
	"time"
)

type ProductStatusServiceInterface interface {
	CreateProductStatus(request *model.ProductStatus) ([]model.ProductStatus, error)
	ReadProductStatus() ([]model.ProductStatus, error)
	UpdateProductStatus(id int, request *model.ProductStatus) ([]model.ProductStatus, error)
	DeleteProductStatus(id int) error
}

type productStatusService struct {
	repository repository.ProductStatusRepostoryInterface
}

func ProductStatusService(repository repository.ProductStatusRepostoryInterface) *productStatusService {
	return &productStatusService{repository}
}

func (service *productStatusService) CreateProductStatus(request *model.ProductStatus) ([]model.ProductStatus, error) {
	now := time.Now()
	request.CreatedAt = now
	request.UpdatedAt = now
	request.IsActive = 1

	productStatus, error := service.repository.CreateProductStatus(request)

	return productStatus, error
}

func (service *productStatusService) ReadProductStatus() ([]model.ProductStatus, error) {

	productStatus, error := service.repository.ReadProductStatus()

	return productStatus, error
}

func (service *productStatusService) UpdateProductStatus(id int, request *model.ProductStatus) ([]model.ProductStatus, error) {
	now := time.Now()
	request.UpdatedAt = now
	productStatus := []model.ProductStatus{}

	error := service.repository.UpdateProductStatus(id, request)
	if error == nil {

		productStatus, error = service.repository.ReadDetailProductStatus(id)
		if error == nil {
			return productStatus, error
		}
	}

	return productStatus, error
}

func (service *productStatusService) DeleteProductStatus(id int) error {

	error := service.repository.DeleteProductStatus(id)

	return error
}
