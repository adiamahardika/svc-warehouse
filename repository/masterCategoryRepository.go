package repository

import (
	"fmt"
	"svc-warehouse/model"
)

type MasterCategoryRepostoryInterface interface {
	CreateMasterCategory(request *model.MasterCategory) ([]model.MasterCategory, error)
	ReadMasterCategory() ([]model.MasterCategory, error)
}

func (repo *repository) CreateMasterCategory(request *model.MasterCategory) ([]model.MasterCategory, error) {
	var master_category []model.MasterCategory

	error := repo.db.Table("master_category").Create(request).Last(&master_category).Error

	return master_category, error
}

func (repo *repository) ReadMasterCategory() ([]model.MasterCategory, error) {
	var master_category []model.MasterCategory

	error := repo.db.Table("master_category").Where("is_active = ?", 1).Order("name").Find(&master_category).Error

	fmt.Println(master_category)
	return master_category, error
}
