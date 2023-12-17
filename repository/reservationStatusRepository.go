package repository

import (
	"fmt"
	"svc-warehouse/model"

	"gorm.io/gorm"
)

type ReservationStatusRepostoryInterface interface {
	CreateReservationStatus(request *model.ReservationStatus) ([]model.ReservationStatus, error)
	ReadReservationStatus() ([]model.ReservationStatus, error)
	UpdateReservationStatus(id int, request *model.ReservationStatus) error
	ReadDetailReservationStatus(id int) ([]model.ReservationStatus, error)
	DeleteReservationStatus(id int) error
	ReadReservationStatusByName(name string, db *gorm.DB) (model.ReservationStatus, error)
}

func (repo *repository) CreateReservationStatus(request *model.ReservationStatus) ([]model.ReservationStatus, error) {
	var reservationStatus []model.ReservationStatus

	error := repo.db.Table("reservation_status").Create(request).Last(&reservationStatus).Error

	return reservationStatus, error
}

func (repo *repository) ReadReservationStatus() ([]model.ReservationStatus, error) {
	var reservationStatus []model.ReservationStatus

	error := repo.db.Table("reservation_status").Where("is_active = ?", 1).Order("name").Find(&reservationStatus).Error

	return reservationStatus, error
}

func (repo *repository) UpdateReservationStatus(id int, request *model.ReservationStatus) error {
	var reservationStatus []model.ReservationStatus

	query := fmt.Sprintf("UPDATE reservation_status SET name = @Name, updated_at = @UpdatedAt WHERE id = %d", id)
	error := repo.db.Raw(query, request).Scan(&reservationStatus).Error

	return error
}

func (repo *repository) ReadDetailReservationStatus(id int) ([]model.ReservationStatus, error) {
	var reservationStatus []model.ReservationStatus

	error := repo.db.Table("reservation_status").Where("id = ?", id).Find(&reservationStatus).Error

	return reservationStatus, error
}

func (repo *repository) DeleteReservationStatus(id int) error {
	var reservationStatus []model.ReservationStatus

	query := fmt.Sprintf("UPDATE reservation_status SET is_active = 0 WHERE id = %d", id)
	error := repo.db.Raw(query).Scan(&reservationStatus).Error

	return error
}

func (repo *repository) ReadReservationStatusByName(name string, db *gorm.DB) (model.ReservationStatus, error) {
	var reservationStatus model.ReservationStatus

	error := db.Table("reservation_status").Where("name = ? AND is_active = ?", name, 1).Find(&reservationStatus).Error

	return reservationStatus, error
}
