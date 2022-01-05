package complexs

import (
	"capstone/business/complexs"
	"context"
	"errors"

	"gorm.io/gorm"
)

type ComplexRepository struct {
	DB *gorm.DB
}

func NewComplexRepository (DB *gorm.DB) complexs.ComplexRepoInterface {
	return &ComplexRepository{DB :DB}
}

func (Repo *ComplexRepository) Add(ctx context.Context, domain complexs.Domain) (complexs.Domain, error) {
	complex := FromDomain(*&domain)
	err := Repo.DB.Create(&complex)
	if err.Error != nil {
		return complexs.Domain{}, err.Error
	}
	return complex.ToDomain(), nil
}

func (Repo *ComplexRepository) GetAll(ctx context.Context) ([]complexs.Domain, error){
	var complex []Complex
	err := Repo.DB.Find(&complex)
	if err.Error != nil {
		return []complexs.Domain{}, err.Error
	}
	return GetAll(complex), nil
}

func (Repo *ComplexRepository) GetByID(ctx context.Context , id uint) (complexs.Domain, error){
	var complex Complex
	err := Repo.DB.Find(&complex, "id=?", id)
	if err.Error != nil {
		return complexs.Domain{}, err.Error
	}
	return complex.ToDomain(), nil
}

func (Repo *ComplexRepository) Update(ctx context.Context, id uint, domain complexs.Domain) (complexs.Domain, error){
	complex := FromDomain(domain)
	if Repo.DB.Save(&complex).Error != nil {
		return complexs.Domain{}, errors.New("id tidak ditemukan")
	}
	return complex.ToDomain(), nil
}


func (Repo *ComplexRepository) Delete(ctx context.Context, id uint) error{
	complex := Complex{}
	err := Repo.DB.Delete(&complex, id)
	if err.Error!= nil {
		return err.Error
		
	}
	if err.RowsAffected == 0 {
		return errors.New("data kosong")
	}
	return nil
}