package repository

import "svc-warehouse/model"

type MasterProductRepostoryInterface interface {
	CreateMasterProduct(request *model.MasterProduct) ([]model.MasterProduct, error)
	ReadMasterProduct() ([]model.MasterProduct, error)
}

func (repo *repository) CreateMasterProduct(request *model.MasterProduct) ([]model.MasterProduct, error) {
	var master_Product []model.MasterProduct

	error := repo.db.Table("master_product").Create(request).Last(&master_Product).Error

	return master_Product, error
}

func (repo *repository) ReadMasterProduct() ([]model.MasterProduct, error) {
	var master_product []model.MasterProduct

	error := repo.db.Table("master_product").Select("master_product.*, master_category.name AS category").Joins("LEFT JOIN master_category on master_product.master_category_id = master_category.id ").Where("master_product.is_active = ?", 1).Order("master_product.name").Scan(&master_product).Error

	return master_product, error
}
