package repository

import (
	"svc-warehouse/model"

	"gorm.io/gorm"
)

type ApprovalDetailRepostoryInterface interface {
	CreateApprovalDetail(request *model.ApprovalDetail, db *gorm.DB) (model.ApprovalDetail, error)
	ReadApprovalDetailByApprovalId(approvalId int) ([]model.ApprovalDetail, error)
}

func (repo *repository) CreateApprovalDetail(request *model.ApprovalDetail, db *gorm.DB) (model.ApprovalDetail, error) {
	var approval model.ApprovalDetail

	error := db.Table("approval_detail").Create(request).Last(&approval).Error

	return approval, error
}

func (repo *repository) ReadApprovalDetailByApprovalId(approvalId int) ([]model.ApprovalDetail, error) {
	var approval []model.ApprovalDetail

	error := repo.db.Table("approval_detail").Select("approval_detail.*, master_product.name as product, product.serial_number, product_status.name as product_status, master_category.name as category").Joins("LEFT JOIN product on approval_detail.product_id = product.id").Joins("LEFT JOIN master_product on product.master_product_id = master_product.id").Joins("LEFT JOIN product_status on product.product_status_id = product_status.id").Joins("LEFT JOIN master_category on master_product.master_category_id = master_category.id").Where("approval_id = ?", approvalId).Find(&approval).Error

	return approval, error
}
