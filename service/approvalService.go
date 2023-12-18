package service

import (
	"fmt"
	"svc-warehouse/model"
	"svc-warehouse/repository"
	"time"

	"gorm.io/gorm"
)

type ApprovalServiceInterface interface {
	CreateApproval(request *model.ApprovalRequest, db *gorm.DB) ([]model.ApprovalRequest, error)
	ReadApproval() ([]model.Approval, error)
	ReadApprovalById(id int) ([]model.ApprovalRequest, error)
}

type approvalService struct {
	approvalRepository       repository.ApprovalRepostoryInterface
	approvalDetailRepository repository.ApprovalDetailRepostoryInterface
}

func ApprovalService(approvalRepository repository.ApprovalRepostoryInterface, approvalDetailRepository repository.ApprovalDetailRepostoryInterface) *approvalService {
	return &approvalService{approvalRepository, approvalDetailRepository}
}

func (service *approvalService) CreateApproval(request *model.ApprovalRequest, db *gorm.DB) ([]model.ApprovalRequest, error) {
	now := time.Now()
	request.CreatedAt = now
	request.UpdatedAt = now

	dbTrx := db.Begin()
	approval, error := service.approvalRepository.CreateApproval(&request.Approval, dbTrx)
	if error != nil {
		fmt.Println(error)
		dbTrx.Rollback()
		return nil, error
	}

	approvalDetails := []model.ApprovalDetail{}
	for _, v := range request.ApprovalDetail {

		v.ApprovalId = approval.ID
		v.CreatedAt = now
		v.UpdatedAt = now

		approvalDetail, error := service.approvalDetailRepository.CreateApprovalDetail(&v, dbTrx)
		if error != nil {
			fmt.Println(error)
			dbTrx.Rollback()
			return nil, error
		}

		approvalDetails = append(approvalDetails, approvalDetail)
	}

	result := []model.ApprovalRequest{}
	result = append(result, model.ApprovalRequest{
		Approval:       approval,
		ApprovalDetail: approvalDetails,
	})
	dbTrx.Commit()
	return result, error
}

func (service *approvalService) ReadApproval() ([]model.Approval, error) {

	approval, error := service.approvalRepository.ReadApproval()

	return approval, error
}

func (service *approvalService) ReadApprovalById(id int) ([]model.ApprovalRequest, error) {

	approval, error := service.approvalRepository.ReadApprovalById(id)
	if error != nil {
		return nil, error
	}

	approvalDetail, error := service.approvalDetailRepository.ReadApprovalDetailByApprovalId(id)
	if error != nil {
		return nil, error
	}

	result := []model.ApprovalRequest{}
	result = append(result, model.ApprovalRequest{
		Approval:       approval,
		ApprovalDetail: approvalDetail,
	})

	return result, error
}
