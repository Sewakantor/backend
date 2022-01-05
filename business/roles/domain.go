package roles

import (
	"context"
	"time"

	"gorm.io/gorm"
)

type Domain struct {
	Id        uint `gorm:"primaryKey"`
	Name      string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"` 
}

type RoleUsecaseInterface interface {
	Add(ctx context.Context, domain Domain) (Domain, error)
	GetAll(ctx context.Context) ([]Domain, error)
	Delete(ctx context.Context, id uint) error
}

type RoleRepoInterface interface {
	Add(ctx context.Context, domain Domain) (Domain, error)
	GetAll(ctx context.Context) ([]Domain, error)
	Delete( ctx context.Context, id uint) error
}