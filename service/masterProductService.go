package service

import (
	"svc-warehouse/model"
	"svc-warehouse/repository"
	"time"
)

type MasterProductServiceInterface interface {
	CreateMasterProduct(request *model.MasterProduct) ([]model.MasterProduct, error)
	ReadMasterProduct() ([]model.MasterProduct, error)
	UpdateMasterProduct(id int, request *model.MasterProduct) ([]model.MasterProduct, error)
	DeleteMasterProduct(id int) error
}

type masterProductService struct {
	repository repository.MasterProductRepostoryInterface
}

func MasterProductService(repository repository.MasterProductRepostoryInterface) *masterProductService {
	return &masterProductService{repository}
}

func (service *masterProductService) CreateMasterProduct(request *model.MasterProduct) ([]model.MasterProduct, error) {
	now := time.Now()
	request.CreatedAt = now
	request.UpdatedAt = now
	request.IsActive = 1

	masterProduct, error := service.repository.CreateMasterProduct(request)

	return masterProduct, error
}

func (service *masterProductService) ReadMasterProduct() ([]model.MasterProduct, error) {

	masterProduct, error := service.repository.ReadMasterProduct()

	return masterProduct, error
}

func (service *masterProductService) UpdateMasterProduct(id int, request *model.MasterProduct) ([]model.MasterProduct, error) {
	now := time.Now()
	request.UpdatedAt = now
	masterProduct := []model.MasterProduct{}

	error := service.repository.UpdateMasterProduct(id, request)
	if error == nil {

		masterProduct, error = service.repository.ReadDetailMasterProduct(id)
		if error == nil {
			return masterProduct, error
		}
	}

	return masterProduct, error
}

func (service *masterProductService) DeleteMasterProduct(id int) error {

	error := service.repository.DeleteMasterProduct(id)

	return error
}
