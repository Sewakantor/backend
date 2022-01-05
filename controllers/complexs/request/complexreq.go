package request

import "capstone/business/complexs"

type ComplexRequest struct {
	Name    string `json:"name"`
	Address string `json:"address"`
}

func (complex *ComplexRequest) ToDomain() *complexs.Domain {
	return &complexs.Domain{
		Name: complex.Name,
		Address: complex.Address,
	}
}