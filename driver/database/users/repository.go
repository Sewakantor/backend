package users

import (
	"capstone/business/users"
	"context"
	"errors"

	"gorm.io/gorm"
)

type UserRepository struct {
	DB *gorm.DB
}

func NewUserRepository(DB *gorm.DB) *UserRepository {
	return &UserRepository{DB: DB}
}

func (Repo *UserRepository) Register(ctx context.Context, domain *users.Domain) (users.Domain, error) {
	user := FromDomain(*domain)
	err := Repo.DB.Create(&user).Joins("Role")
	if err.Error != nil {
		return users.Domain{}, err.Error
	}
	return user.ToDomain(), nil
} 

func (Repo *UserRepository) GetEmail(ctx context.Context, email string, password string) (users.Domain, error){
	var user User
	err := Repo.DB.Find(&user, "email=?", email)
	if err.Error != nil {
		return users.Domain{}, errors.New("email not registered")
	}
	if user.Password != password {
        return users.Domain{}, errors.New("wrong password")
    }
	return user.ToDomain(), nil
	
}