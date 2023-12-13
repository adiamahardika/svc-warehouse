package model

import "time"

type InboundDetail struct {
	ID        int       `json:"id" gorm:"primaryKey"`
	InboundId int       `json:"inbound_id"`
	ProductId int       `json:"product_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type InboundDetailResponse struct {
	StandardResponse
	Result []InboundDetail `json:"result,omitempty"`
}
