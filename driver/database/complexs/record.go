package complexs

import (
	"capstone/business/complexs"
	"time"

	"gorm.io/gorm"
)

type Complex struct {
	Id        uint   `gorm:"primaryKey"`
	Name      string `gorm:"unique"`
	Address   string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

func (complex *Complex) ToDomain() complexs.Domain {
	return complexs.Domain{
		Id: complex.Id,
		Name: complex.Name,
		Address: complex.Address,
		CreatedAt: complex.CreatedAt,
		UpdatedAt: complex.UpdatedAt,
	}
}

func FromDomain (domain complexs.Domain) Complex{
	return Complex{
		Id: domain.Id,
		Name: domain.Name,
		Address: domain.Address,
		CreatedAt: domain.CreatedAt,
		UpdatedAt: domain.UpdatedAt,
	}
}

func GetAll(data []Complex) []complexs.Domain {
	res := []complexs.Domain {}
	for _, val := range data {
		res = append(res, val.ToDomain())
	}
	return res
	
}