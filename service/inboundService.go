package service

import (
	"fmt"
	"svc-warehouse/model"
	"svc-warehouse/repository"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type InboundServiceInterface interface {
	CreateInbound(request *model.InboundRequest, db *gorm.DB) ([]model.InboundRequest, error)
}

type inboundService struct {
	inboundRepository       repository.InboundRepostoryInterface
	productRepository       repository.ProductRepostoryInterface
	inboundDetailRepository repository.InboundDetailRepostoryInterface
}

func InboundService(inboundRepository repository.InboundRepostoryInterface, productRepository repository.ProductRepostoryInterface, inboundDetailRepository repository.InboundDetailRepostoryInterface) *inboundService {
	return &inboundService{inboundRepository, productRepository, inboundDetailRepository}
}

func (service *inboundService) CreateInbound(request *model.InboundRequest, db *gorm.DB) ([]model.InboundRequest, error) {
	var result []model.InboundRequest

	now := time.Now()
	request.CreatedAt = now
	request.UpdatedAt = now
	request.InboundNumber = now.Format("2006-01-02") + "-" + uuid.NewString()

	dbTrx := db.Begin()
	inbound, error := service.inboundRepository.CreateInbound(&request.Inbound, dbTrx)
	if error != nil {
		fmt.Println(error)
		dbTrx.Rollback()
		return result, error
	} else {
		request.Inbound = inbound

		products := []model.Product{}
		for _, v := range request.Product {

			v.IsActive = 1
			v.CreatedAt = now
			v.UpdatedAt = now
			product := model.Product{}
			product, error = service.productRepository.CreateProduct(&v, dbTrx)

			if error != nil {
				fmt.Println(error)
				dbTrx.Rollback()
				return result, error
			} else {

				products = append(products, product)
				inboundDetail := model.InboundDetail{
					InboundId: inbound.ID,
					ProductId: product.ID,
					CreatedAt: now,
					UpdatedAt: now,
				}

				_, error = service.inboundDetailRepository.CreateInboundDetail(&inboundDetail, dbTrx)
				if error != nil {
					fmt.Println(error)
					dbTrx.Rollback()
					return result, error
				}
			}
		}
		request.Product = products
	}

	result = append(result, *request)
	dbTrx.Commit()

	return result, error
}
