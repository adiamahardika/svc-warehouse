package repository

import (
	"fmt"
	"svc-warehouse/model"

	"gorm.io/gorm"
)

type ReservationRepostoryInterface interface {
	CreateReservation(request *model.Reservation, db *gorm.DB) (model.Reservation, error)
	ReadReservation() ([]model.Reservation, error)
	ReadReservationById(id int) (model.Reservation, error)
	UpdateReservation(id int, request *model.Reservation, db *gorm.DB) error
}

func (repo *repository) CreateReservation(request *model.Reservation, db *gorm.DB) (model.Reservation, error) {
	var reservation model.Reservation

	error := db.Table("reservation").Create(request).Last(&reservation).Error

	return reservation, error
}

func (repo *repository) ReadReservation() ([]model.Reservation, error) {
	var reservation []model.Reservation

	error := repo.db.Table("reservation").Select("reservation.*, reservation_status.name as reservation_status").Joins("LEFT JOIN reservation_status on reservation.reservation_status_id = reservation_status.id").Order("created_at desc").Find(&reservation).Error

	return reservation, error
}

func (repo *repository) ReadReservationById(id int) (model.Reservation, error) {
	var reservation model.Reservation

	error := repo.db.Table("reservation").Where("id = ?", id).Find(&reservation).Error

	return reservation, error
}

func (repo *repository) UpdateReservation(id int, request *model.Reservation, db *gorm.DB) error {
	var reservation model.Reservation

	query := fmt.Sprintf("UPDATE reservation SET reservation_status_id = @ReservationStatusId, updated_at = @UpdatedAt WHERE id = %d", id)
	error := db.Raw(query, request).Scan(&reservation).Error

	return error
}
