package model

import "time"

type Approval struct {
	ID                  int       `json:"id" gorm:"primaryKey"`
	ReservationId       int       `json:"reservation_id" binding:"required"`
	ReservationNumber   string    `json:"reservation_number" gorm:"<-:false"`
	ReservationStatusId int       `json:"reservation_status_id" binding:"required"`
	ReservationStatus   string    `json:"reservation_status" gorm:"<-:false"`
	CreatedBy           int       `json:"created_by"`
	CreatedAt           time.Time `json:"created_at"`
	UpdatedAt           time.Time `json:"updated_at"`
}

type ApprovalRequest struct {
	Approval
	ApprovalDetail []ApprovalDetail `json:"approval_detail,omitempty" binding:"required"`
}
type ApprovalResponse struct {
	StandardResponse
	Result []ApprovalRequest `json:"result,omitempty"`
}

type GetApprovalReponse struct {
	StandardResponse
	Result []Approval `json:"result,omitempty"`
}
