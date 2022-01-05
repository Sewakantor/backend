package addresses

import (
	"capstone/business/addresses"
	"context"
	"errors"

	"gorm.io/gorm"
)

type AddressRepository struct {
	DB *gorm.DB
}

func NewAddressRepository(DB *gorm.DB) *AddressRepository {
	return &AddressRepository{DB: DB}
}

func(Repo *AddressRepository) Add(ctx context.Context, domain addresses.Domain) (addresses.Domain, error) {
	address := FromDomain(*&domain)
	err := Repo.DB.Create(&address)
	if err.Error != nil {
		return addresses.Domain{}, err.Error
	}
	return address.ToDomain(), nil
}

func (Repo *AddressRepository) GetAll(ctx context.Context) ([]addresses.Domain, error) {
	var address []Address
	err := Repo.DB.Find(&address)
	if err.Error != nil {
		return []addresses.Domain{}, err.Error
	}
	return GetAllAddress(address), nil
}

func (Repo *AddressRepository) Update(ctx context.Context, id uint, domain addresses.Domain) (addresses.Domain, error){
	address := FromDomain(domain)
	if Repo.DB.Save(&address).Error != nil {
		return addresses.Domain{}, errors.New("id tidak ditemukan")
	}
	return address.ToDomain(), nil
}


func (Repo *AddressRepository) Delete(ctx context.Context, id uint) error {
	address := Address{}
	err := Repo.DB.Delete(&address, id)
	if err.Error != nil {
		return err.Error

	}
	if err.RowsAffected == 0 {
		return errors.New("data kosong")
	}
	return nil
}