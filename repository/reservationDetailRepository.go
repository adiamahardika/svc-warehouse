package repository

import (
	"svc-warehouse/model"

	"gorm.io/gorm"
)

type ReservationDetailRepostoryInterface interface {
	CreateReservationDetail(request *model.ReservationDetail, db *gorm.DB) (model.ReservationDetail, error)
}

func (repo *repository) CreateReservationDetail(request *model.ReservationDetail, db *gorm.DB) (model.ReservationDetail, error) {
	var reservation model.ReservationDetail

	error := db.Table("reservation_detail").Create(request).Last(&reservation).Error

	return reservation, error
}
