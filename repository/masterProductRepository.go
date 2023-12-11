package repository

import (
	"fmt"
	"svc-warehouse/model"
)

type MasterProductRepostoryInterface interface {
	CreateMasterProduct(request *model.MasterProduct) ([]model.MasterProduct, error)
	ReadMasterProduct() ([]model.MasterProduct, error)
	UpdateMasterProduct(id int, request *model.MasterProduct) error
	ReadDetailMasterProduct(id int) ([]model.MasterProduct, error)
	DeleteMasterProduct(id int) error
}

func (repo *repository) CreateMasterProduct(request *model.MasterProduct) ([]model.MasterProduct, error) {
	var masterProduct []model.MasterProduct

	error := repo.db.Table("master_product").Create(request).Last(&masterProduct).Error

	return masterProduct, error
}

func (repo *repository) ReadMasterProduct() ([]model.MasterProduct, error) {
	var masterProduct []model.MasterProduct

	error := repo.db.Table("master_product").Select("master_product.*, master_category.name AS category").Joins("LEFT JOIN master_category on master_product.master_category_id = master_category.id ").Where("master_product.is_active = ?", 1).Order("master_product.name").Scan(&masterProduct).Error

	return masterProduct, error
}

func (repo *repository) UpdateMasterProduct(id int, request *model.MasterProduct) error {
	var masterProduct []model.MasterProduct

	query := fmt.Sprintf("UPDATE master_product SET name = @Name, master_category_id = @MasterCategoryId, updated_at = @UpdatedAt WHERE id = %d", id)
	error := repo.db.Raw(query, request).Scan(&masterProduct).Error

	return error
}

func (repo *repository) ReadDetailMasterProduct(id int) ([]model.MasterProduct, error) {
	var masterProduct []model.MasterProduct
	error := repo.db.Table("master_product").Select("master_product.*, master_category.name AS category").Joins("LEFT JOIN master_category on master_product.master_category_id = master_category.id ").Where("master_product.id = ?", id).Scan(&masterProduct).Error

	return masterProduct, error
}

func (repo *repository) DeleteMasterProduct(id int) error {
	var masterProduct []model.MasterProduct

	query := fmt.Sprintf("UPDATE master_product SET is_active = 0 WHERE id = %d", id)
	error := repo.db.Raw(query).Scan(&masterProduct).Error

	return error
}
