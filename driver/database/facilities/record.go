package facilities

import (
	"capstone/business/facilities"
	"time"

	"gorm.io/gorm"
)

type Facility struct {
	Id   		uint `gorm:"primaryKey"`
	Name 		string
	Types		string
	Lat       	float64
	Long      	float64
	CreatedAt 	time.Time
	UpdatedAt 	time.Time
	DeletedAt 	gorm.DeletedAt `gorm:"index"`
}

func (facility *Facility) ToDomain() facilities.Domain {
	return facilities.Domain{
		Id: facility.Id,
		Name: facility.Name,
		Types: facility.Types,
		Lat: facility.Lat,
		Long: facility.Long,
		CreatedAt: facility.CreatedAt,
		UpdatedAt: facility.UpdatedAt,
	}
}

func FromDomain(domain facilities.Domain) Facility {
	return Facility{
		Id: domain.Id,
		Name: domain.Name,
		Types: domain.Types,
		Lat: domain.Lat,
		Long: domain.Long,
		CreatedAt: domain.CreatedAt,
		UpdatedAt: domain.UpdatedAt,
	}
}

func GetAllFacility(data []Facility) []facilities.Domain{
	res := []facilities.Domain{}
	for _, val := range data{
		res = append(res, val.ToDomain())
	}
	return res
}