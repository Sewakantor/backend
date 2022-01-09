package request

import "capstone/business/complexs"

type ComplexRequest struct {
	Name       string `json:"name"`
	Street     string `json:"street"`
	City       string `json:"city"`
	State      string `json:"state"`
	Country    string `json:"country"`
	PostalCode string `json:"postalCode"`
	Latitude   float64 `json:"latitude"`
	Longtitude float64 `json:"longtitude"`
}

func (complex *ComplexRequest) ToDomain() *complexs.Domain {
	return &complexs.Domain{
		Name: complex.Name,
		Street: complex.Street,
		City: complex.City,
		State: complex.State,
		Country: complex.Country,
		PostalCode: complex.PostalCode,
		Latitude: complex.Latitude,
	}
}