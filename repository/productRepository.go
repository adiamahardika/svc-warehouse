package repository

import (
	"svc-warehouse/model"

	"gorm.io/gorm"
)

type ProductRepostoryInterface interface {
	CreateProduct(request *model.Product, db *gorm.DB) (model.Product, error)
	ReadProduct() ([]model.Product, error)
}

func (repo *repository) CreateProduct(request *model.Product, db *gorm.DB) (model.Product, error) {
	var inbound model.Product

	error := db.Table("product").Create(request).Last(&inbound).Error

	return inbound, error
}

func (repo *repository) ReadProduct() ([]model.Product, error) {
	var inbound []model.Product

	error := repo.db.Table("product").Select("product.*, master_product.name, product_status.name as product_status").Joins("LEFT JOIN master_product on product.master_product_id = master_product.id").Joins("LEFT JOIN product_status on product.product_status_id = product_status.id").Order("master_product.name").Scan(&inbound).Error

	return inbound, error
}
