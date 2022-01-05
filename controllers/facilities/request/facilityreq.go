package request

import "capstone/business/facilities"

type FacilityRequest struct {
	Name  string  `json:"name"`
	Types string  `json:"types"`
	Lat   float64 `json:"lat"`
	Long  float64 `json:"long"`
}

func (facility *FacilityRequest) ToDomain() *facilities.Domain {
	return &facilities.Domain{
		Name: facility.Name,
		Types: facility.Types,
		Lat: facility.Lat,
		Long: facility.Long,
	}
}