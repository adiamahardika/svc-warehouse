package service

import (
	"svc-warehouse/model"
	"svc-warehouse/repository"
	"time"
)

type ReservationStatusServiceInterface interface {
	CreateReservationStatus(request *model.ReservationStatus) ([]model.ReservationStatus, error)
	ReadReservationStatus() ([]model.ReservationStatus, error)
	UpdateReservationStatus(id int, request *model.ReservationStatus) ([]model.ReservationStatus, error)
	DeleteReservationStatus(id int) error
}

type reservationStatusService struct {
	repository repository.ReservationStatusRepostoryInterface
}

func ReservationStatusService(repository repository.ReservationStatusRepostoryInterface) *reservationStatusService {
	return &reservationStatusService{repository}
}

func (service *reservationStatusService) CreateReservationStatus(request *model.ReservationStatus) ([]model.ReservationStatus, error) {
	now := time.Now()
	request.CreatedAt = now
	request.UpdatedAt = now
	request.IsActive = 1

	reservationStatus, error := service.repository.CreateReservationStatus(request)

	return reservationStatus, error
}

func (service *reservationStatusService) ReadReservationStatus() ([]model.ReservationStatus, error) {

	reservationStatus, error := service.repository.ReadReservationStatus()

	return reservationStatus, error
}

func (service *reservationStatusService) UpdateReservationStatus(id int, request *model.ReservationStatus) ([]model.ReservationStatus, error) {
	now := time.Now()
	request.UpdatedAt = now
	reservationStatus := []model.ReservationStatus{}

	error := service.repository.UpdateReservationStatus(id, request)
	if error == nil {

		reservationStatus, error = service.repository.ReadDetailReservationStatus(id)
		if error == nil {
			return reservationStatus, error
		}
	}

	return reservationStatus, error
}

func (service *reservationStatusService) DeleteReservationStatus(id int) error {

	error := service.repository.DeleteReservationStatus(id)

	return error
}
