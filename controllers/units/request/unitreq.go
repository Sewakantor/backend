package request

import "capstone/business/units"

type UnitRequest struct {
	Name        string `json:"name"`
	Category    string `json:"category"`
	Building_ID uint   `json:"building_id"`
	Surface     uint   `json:"surface"`
	Capacity    uint   `json:"capacity"`
}

func (Unit *UnitRequest) ToDomain() *units.Domain {
	return &units.Domain{
		Name: Unit.Name,
		Category: Unit.Category,
		Building_ID: Unit.Building_ID,
		Surface: Unit.Surface,
		Capacity: Unit.Capacity,
	}
}