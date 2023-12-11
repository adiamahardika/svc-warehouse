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
	DeleteMasterCategory(id int) error
}

func (repo *repository) CreateMasterCategory(request *model.MasterCategory) ([]model.MasterCategory, error) {
	var masterCategory []model.MasterCategory

	error := repo.db.Table("master_category").Create(request).Last(&masterCategory).Error

	return masterCategory, error
}

func (repo *repository) ReadMasterCategory() ([]model.MasterCategory, error) {
	var masterCategory []model.MasterCategory

	error := repo.db.Table("master_category").Where("is_active = ?", 1).Order("name").Find(&masterCategory).Error

	return masterCategory, error
}

func (repo *repository) UpdateMasterCategory(id int, request *model.MasterCategory) error {
	var masterCategory []model.MasterCategory

	query := fmt.Sprintf("UPDATE master_category SET name = @Name, updated_at = @UpdatedAt WHERE id = %d", id)
	error := repo.db.Raw(query, request).Scan(&masterCategory).Error

	return error
}

func (repo *repository) ReadDetailMasterCategory(id int) ([]model.MasterCategory, error) {
	var masterCategory []model.MasterCategory

	error := repo.db.Table("master_category").Where("id = ?", id).Find(&masterCategory).Error

	return masterCategory, error
}

func (repo *repository) DeleteMasterCategory(id int) error {
	var masterCategory []model.MasterCategory

	query := fmt.Sprintf("UPDATE master_category SET is_active = 0 WHERE id = %d", id)
	error := repo.db.Raw(query).Scan(&masterCategory).Error

	return error
}
