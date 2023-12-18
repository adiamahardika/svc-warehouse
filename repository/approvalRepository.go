package repository

import (
	"svc-warehouse/model"

	"gorm.io/gorm"
)

type ApprovalRepostoryInterface interface {
	CreateApproval(request *model.Approval, db *gorm.DB) (model.Approval, error)
	ReadApproval() ([]model.Approval, error)
	ReadApprovalById(id int) (model.Approval, error)
}

func (repo *repository) CreateApproval(request *model.Approval, db *gorm.DB) (model.Approval, error) {
	var approval model.Approval

	error := db.Table("approval").Create(request).Last(&approval).Error

	return approval, error
}

func (repo *repository) ReadApproval() ([]model.Approval, error) {
	var approval []model.Approval

	error := repo.db.Table("approval").Select("approval.*, reservation.reservation_number, reservation_status.name as reservation_status").Joins("LEFT JOIN reservation on approval.reservation_id = reservation.id").Joins("LEFT JOIN reservation_status on approval.reservation_status_id = reservation_status.id").Order("created_at desc").Find(&approval).Error

	return approval, error
}

func (repo *repository) ReadApprovalById(id int) (model.Approval, error) {
	var approval model.Approval

	error := repo.db.Table("approval").Select("approval.*, reservation.reservation_number, reservation_status.name as reservation_status").Joins("LEFT JOIN reservation on approval.reservation_id = reservation.id").Joins("LEFT JOIN reservation_status on approval.reservation_status_id = reservation_status.id").Where("approval.id = ?", id).Find(&approval).Error

	return approval, error
}
