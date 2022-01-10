package units

import (
	"capstone/business/units"
	"context"
	"errors"

	"gorm.io/gorm"
)

type UnitRepository struct {
	DB *gorm.DB
}

func NewUnitRepository(DB *gorm.DB) *UnitRepository {
	return &UnitRepository{DB: DB}
}

func (Repo *UnitRepository) Add(ctx context.Context, domain units.Domain) (units.Domain, error) {
	unit := FromDomain(*&domain)
	err := Repo.DB.Create(&unit)
	if err.Error != nil {
		return units.Domain{}, err.Error
	}
	return unit.ToDomain(), nil
}

func (Repo *UnitRepository) GetAll(ctx context.Context) ([]units.Domain, error) {
	var unit []Unit
	err := Repo.DB.Preload("Building").Find(&unit)
	if err.Error != nil {
		return []units.Domain{}, err.Error
	}
	return GetAll(unit), nil
}

func (Repo *UnitRepository) GetByID(ctx context.Context, id uint) (units.Domain, error){
	var unit Unit
	err := Repo.DB.Preload("Building").Find(&unit, "id=?", id)
	if err.Error != nil {
		return units.Domain{}, err.Error
	}
	return unit.ToDomain(), nil
}

func (Repo *UnitRepository) Update(ctx context.Context, id uint, domain units.Domain) (units.Domain, error){
	unit := FromDomain(domain)
	if Repo.DB.Save(&unit).Error != nil {
		return units.Domain{}, errors.New("id tidak ditemukan")
	}
	return unit.ToDomain(), nil
}


func (Repo *UnitRepository) Delete(ctx context.Context, id uint) error {
	unit := Unit{}
	err := Repo.DB.Delete(&unit, id)
	if err.Error != nil {
		return err.Error

	}
	if err.RowsAffected == 0 {
		return errors.New("data kosong")
	}
	return nil
}