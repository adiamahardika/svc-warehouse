package repository

import (
	"svc-warehouse/model"

	"gorm.io/gorm"
)

type ProductRepostoryInterface interface {
	CreateProduct(request *model.Product, db *gorm.DB) (model.Product, error)
}

func (repo *repository) CreateProduct(request *model.Product, db *gorm.DB) (model.Product, error) {
	var inbound model.Product

	error := db.Table("product").Create(request).Last(&inbound).Error

	return inbound, error
}
