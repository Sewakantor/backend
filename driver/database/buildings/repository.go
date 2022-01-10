package buildings

import (
	"capstone/business/buildings"
	"context"
	"errors"

	"gorm.io/gorm"
)

type BuildingRepository struct {
	DB *gorm.DB
}

func NewBuildingRepository(DB *gorm.DB) *BuildingRepository {
	return &BuildingRepository{DB: DB}
}

func (Repo *BuildingRepository) Add(ctx context.Context, domain buildings.Domain) (buildings.Domain, error) {
	building := FromDomain(*&domain)
	err := Repo.DB.Create(&building)
	if err.Error != nil {
		return buildings.Domain{}, err.Error
	}
	return building.ToDomain(), nil
}

func (Repo *BuildingRepository) GetAll(ctx context.Context) ([]buildings.Domain, error) {
	var building []Building
	err := Repo.DB.Preload("Complex").Find(&building)
	if err.Error != nil {
		return []buildings.Domain{}, err.Error
	}
	return GetAll(building), nil
}

func (Repo *BuildingRepository) GetByID(ctx context.Context, id uint) (buildings.Domain, error){
	var building Building
	err := Repo.DB.Preload("Complex").Find(&building, "id=?", id)
	if err.Error != nil {
		return buildings.Domain{}, err.Error
	}
	return building.ToDomain(), nil
}

func (Repo *BuildingRepository) GetByComplexID(complexid uint, ctx context.Context ) ([]buildings.Domain, error){
	var building []Building
	err := Repo.DB.Preload("Complex").Find(&building, "complex_id=?", complexid)
	if err.Error != nil {
		return []buildings.Domain{}, err.Error
	}
	return GetAll(building), nil
}

func (Repo *BuildingRepository) Update(ctx context.Context, id uint, domain buildings.Domain) (buildings.Domain, error){
	building := FromDomain(domain)
	if Repo.DB.Save(&building).Error != nil {
		return buildings.Domain{}, errors.New("id tidak ditemukan")
	}
	return building.ToDomain(), nil
}


func (Repo *BuildingRepository) Delete(ctx context.Context, id uint) error {
	building := Building{}
	err := Repo.DB.Delete(&building, id)
	if err.Error != nil {
		return err.Error

	}
	if err.RowsAffected == 0 {
		return errors.New("data kosong")
	}
	return nil
}