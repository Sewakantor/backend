package response

import (
	"capstone/business/facilities"
	"time"

	"gorm.io/gorm"
)

type FacilityResponse struct {
	Id        uint 			`gorm:"primaryKey"`
	Name      string		`json:"name"`
	Types     string		`json:"types"`
	Lat       float64		`json:"lat"`
	Long      float64		`json:"long"`
	CreatedAt time.Time		`json:"createdAt"`
	UpdatedAt time.Time		`json:"updatedAt"`
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

func FromDomainFacility(domain facilities.Domain) FacilityResponse {
	return FacilityResponse{
		Id: domain.Id,
		Name: domain.Name,
		Types: domain.Types,
		Lat: domain.Lat,
		Long: domain.Long,
		CreatedAt: domain.CreatedAt,
		UpdatedAt: domain.UpdatedAt,
	}
}

func FromDomainFacilityArray(data []facilities.Domain) []FacilityResponse{
	var res []FacilityResponse
	for _, val := range data{
		res = append(res, FromDomainFacility(val))
	}
	return res
}
