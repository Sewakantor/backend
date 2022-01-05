package request

import "capstone/business/addresses"

type AddressRequest struct {
	Street     string  `json:"street"`
	City       string  `json:"city"`
	State      string  `json:"state"`
	Country    string  `json:"country"`
	PostalCode string  `json:"postalcode"`
	Latitude   float64 `json:"latitude"`
	Longtitude float64 `json:"longtitude"`
}

func (address *AddressRequest) ToDomain() *addresses.Domain {
	return &addresses.Domain{
		Street:     address.Street,
		City:       address.City,
		State:      address.State,
		Country:    address.Country,
		PostalCode: address.PostalCode,
		Latitude:   address.Latitude,
		Longtitude: address.Longtitude,
	}
}