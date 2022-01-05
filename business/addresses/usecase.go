package addresses

import (
	"context"
	"errors"
	"time"
)

type AddressUseCase struct {
	repo AddressRepoInterface
	ctx  time.Duration
}

func NewAddressUseCase(addressRepo AddressRepoInterface, ctx time.Duration) *AddressUseCase {
	return &AddressUseCase{
		repo: addressRepo,
		ctx: ctx,
	}
}

func (usecase *AddressUseCase) Add(ctx context.Context, domain Domain) (Domain, error) {
	if domain.Street == "" {
		return Domain{}, errors.New("Street harus di isi")
	}
	if domain.City == "" {
		return Domain{}, errors.New("City harus di isi")
	}
	if domain.State == "" {
		return Domain{}, errors.New("State harus di isi")
	}
	if domain.Country == "" {
		return Domain{}, errors.New("Country harus di isi")
	}
	if domain.PostalCode == "" {
		return Domain{}, errors.New("PostalCode harus di isi")
	}
	if domain.Latitude == 0 {
		return Domain{}, errors.New("Latitude harus di isi")
	}
	if domain.Longtitude == 0 {
		return Domain{}, errors.New("Longtitude harus di isi")
	}
	address, err := usecase.repo.Add(ctx, domain)
	if err != nil {
		return Domain{}, err
	}

	return address, nil
}

func (usecase *AddressUseCase) GetAll(ctx context.Context) ([]Domain, error) {
	address, err := usecase.repo.GetAll(ctx)
	if err != nil {
		return []Domain{}, errors.New("tidak ada address")
	}
	return address, nil
}

func (usecase *AddressUseCase) Update(ctx context.Context, id uint, domain Domain) (Domain, error){
	domain.Id = (id)
	building, err := usecase.repo.Update(ctx, id, domain)
	if err != nil {
		return Domain{}, errors.New("tidak ada address dengan ID tersebut")
	}
	return building, nil
}

func (usecase *AddressUseCase) Delete(ctx context.Context, id uint) error{
	err := usecase.repo.Delete(ctx, id)
	if err != nil {
		return errors.New("tidak ada address dengan ID tersebut")
	}
	return nil
}