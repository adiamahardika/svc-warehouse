package repository

import (
	"svc-warehouse/model"

	"gorm.io/gorm"
)

type ReservationDetailRepostoryInterface interface {
	CreateReservationDetail(request *model.ReservationDetail, db *gorm.DB) (model.ReservationDetail, error)
	ReadReservationDetailByReservationId(reservationId int) ([]model.ReservationDetail, error)
}

func (repo *repository) CreateReservationDetail(request *model.ReservationDetail, db *gorm.DB) (model.ReservationDetail, error) {
	var reservation model.ReservationDetail

	error := db.Table("reservation_detail").Create(request).Last(&reservation).Error

	return reservation, error
}

func (repo *repository) ReadReservationDetailByReservationId(reservationId int) ([]model.ReservationDetail, error) {
	var reservation []model.ReservationDetail

	error := repo.db.Table("reservation_detail").Select("reservation_detail.*, master_product.name as product, master_category.name as category").Joins("LEFT JOIN master_product on reservation_detail.master_product_id = master_product.id").Joins("LEFT JOIN master_category on master_product.master_category_id = master_category.id").Where("reservation_id = ?", reservationId).Find(&reservation).Error

	return reservation, error
}
