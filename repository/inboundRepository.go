package repository

import (
	"svc-warehouse/model"

	"gorm.io/gorm"
)

type InboundRepostoryInterface interface {
	CreateInbound(request *model.Inbound, db *gorm.DB) (model.Inbound, error)
}

func (repo *repository) CreateInbound(request *model.Inbound, db *gorm.DB) (model.Inbound, error) {
	var inbound model.Inbound

	error := db.Table("inbound").Create(request).Last(&inbound).Error

	return inbound, error
}
