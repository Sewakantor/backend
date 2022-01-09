package response

import (
	"capstone/business/complexs"
	"time"

	"gorm.io/gorm"
)

type ComplexResponse struct {
	Id        uint   `json:"id"`
	Name      string `json:"name"`
	Street     string `json:"street"`
	City       string `json:"city"`
	State      string `json:"state"`
	Country    string `json:"country"`
	PostalCode string `json:"postalCode"`
	Latitude   float64 `json:"latitude"`
	Longtitude float64 `json:"longtitude"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	DeletedAt gorm.DeletedAt `json:"deletedAt"`
}

func FromDomainComplex(domain complexs.Domain) ComplexResponse {
	return ComplexResponse{
		Id:        domain.Id,
		Name:      domain.Name,
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

func GetAll(domain []complexs.Domain) []ComplexResponse {
	var GetAll []ComplexResponse
	for _, val := range domain {
		GetAll = append(GetAll, FromDomainComplex(val))
	}
	return GetAll
}
