package response

import (
	"capstone/business/addresses"
	"time"

	"gorm.io/gorm"
)

type AddressResponse struct {
	Id         uint				`gorm:"primaryKey"`
	Street     string  			`json:"street"`
	City       string  			`json:"city"`
	State      string  			`json:"state"`
	Country    string  			`json:"country"`
	PostalCode string  			`json:"postalcode"`
	Latitude   float64 			`json:"latitude"`
	Longtitude float64 			`json:"longtitude"`
	CreatedAt  time.Time		`json:"createdAt"`
	UpdatedAt  time.Time		`json:"updatedAt"`
	DeletedAt  gorm.DeletedAt	`gorm:"index"`
}

func FromDomainAddress(domain addresses.Domain) AddressResponse {
	return AddressResponse{
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

func FromDomainAddressArray(data []addresses.Domain) []AddressResponse{
	var res []AddressResponse
	for _, val := range data{
		res = append(res, FromDomainAddress(val))
	}
	return res
}

