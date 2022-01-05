package addresses

import (
	"capstone/business/addresses"
	"time"

	"gorm.io/gorm"
)

type Address struct {
	Id         uint			`gorm:"primaryKey"`
	Street     string
	City       string
	State      string
	Country    string
	PostalCode string
	Latitude   float64
	Longtitude float64
	CreatedAt  time.Time
	UpdatedAt  time.Time
	DeletedAt  gorm.DeletedAt	`gorm:"index"`
}

func (address *Address) ToDomain() addresses.Domain {
	return addresses.Domain{
		Id: address.Id,
		Street: address.Street,
		City: address.City,
		State: address.State,
		Country: address.Country,
		PostalCode: address.PostalCode,
		Latitude: address.Latitude,
		Longtitude: address.Longtitude,
		CreatedAt: address.CreatedAt,
		UpdatedAt: address.UpdatedAt,
	}
}

func FromDomain(domain addresses.Domain) Address {
	return Address{
		Id: domain.Id,
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

func GetAllAddress(data []Address) []addresses.Domain{
	res := []addresses.Domain{}
	for _, val := range data{
		res = append(res, val.ToDomain())
	}
	return res
}