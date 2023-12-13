package model

import "time"

type Inbound struct {
	ID            int       `json:"id" gorm:"primaryKey"`
	InboundNumber string    `json:"inbound_number"`
	Description   string    `json:"description"`
	CreatedBy     int       `json:"created_by"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}

type InboundRequest struct {
	Inbound
	Product []Product `json:"product,omitempty" binding:"required"`
}
type InboundResponse struct {
	StandardResponse
	Result []InboundRequest `json:"result,omitempty"`
}

type GetInboundReponse struct {
	StandardResponse
	Result []Inbound `json:"result,omitempty"`
}
