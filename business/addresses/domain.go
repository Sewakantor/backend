package addresses

import (
	"context"
	"time"

	"gorm.io/gorm"
)

type Domain struct {
	Id         uint
	Street     string
	City       string
	State      string
	Country    string
	PostalCode string
	Latitude   float64
	Longtitude float64
	CreatedAt  time.Time
	UpdatedAt  time.Time
	DeletedAt  gorm.DeletedAt
}

type AddressUsecaseInterface interface {
	Add(ctx context.Context, domain Domain) (Domain, error)
	GetAll(ctx context.Context) ([]Domain, error)
	Update(ctx context.Context, id uint, domain Domain) (Domain, error)
	Delete(ctx context.Context, id uint )error
}

type AddressRepoInterface interface{
	Add(ctx context.Context, domain Domain) (Domain, error)
	GetAll(ctx context.Context) ([]Domain, error)
	Update(ctx context.Context, id uint, domain Domain) (Domain, error)
	Delete(ctx context.Context, id uint )error
}