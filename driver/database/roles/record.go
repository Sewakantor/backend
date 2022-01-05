package roles

import (
	"capstone/business/roles"
	"time"

	"gorm.io/gorm"
)

type Role struct {
	Id        uint `gorm:"primaryKey"`
	Name      string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

func (role *Role) ToDomain() roles.Domain {
	return roles.Domain{
		Id: role.Id,
		Name: role.Name,
		CreatedAt: role.CreatedAt,
		UpdatedAt: role.UpdatedAt,
	}
}

func FromDomain(domain roles.Domain) Role {
	return Role{
		Id: domain.Id,
		Name: domain.Name,
		CreatedAt: domain.CreatedAt,
		UpdatedAt: domain.UpdatedAt,
	}
}

func GetAllRoles(data []Role) []roles.Domain{
	res := []roles.Domain{}
	for _, val := range data {
		res = append(res, val.ToDomain())
	}
	return res
}