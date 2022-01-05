package users

import (
	"capstone/app/middlewares"
	"context"
	"errors"
	"log"
	"time"
)

type UserUseCase struct {
	repo UserRepoInterface
	ctx  time.Duration
	JWTAuth *middlewares.ConfigJWT
	
}

func NewUserUseCase(userRepo UserRepoInterface, ctx time.Duration, JWTAuth *middlewares.ConfigJWT) *UserUseCase {
	return &UserUseCase{
		repo: 		userRepo,
		ctx:		ctx,
	}
}

func (usecase *UserUseCase) Register(ctx context.Context, domain Domain) (Domain, error) {
	if domain.Name == "" {
		return Domain{}, errors.New("Nama Belum di Isi")
	}
	if domain.Email == "" {
		return Domain{}, errors.New("email belum di isi")
	}
	if domain.Password == "" {
		return Domain{}, errors.New("password belum di isi")
	}
	if domain.Company == "" {
		return Domain{}, errors.New("Company belum di isi")
	}
	if domain.Phone == "" {
		return Domain{}, errors.New("Phone Number Belum di Isi")
	}
	user, err := usecase.repo.Register(ctx, &domain)
	if err != nil {
		return Domain{}, err
	}
	
	return user, nil
}

func (usecase *UserUseCase) Login(ctx context.Context, email string, password string, ) (Domain, string, error){
	if email == "" {
		return Domain{},"", errors.New("email belum di isi")
	}
	if password == "" {
		return Domain{},"", errors.New("password belum di isi")
	
	}
	user, err := usecase.repo.GetEmail(ctx, email, password)
	if err != nil {
		return Domain{},"", err
	}
	token, errToken := usecase.JWTAuth.GenerateTokenJWT(user.Id, user.Email, user.Name, user.Phone , user.Roles_ID)
	if errToken != nil {
		log.Println(errToken)
	}
	if token == "" {
		return Domain{}, "", errors.New("Token Kosong")
	}
	return user, token, nil
}