package request

import "capstone/business/buildings"

type BuildingRequest struct {
	Name          string `json:"name"`
	Complex_ID    uint   `json:"complex_id"`
	Photo         string `json:"photo"`
	Year          uint32 `json:"year"`
	Floors        uint32 `json:"floors"`
	Floor_Surface uint32 `json:"floor_surface"`
	Total_Surface uint32 `json:"total_surface"`
}

func (Building *BuildingRequest) ToDomain() *buildings.Domain {
	return &buildings.Domain{
		Name: Building.Name,
		Complex_ID: Building.Complex_ID,
		Photo: Building.Photo,
		Year: Building.Year,
		Floors: Building.Floors,
		Floor_Surface: Building.Floor_Surface,
		Total_Surface: Building.Total_Surface,
	}
}