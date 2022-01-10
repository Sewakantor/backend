package buildings

import (
	"capstone/business/buildings"
	"capstone/driver/database/complexs"
	"time"

	"gorm.io/gorm"
)

type Building struct {
	Id         		uint `gorm:"primaryKey"`
	Name       		string `gorm:"unique"`
	Complex_ID 		uint
	Complex    		complexs.Complex `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Photo			string
	Year			uint32
	Floors 			uint32
	Floor_Surface	uint32
	Total_Surface	uint32
	CreatedAt  		time.Time
	UpdatedAt  		time.Time
	DeletedAt  		gorm.DeletedAt `gorm:"index"`
}

func (building *Building) ToDomain() buildings.Domain {
	return buildings.Domain{
		Id:         	building.Id,
		Name:       	building.Name,
		Complex_ID: 	building.Complex_ID,
		Complex:		building.Complex.ToDomain(),
		Photo: 			building.Photo,
		Year: 			building.Year,
		Floors: 		building.Floors,
		Floor_Surface: 	building.Floor_Surface,
		Total_Surface: 	building.Total_Surface,	
		CreatedAt:  	building.CreatedAt,
		UpdatedAt:  	building.UpdatedAt,
	}
}

func FromDomain(domain buildings.Domain) Building {
	return Building{
		Id:         	domain.Id,
		Name:       	domain.Name,
		Complex_ID: 	domain.Complex_ID,
		Complex:		complexs.FromDomain(domain.Complex),
		Photo: 			domain.Photo,
		Year: 			domain.Year,
		Floors: 		domain.Floors,
		Floor_Surface: 	domain.Floor_Surface,
		Total_Surface: 	domain.Total_Surface,	
		CreatedAt:  	domain.CreatedAt,
		UpdatedAt:  	domain.UpdatedAt,
	}
}

func GetAll(data []Building) []buildings.Domain {
	res := []buildings.Domain{}
	for _, val := range data {
		res = append(res, val.ToDomain())
	}
	return res

}