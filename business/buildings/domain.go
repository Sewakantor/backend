package buildings

import (
	"capstone/business/complexs"
	"context"
	"time"

	"gorm.io/gorm"
)

type Domain struct {
	Id         		uint `gorm:"primaryKey"`
	Name       		string
	Complex_ID 		uint
	Complex 		complexs.Domain 
	Photo			string
	Year			uint32
	Floors 			uint32
	Floor_Surface	uint32
	Total_Surface	uint32
	CreatedAt		time.Time
	UpdatedAt		time.Time
	DeletedAt		gorm.DeletedAt `gorm:"index"`

}

type BuildingUsecaseInterface interface {
	Add(ctx context.Context, domain Domain) (Domain, error)
	GetAll(ctx context.Context) ([]Domain, error)
	GetByID(ctx context.Context, id uint) (Domain, error)
	GetByComplexID(complexid uint, ctx context.Context) ([]Domain, error)
	Update(ctx context.Context, id uint, domain Domain) (Domain, error)
	Delete(ctx context.Context, id uint) error
}

type BuildingRepoInterface interface {
	Add(ctx context.Context, domain Domain) (Domain, error)
	GetAll(ctx context.Context) ([]Domain, error)
	GetByID(ctx context.Context, id uint) (Domain, error)
	GetByComplexID(complexid uint, ctx context.Context) ([]Domain, error)
	Update(ctx context.Context, id uint, domain Domain) (Domain, error)
	Delete(ctx context.Context, id uint) error
}