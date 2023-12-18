package service

import (
	"svc-warehouse/model"
	"svc-warehouse/repository"
)

type ProductServiceInterface interface {
	ReadProduct() ([]model.Product, error)
}

type productService struct {
	repository repository.ProductRepostoryInterface
}

func ProductService(repository repository.ProductRepostoryInterface) *productService {
	return &productService{repository}
}

func (service *productService) ReadProduct() ([]model.Product, error) {

	product, error := service.repository.ReadProduct()

	return product, error
}
