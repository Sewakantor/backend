package response

import (
	"capstone/business/complexs"
	"time"

	"gorm.io/gorm"
)

type ComplexResponse struct {
	Id        uint   `json:"id"`
	Name      string `json:"name"`
	Address   string `json:"address"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	DeletedAt gorm.DeletedAt `json:"deletedAt"`
}

func FromDomainComplex(domain complexs.Domain) ComplexResponse {
	return ComplexResponse{
		Id:        domain.Id,
		Name:      domain.Name,
		Address:   domain.Address,
		CreatedAt: domain.CreatedAt,
		UpdatedAt: domain.UpdatedAt,
	}
}

func GetAll(domain []complexs.Domain) []ComplexResponse {
	var GetAll []ComplexResponse
	for _, val := range domain {
		GetAll = append(GetAll, FromDomainComplex(val))
	}
	return GetAll
}
