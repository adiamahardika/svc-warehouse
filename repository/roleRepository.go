package repository

import (
	"fmt"
	"svc-warehouse/model"
)

type RoleRepostoryInterface interface {
	CreateRole(request *model.Role) ([]model.Role, error)
	ReadRole() ([]model.Role, error)
	UpdateRole(id int, request *model.Role) error
	ReadDetailRole(id int) ([]model.Role, error)
	DeleteRole(id int) error
}

func (repo *repository) CreateRole(request *model.Role) ([]model.Role, error) {
	var role []model.Role

	error := repo.db.Table("role").Create(request).Last(&role).Error

	return role, error
}

func (repo *repository) ReadRole() ([]model.Role, error) {
	var role []model.Role

	error := repo.db.Table("role").Where("is_active = ?", 1).Order("name").Find(&role).Error

	return role, error
}

func (repo *repository) UpdateRole(id int, request *model.Role) error {
	var role []model.Role

	query := fmt.Sprintf("UPDATE role SET name = @Name, updated_at = @UpdatedAt WHERE id = %d", id)
	error := repo.db.Raw(query, request).Scan(&role).Error

	return error
}

func (repo *repository) ReadDetailRole(id int) ([]model.Role, error) {
	var role []model.Role

	error := repo.db.Table("role").Where("id = ?", id).Find(&role).Error

	return role, error
}

func (repo *repository) DeleteRole(id int) error {
	var role []model.Role

	query := fmt.Sprintf("UPDATE role SET is_active = 0 WHERE id = %d", id)
	error := repo.db.Raw(query).Scan(&role).Error

	return error
}
