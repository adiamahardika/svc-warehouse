package repository

import "svc-warehouse/model"

type MasterProductRepostoryInterface interface {
	CreateMasterProduct(request *model.MasterProduct) ([]model.MasterProduct, error)
}

func (repo *repository) CreateMasterProduct(request *model.MasterProduct) ([]model.MasterProduct, error) {
	var master_Product []model.MasterProduct

	error := repo.db.Table("master_product").Create(request).Last(&master_Product).Error

	return master_Product, error
}
