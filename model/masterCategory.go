package model

import "time"

type MasterCategory struct {
	ID        int       `json:"id,omitempty" gorm:"primaryKey"`
	Name      string    `json:"name,omitempty" gorm:"unique" binding:"required"`
	IsActive  int       `json:"is_active,omitempty"`
	CreatedAt time.Time `json:"created_at,omitempty"`
	UpdatedAt time.Time `json:"updated_at,omitempty"`
}

type MasterCategoryResponse struct {
	StandardResponse
	Result []MasterCategory `json:"result"`
}
