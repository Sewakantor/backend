package users

import (
	"capstone/business/users"
	"capstone/driver/database/roles"
	"time"

	"gorm.io/gorm"
)

type User struct {
	Id 		  uint 		`gorm:"primaryKey"`
	// User_ID   uint
	Name      string
	Email     string 	`gorm:"unique"`
	Password  string
	Photo     string
	Company   string
	Phone     string
	Status    int
	Roles_ID  uint
	Roles     roles.Role `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Token     string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

func (user *User) ToDomain() users.Domain {
	return users.Domain{
		Id: user.Id,
		Name: user.Name,
		Email: user.Email,
		Password: user.Password,
		Photo: user.Photo,
		Company: user.Company,
		Phone: user.Phone,
		Status: user.Status,
		Roles_ID: user.Roles_ID,
		Roles: user.Roles.ToDomain(),
		Token: user.Token,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}
}

func FromDomain(domain users.Domain) User {
	return User{
		Id: domain.Id,
		Name: domain.Name,
		Email: domain.Email,
		Password: domain.Password,
		Photo: domain.Photo,
		Company: domain.Company,
		Phone: domain.Phone,
		Status: domain.Status,
		Roles_ID: domain.Roles_ID,
		Roles: roles.FromDomain(domain.Roles),
		Token: domain.Token,
		CreatedAt: domain.CreatedAt,
		UpdatedAt: domain.UpdatedAt,
	}
}

func GetAllUsers(data []User) []users.Domain{
	res := []users.Domain{}
	for _, val := range data{
		res = append(res, val.ToDomain())
	}
	return res
}