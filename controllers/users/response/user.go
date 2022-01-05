package response

import (
	"capstone/business/users"
	"time"

	"gorm.io/gorm"
)

type UserResponse struct {
	Id uint 			`json:"id"`
	// User_ID   uint
	Name      string 	`json:"name"`
	Email     string 	`json:"email"`
	Photo     string 	`json:"photo"`
	Company   string 	`json:"company"`
	Phone     string 	`json:"phone"`
	Status    int 		`json:"status"`
	Roles_ID  uint 		`json:"roles_id"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	DeletedAt gorm.DeletedAt `json:"deletedAt"`
}

type JWTResponse struct {
	Token string		`json:"token"`
	User interface{}	`json:"user"`
}

func FromDomain(domain users.Domain) UserResponse {
	return UserResponse{
		Id: domain.Id,
		Name: domain.Name,
		Email: domain.Email,
		Photo: domain.Photo,
		Company: domain.Company,
		Phone: domain.Phone,
		Status: domain.Status,
		Roles_ID: domain.Roles_ID,
		CreatedAt: domain.CreatedAt,
		UpdatedAt: domain.UpdatedAt,
	}
}

func GetAllUsers(domain []users.Domain) []UserResponse {
	var GetAllUsers []UserResponse
	for _, val := range domain{
		GetAllUsers = append(GetAllUsers, FromDomain(val))
	}
	return GetAllUsers 
}