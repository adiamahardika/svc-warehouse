package service

import (
	"svc-warehouse/model"
	"svc-warehouse/repository"
	"time"
)

type MasterCategoryServiceInterface interface {
	CreateMasterCategory(request *model.MasterCategory) ([]model.MasterCategory, error)
	ReadMasterCategory() ([]model.MasterCategory, error)
	UpdateMasterCategory(id int, request *model.MasterCategory) ([]model.MasterCategory, error)
}

type masterCategoryService struct {
	repository repository.MasterCategoryRepostoryInterface
}

func MasterCategoryService(repository repository.MasterCategoryRepostoryInterface) *masterCategoryService {
	return &masterCategoryService{repository}
}

func (service *masterCategoryService) CreateMasterCategory(request *model.MasterCategory) ([]model.MasterCategory, error) {
	now := time.Now()
	request.CreatedAt = now
	request.UpdatedAt = now
	request.IsActive = 1

	masterCategory, error := service.repository.CreateMasterCategory(request)

	return masterCategory, error
}

func (service *masterCategoryService) ReadMasterCategory() ([]model.MasterCategory, error) {

	masterCategory, error := service.repository.ReadMasterCategory()

	return masterCategory, error
}

func (service *masterCategoryService) UpdateMasterCategory(id int, request *model.MasterCategory) ([]model.MasterCategory, error) {
	now := time.Now()
	request.UpdatedAt = now
	masterCategory := []model.MasterCategory{}

	error := service.repository.UpdateMasterCategory(id, request)
	if error == nil {

		masterCategory, error = service.repository.ReadDetailMasterCategory(id)
		if error == nil {
			return masterCategory, error
		}
	}

	return masterCategory, error
}
