package repository

import (
	"fmt"
	"svc-warehouse/model"
)

type ProductStatusRepostoryInterface interface {
	CreateProductStatus(request *model.ProductStatus) ([]model.ProductStatus, error)
	ReadProductStatus() ([]model.ProductStatus, error)
	UpdateProductStatus(id int, request *model.ProductStatus) error
	ReadDetailProductStatus(id int) ([]model.ProductStatus, error)
	DeleteProductStatus(id int) error
}

func (repo *repository) CreateProductStatus(request *model.ProductStatus) ([]model.ProductStatus, error) {
	var productStatus []model.ProductStatus

	error := repo.db.Table("product_status").Create(request).Last(&productStatus).Error

	return productStatus, error
}

func (repo *repository) ReadProductStatus() ([]model.ProductStatus, error) {
	var productStatus []model.ProductStatus

	error := repo.db.Table("product_status").Where("is_active = ?", 1).Order("name").Find(&productStatus).Error

	return productStatus, error
}

func (repo *repository) UpdateProductStatus(id int, request *model.ProductStatus) error {
	var productStatus []model.ProductStatus

	query := fmt.Sprintf("UPDATE product_status SET name = @Name, updated_at = @UpdatedAt WHERE id = %d", id)
	error := repo.db.Raw(query, request).Scan(&productStatus).Error

	return error
}

func (repo *repository) ReadDetailProductStatus(id int) ([]model.ProductStatus, error) {
	var productStatus []model.ProductStatus

	error := repo.db.Table("product_status").Where("id = ?", id).Find(&productStatus).Error

	return productStatus, error
}

func (repo *repository) DeleteProductStatus(id int) error {
	var productStatus []model.ProductStatus

	query := fmt.Sprintf("UPDATE product_status SET is_active = 0 WHERE id = %d", id)
	error := repo.db.Raw(query).Scan(&productStatus).Error

	return error
}
