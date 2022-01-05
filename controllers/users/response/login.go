package response

import "capstone/business/users"

func UserLogin(domain users.Domain, token string) JWTResponse {
	Response := UserResponse {
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

	TokenResponse := JWTResponse{}
	TokenResponse.Token = token
	TokenResponse.User =Response
	return TokenResponse
}