package model

import "time"

type Reservation struct {
	ID                  int       `json:"id" gorm:"primaryKey"`
	ReservationNumber   string    `json:"reservation_number" gorm:"unique" binding:"required"`
	ReservationStatusId int       `json:"reservation_status_id"`
	Description         string    `json:"description"`
	CreatedBy           int       `json:"created_by" binding:"required"`
	CreatedAt           time.Time `json:"created_at"`
	UpdatedAt           time.Time `json:"updated_at"`
}

type ReservationRequest struct {
	Reservation
	ReservationDetail []ReservationDetail `json:"reservation_detail"`
}
type ReservationResponse struct {
	StandardResponse
	Result []ReservationRequest `json:"result,omitempty"`
}

type GetReservationReponse struct {
	StandardResponse
	Result []Reservation `json:"result,omitempty"`
}
