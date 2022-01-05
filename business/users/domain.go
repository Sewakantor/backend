package users

import (
	"capstone/business/roles"
	"context"
	"time"

	"gorm.io/gorm"
)

type Domain struct {
	Id        uint `gorm:"primaryKey"`
	// User_ID   uint
	Name      string
	Email     string
	Password  string
	Photo     string
	Company   string
	Phone     string
	Status    int
	Roles_ID  uint
	Roles	  roles.Domain
	Token	  string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

type UserUsecaseInterface interface {
	Register(ctx context.Context, domain Domain) (Domain, error)
	Login (ctx context.Context, email string, password string, ) (Domain, string, error)
}

type UserRepoInterface interface {
	Register(ctx context.Context, domain *Domain) (Domain, error)
	GetEmail(ctx context.Context, email string, password string) (Domain, error)
}
