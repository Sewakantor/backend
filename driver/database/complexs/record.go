package complexs

import (
	"capstone/business/complexs"
	"time"

	"gorm.io/gorm"
)

type Complex struct {
	Id        uint   `gorm:"primaryKey"`
	Name      string `gorm:"unique"`
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

func (complex *Complex) ToDomain() complexs.Domain {
	return complexs.Domain{
		Id: complex.Id,
		Name: complex.Name,
		Street: complex.Street,
		City: complex.City,
		State: complex.State,
		Country: complex.Country,
		PostalCode: complex.PostalCode,
		Latitude: complex.Latitude,
		Longtitude: complex.Longtitude,
		CreatedAt: complex.CreatedAt,
		UpdatedAt: complex.UpdatedAt,
	}
}

func FromDomain (domain complexs.Domain) Complex{
	return Complex{
		Id: domain.Id,
		Name: domain.Name,
		Street: domain.Street,
		City: domain.City,
		State: domain.State,
		Country: domain.Country,
		PostalCode: domain.PostalCode,
		Latitude: domain.Latitude,
		Longtitude: domain.Longtitude,
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