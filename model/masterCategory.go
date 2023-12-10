package model

import "time"

type MasterCategory struct {
	ID        int       `json:"id" gorm:"primaryKey"`
	Name      string    `json:"name" gorm:"unique" binding:"required"`
	IsActive  int       `json:"is_active"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type MasterCategoryResponse struct {
	StandardResponse
	Result []MasterCategory `json:"result,omitempty"`
}
