package repository

import "svc-warehouse/model"

type MasterCategoryRepostoryInterface interface {
	CreateMasterCategory(request *model.MasterCategory) ([]model.MasterCategory, error)
}

func (repo *repository) CreateMasterCategory(request *model.MasterCategory) ([]model.MasterCategory, error) {
	var master_category []model.MasterCategory

	error := repo.db.Table("master_category").Create(request).Last(&master_category).Error

	return master_category, error
}
