package model

import "time"

type Product struct {
	ID              int       `json:"id" gorm:"primaryKey"`
	MasterProductId int       `json:"master_product_id"`
	SerialNumber    string    `json:"serial_number"`
	ProductStatusId int       `json:"product_status_id"`
	IsActive        int       `json:"is_active"`
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"updated_at"`
}

type ProductResponse struct {
	StandardResponse
	Result []Product `json:"result,omitempty"`
}
