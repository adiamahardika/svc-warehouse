package service

import (
	"svc-warehouse/model"
	"svc-warehouse/repository"
	"time"
)

type RoleServiceInterface interface {
	CreateRole(request *model.Role) ([]model.Role, error)
	ReadRole() ([]model.Role, error)
	UpdateRole(id int, request *model.Role) ([]model.Role, error)
	DeleteRole(id int) error
}

type roleService struct {
	repository repository.RoleRepostoryInterface
}

func RoleService(repository repository.RoleRepostoryInterface) *roleService {
	return &roleService{repository}
}

func (service *roleService) CreateRole(request *model.Role) ([]model.Role, error) {
	now := time.Now()
	request.CreatedAt = now
	request.UpdatedAt = now
	request.IsActive = 1

	role, error := service.repository.CreateRole(request)

	return role, error
}

func (service *roleService) ReadRole() ([]model.Role, error) {

	role, error := service.repository.ReadRole()

	return role, error
}

func (service *roleService) UpdateRole(id int, request *model.Role) ([]model.Role, error) {
	now := time.Now()
	request.UpdatedAt = now
	role := []model.Role{}

	error := service.repository.UpdateRole(id, request)
	if error == nil {

		role, error = service.repository.ReadDetailRole(id)
		if error == nil {
			return role, error
		}
	}

	return role, error
}

func (service *roleService) DeleteRole(id int) error {

	error := service.repository.DeleteRole(id)

	return error
}
