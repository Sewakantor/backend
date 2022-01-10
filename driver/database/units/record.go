package units

import (
	"capstone/business/units"
	"capstone/driver/database/buildings"
	"time"

	"gorm.io/gorm"
)

type Unit struct {
	Id          uint `gorm:"primaryKey"`
	Name        string
	Category    string
	Building_ID uint
	Building    buildings.Building `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Surface     uint
	Capacity    uint
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   gorm.DeletedAt `gorm:"index"`
}

func (unit *Unit) ToDomain() units.Domain {
	return units.Domain{
		Id: 			unit.Id,
		Name: 			unit.Name,
		Category: 		unit.Category,
		Building_ID: 	unit.Building_ID,
		Building: 		unit.Building.ToDomain(),
		Surface: 		unit.Surface,
		Capacity: 		unit.Capacity,
		CreatedAt:  	unit.CreatedAt,
		UpdatedAt:  	unit.UpdatedAt,
	}
}

func FromDomain(domain units.Domain) Unit {
	return Unit{
		Id: 			domain.Id,
		Name: 			domain.Name,
		Category: 		domain.Category,
		Building_ID: 	domain.Building_ID,
		Building: 		buildings.FromDomain(domain.Building),
		Surface: 		domain.Surface,
		Capacity: 		domain.Capacity,
		CreatedAt:  	domain.CreatedAt,
		UpdatedAt:  	domain.UpdatedAt,
	}
}

func GetAll(data []Unit) []units.Domain {
	res := []units.Domain{}
	for _, val := range data {
		res = append(res, val.ToDomain())
	}
	return res

}