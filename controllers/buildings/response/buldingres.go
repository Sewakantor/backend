package response

import (
	"capstone/business/buildings"
	"capstone/controllers/complexs/response"
	"time"

	"gorm.io/gorm"
)

type BuildingResponse struct {
	Id            uint   					`gorm:"primaryKey" json:"id"`
	Name          string 					`json:"name"`
	Complex_ID    uint   					`json:"complex_id"`
	Complex       response.ComplexResponse	`json:"complex"`
	Photo         string					`json:"photo"`
	Year          uint32					`json:"year"`
	Floors        uint32					`json:"floors"`
	Floor_Surface uint32					`json:"floor_surface"`
	Total_Surface uint32					`json:"total_surface"`
	CreatedAt     time.Time					
	UpdatedAt     time.Time
	DeletedAt     gorm.DeletedAt 			`gorm:"index"`
}

func FromDomainBuilding(domain buildings.Domain) BuildingResponse {
	return BuildingResponse{
		Id:         	domain.Id,
		Name:       	domain.Name,
		Complex_ID: 	domain.Complex_ID,
		Complex:		response.FromDomainComplex(domain.Complex),
		Photo: 			domain.Photo,
		Year: 			domain.Year,
		Floors: 		domain.Floors,
		Floor_Surface: 	domain.Floor_Surface,
		Total_Surface: 	domain.Total_Surface,	
		CreatedAt:  	domain.CreatedAt,
		UpdatedAt:  	domain.UpdatedAt,
	}
}

func FromDomainBuildingAll(data []buildings.Domain) []BuildingResponse{
	var res []BuildingResponse
	for _, val := range data{
		res = append(res, FromDomainBuilding(val))
	}
	return res
}