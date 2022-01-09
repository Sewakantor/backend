package complexs

import (
	"context"
	"time"

	"gorm.io/gorm"
)

type Domain struct {
	Id        uint `gorm:"primaryKey"`
	Name      string
	Street     string
	City       string
	State      string
	Country    string
	PostalCode string
	Latitude   float64
	Longtitude float64
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

type ComplexUsecaseInterface interface {
	Add(ctx context.Context, domain Domain) (Domain, error)
	GetAll(ctx context.Context) ([]Domain, error)
	GetByID(ctx context.Context, id uint) (Domain, error)
	Update(ctx context.Context, id uint, domain Domain) (Domain, error)
	Delete(ctx context.Context, id uint) error
}

type ComplexRepoInterface interface {
	Add(ctx context.Context, domain Domain) (Domain, error)
	GetAll(ctx context.Context) ([]Domain, error)
	GetByID(ctx context.Context, id uint) (Domain, error)
	Update(ctx context.Context, id uint, domain Domain) (Domain, error)
	Delete(ctx context.Context, id uint) error
}