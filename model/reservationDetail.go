package model

import "time"

type ReservationDetail struct {
	ID              int       `json:"id" gorm:"primaryKey"`
	ReservationId   int       `json:"reservation_id" binding:"required"`
	MasterProductId int       `json:"master_product_id"`
	Product         string    `json:"product" gorm:"<-:false"`
	Category        string    `json:"category" gorm:"<-:false"`
	Quantity        int       `json:"quantity"`
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"updated_at"`
}

type ReservationDetailResponse struct {
	StandardResponse
	Result []ReservationDetail `json:"result,omitempty"`
}
