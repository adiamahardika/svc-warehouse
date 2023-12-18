package model

import "time"

type ApprovalDetail struct {
	ID            int       `json:"id" gorm:"primaryKey"`
	ApprovalId    int       `json:"approval_id"`
	ProductId     int       `json:"product_id"`
	Product       string    `json:"product" gorm:"<-:false"`
	SerialNumber  string    `json:"serial_number" gorm:"<-:false"`
	ProductStatus string    `json:"product_status" gorm:"<-:false"`
	Category      string    `json:"category" gorm:"<-:false"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}
