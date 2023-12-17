package service

import (
	"fmt"
	"svc-warehouse/model"
	"svc-warehouse/repository"
	"time"

	"gorm.io/gorm"
)

type ReservationServiceInterface interface {
	CreateReservation(request *model.ReservationRequest, db *gorm.DB) ([]model.ReservationRequest, error)
	ReadReservation() ([]model.Reservation, error)
	ReadReservationById(id int) ([]model.ReservationRequest, error)
}

type reservationService struct {
	reservationRepository       repository.ReservationRepostoryInterface
	reservationDetailRepository repository.ReservationDetailRepostoryInterface
	reservationStatusRepository repository.ReservationStatusRepostoryInterface
}

func ReservationService(reservationRepository repository.ReservationRepostoryInterface, reservationDetailRepository repository.ReservationDetailRepostoryInterface, reservationStatusRepository repository.ReservationStatusRepostoryInterface) *reservationService {
	return &reservationService{reservationRepository, reservationDetailRepository, reservationStatusRepository}
}

func (service *reservationService) CreateReservation(request *model.ReservationRequest, db *gorm.DB) ([]model.ReservationRequest, error) {
	now := time.Now()
	request.CreatedAt = now
	request.UpdatedAt = now

	dbTrx := db.Begin()
	reservationStatus, error := service.reservationStatusRepository.ReadReservationStatusByName("Pending", dbTrx)
	if error != nil {
		fmt.Println(error)
		dbTrx.Rollback()
		return nil, error
	}

	request.ReservationStatusId = reservationStatus.ID
	reservation, error := service.reservationRepository.CreateReservation(&request.Reservation, dbTrx)
	if error != nil {
		fmt.Println(error)
		dbTrx.Rollback()
		return nil, error
	}

	reservationDetails := []model.ReservationDetail{}
	for _, v := range request.ReservationDetail {

		v.ReservationId = reservation.ID
		v.CreatedAt = now
		v.UpdatedAt = now

		reservationDetail, error := service.reservationDetailRepository.CreateReservationDetail(&v, dbTrx)
		if error != nil {
			fmt.Println(error)
			dbTrx.Rollback()
			return nil, error
		}

		reservationDetails = append(reservationDetails, reservationDetail)
	}

	result := []model.ReservationRequest{}
	result = append(result, model.ReservationRequest{
		Reservation:       reservation,
		ReservationDetail: reservationDetails,
	})
	dbTrx.Commit()
	return result, error
}

func (service *reservationService) ReadReservation() ([]model.Reservation, error) {

	reservation, error := service.reservationRepository.ReadReservation()

	return reservation, error
}

func (service *reservationService) ReadReservationById(id int) ([]model.ReservationRequest, error) {

	reservation, error := service.reservationRepository.ReadReservationById(id)
	if error != nil {
		return nil, error
	}

	reservationDetail, error := service.reservationDetailRepository.ReadReservationDetailByReservationId(id)
	if error != nil {
		return nil, error
	}

	result := []model.ReservationRequest{}
	result = append(result, model.ReservationRequest{
		Reservation:       reservation,
		ReservationDetail: reservationDetail,
	})

	return result, error
}
