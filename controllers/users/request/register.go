package request

import "capstone/business/users"

type RegisterUserRequest struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Photo 	 string `json:"photo"`
	Password string `json:"password"`
	Company  string `json:"company"`
	Phone 	 string `json:"phone"`
	Status 	 int 	`json:"status"`
	Roles_ID uint 	`json:"roles_id"`
}

func (user *RegisterUserRequest) ToDomain() *users.Domain {
	return &users.Domain{
		Name: user.Name,
		Email: user.Email,
		Photo: user.Photo,
		Password: user.Password,
		Company: user.Company,
		Phone: user.Phone,
		Status: user.Status,
		Roles_ID: user.Roles_ID,
	}
}