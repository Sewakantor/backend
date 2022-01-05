package facilities

import (
	"context"
	"time"

	"gorm.io/gorm"
)

type Domain struct {
	Id        uint `gorm:"primaryKey"`
	Name      string
	Types     string
	Lat       float64
	Long      float64
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

type FacilityUsecaseInterface interface {
	Add(ctx context.Context, domain Domain) (Domain, error)
	GetAll(ctx context.Context) ([]Domain, error)
	Update(ctx context.Context, id uint, domain Domain) (Domain, error)
	Delete(ctx context.Context, id uint) error
}

type FacilityRepoInterface interface {
	Add(ctx context.Context, domain Domain) (Domain, error)
	GetAll(ctx context.Context) ([]Domain, error)
	Update(ctx context.Context, id uint, domain Domain) (Domain, error)
	Delete( ctx context.Context, id uint) error
}