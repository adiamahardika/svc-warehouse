package repository

import (
	"svc-warehouse/model"

	"gorm.io/gorm"
)

type ReservationRepostoryInterface interface {
	CreateReservation(request *model.Reservation, db *gorm.DB) (model.Reservation, error)
	ReadReservation() ([]model.Reservation, error)
}

func (repo *repository) CreateReservation(request *model.Reservation, db *gorm.DB) (model.Reservation, error) {
	var reservation model.Reservation

	error := db.Table("reservation").Create(request).Last(&reservation).Error

	return reservation, error
}

func (repo *repository) ReadReservation() ([]model.Reservation, error) {
	var reservation []model.Reservation

	error := repo.db.Table("reservation").Order("created_at desc").Find(&reservation).Error

	return reservation, error
}
