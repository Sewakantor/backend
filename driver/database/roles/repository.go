package roles

import (
	"capstone/business/roles"
	"context"
	"errors"

	"gorm.io/gorm"
)

type RoleRepository struct {
	DB *gorm.DB
}

func NewRoleRepository(DB *gorm.DB) roles.RoleRepoInterface {
	return &RoleRepository{
		DB: DB,
	}
}

func (Repo *RoleRepository) Add(ctx context.Context, domain roles.Domain) (roles.Domain, error) {
	role := Role{
		Id 			:domain.Id,
		Name 		:domain.Name,
	}
	err := Repo.DB.Create(&role)
	if err.Error != nil {
		return roles.Domain{}, err.Error
	}
	return role.ToDomain(), nil
}

func (Repo *RoleRepository) GetAll(ctx context.Context) ([]roles.Domain, error){
	var role []Role
	err := Repo.DB.Find(&role)
	if err.Error != nil {
		return []roles.Domain{}, err.Error
	}
	return GetAllRoles(role), nil
}

func (Repo *RoleRepository) Delete(ctx context.Context, id uint) error{
	role := Role{}
	err := Repo.DB.Delete(&role, id)
	if err.Error!= nil {
		return err.Error
		
	}
	if err.RowsAffected == 0 {
		return errors.New("data kosong")
	}
	return nil
}