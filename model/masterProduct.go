package model

import "time"

type MasterProduct struct {
	ID               int       `json:"id" gorm:"primaryKey"`
	Name             string    `json:"name" gorm:"unique" binding:"required"`
	MasterCategoryId int       `json:"master_category_id" binding:"required"`
	Category         string    `json:"category,omitempty" gorm:"<-:false"`
	IsActive         int       `json:"is_active"`
	CreatedAt        time.Time `json:"created_at"`
	UpdatedAt        time.Time `json:"updated_at"`
}

type MasterProductResponse struct {
	StandardResponse
	Result []MasterProduct `json:"result,omitempty"`
}
