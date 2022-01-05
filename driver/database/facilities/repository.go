package facilities

import (
	"capstone/business/facilities"
	"context"
	"errors"

	"gorm.io/gorm"
)

type FacilityRepository struct {
	DB *gorm.DB
}

func NewFacilityRepository(DB *gorm.DB) *FacilityRepository {
	return &FacilityRepository{DB: DB}
}

func(Repo *FacilityRepository) Add(ctx context.Context, domain facilities.Domain) (facilities.Domain, error) {
	facility := FromDomain(*&domain)
	err := Repo.DB.Create(&facility)
	if err.Error != nil {
		return facilities.Domain{}, err.Error
	}
	return facility.ToDomain(), nil
}

func (Repo *FacilityRepository) GetAll(ctx context.Context) ([]facilities.Domain, error) {
	var facility []Facility
	err := Repo.DB.Find(&facility)
	if err.Error != nil {
		return []facilities.Domain{}, err.Error
	}
	return GetAllFacility(facility), nil
}

func (Repo *FacilityRepository) Update(ctx context.Context, id uint, domain facilities.Domain) (facilities.Domain, error){
	facility := FromDomain(domain)
	if Repo.DB.Save(&facility).Error != nil {
		return facilities.Domain{}, errors.New("id tidak ditemukan")
	}
	return facility.ToDomain(), nil
}


func (Repo *FacilityRepository) Delete(ctx context.Context, id uint) error {
	facility := Facility{}
	err := Repo.DB.Delete(&facility, id)
	if err.Error != nil {
		return err.Error
	}
	if err.RowsAffected == 0 {
		return errors.New("data kosong")
	}
	return nil
}