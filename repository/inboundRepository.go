package repository

import (
	"svc-warehouse/model"

	"gorm.io/gorm"
)

type InboundRepostoryInterface interface {
	CreateInbound(request *model.Inbound, db *gorm.DB) (model.Inbound, error)
	ReadInbound() ([]model.Inbound, error)
}

func (repo *repository) CreateInbound(request *model.Inbound, db *gorm.DB) (model.Inbound, error) {
	var inbound model.Inbound

	error := db.Table("inbound").Create(request).Last(&inbound).Error

	return inbound, error
}

func (repo *repository) ReadInbound() ([]model.Inbound, error) {
	var inbound []model.Inbound

	error := repo.db.Table("inbound").Order("created_at desc").Find(&inbound).Error

	return inbound, error
}
