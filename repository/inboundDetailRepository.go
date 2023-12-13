package repository

import (
	"svc-warehouse/model"

	"gorm.io/gorm"
)

type InboundDetailRepostoryInterface interface {
	CreateInboundDetail(request *model.InboundDetail, db *gorm.DB) ([]model.InboundDetail, error)
}

func (repo *repository) CreateInboundDetail(request *model.InboundDetail, db *gorm.DB) ([]model.InboundDetail, error) {
	var inbound []model.InboundDetail

	error := db.Table("inbound_detail").Create(request).Last(&inbound).Error

	return inbound, error
}
