package request

import "capstone/business/roles"

type AddRoleRequest struct {
	Name string `json:"name"`
}

func (Role *AddRoleRequest) ToDomain() *roles.Domain{
	return &roles.Domain{
		Name: Role.Name,
	}
}