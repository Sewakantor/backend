package response

import (
	"capstone/business/units"
	"capstone/controllers/buildings/response"
	"time"

	"gorm.io/gorm"
)

type UnitResponse struct {
	Id          uint   						`gorm:"primaryKey" json:"id"`
	Name        string 						`json:"name"`
	Category    string 						`json:"category"`
	Building_ID uint   						`json:"building_id"`
	Building    response.BuildingResponse	`json:"building"`
	Surface     uint						`json:"surface"`
	Capacity    uint						`json:"capacity"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   gorm.DeletedAt 				`gorm:"index"`
}

func FromDomainUnit(domain units.Domain) UnitResponse {
	return UnitResponse{
		Id:         	domain.Id,
		Name:       	domain.Name,
		Category: 		domain.Category,
		Building_ID: 	domain.Building_ID,
		Building:		response.FromDomainBuilding(domain.Building),
		Surface: 		domain.Surface,
		Capacity: 		domain.Capacity,	
		CreatedAt:  	domain.CreatedAt,
		UpdatedAt:  	domain.UpdatedAt,
	}
}

func FromDomainUnitAll(data []units.Domain) []UnitResponse{
	var res []UnitResponse
	for _, val := range data{
		res = append(res, FromDomainUnit(val))
	}
	return res
}