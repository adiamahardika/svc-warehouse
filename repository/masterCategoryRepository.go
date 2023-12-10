package repository

import (
	"fmt"
	"svc-warehouse/model"
)

type MasterCategoryRepostoryInterface interface {
	CreateMasterCategory(request *model.MasterCategory) ([]model.MasterCategory, error)
	ReadMasterCategory() ([]model.MasterCategory, error)
	UpdateMasterCategory(id int, request *model.MasterCategory) error
	ReadDetailMasterCategory(id int) ([]model.MasterCategory, error)
}

func (repo *repository) CreateMasterCategory(request *model.MasterCategory) ([]model.MasterCategory, error) {
	var master_category []model.MasterCategory

	error := repo.db.Table("master_category").Create(request).Last(&master_category).Error

	return master_category, error
}

func (repo *repository) ReadMasterCategory() ([]model.MasterCategory, error) {
	var master_category []model.MasterCategory

	error := repo.db.Table("master_category").Where("is_active = ?", 1).Order("name").Find(&master_category).Error

	return master_category, error
}

func (repo *repository) UpdateMasterCategory(id int, request *model.MasterCategory) error {
	var master_category []model.MasterCategory

	query := fmt.Sprintf("UPDATE master_category SET name = @Name, updated_at = @UpdatedAt WHERE id = %d", id)
	error := repo.db.Raw(query, request).Scan(&master_category).Error

	return error
}

func (repo *repository) ReadDetailMasterCategory(id int) ([]model.MasterCategory, error) {
	var master_category []model.MasterCategory

	error := repo.db.Table("master_category").Where("id = ?", id).Find(&master_category).Error

	return master_category, error
}
