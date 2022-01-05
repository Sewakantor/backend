package facilities

import (
	"context"
	"errors"
	"time"
)

type FacilityUseCase struct {
	repo FacilityRepoInterface
	ctx  time.Duration
}

func NewFacilityUseCase(facilityRepo FacilityRepoInterface, ctx time.Duration) *FacilityUseCase {
	return &FacilityUseCase{
		repo: facilityRepo,
		ctx: ctx,
	}
}

func (usecase *FacilityUseCase) Add(ctx context.Context, domain Domain) (Domain, error) {
	if domain.Name == "" {
		return Domain{}, errors.New("Nama harus di isi")
	}
	if domain.Types == "" {
		return Domain{}, errors.New("Type harus di isi")
	}
	if domain.Lat == 0 {
		return Domain{}, errors.New("Latitude harus di isi")
	}
	if domain.Long == 0 {
		return Domain{}, errors.New("Longtitude harus di isi")
	}
	facility, err := usecase.repo.Add(ctx, domain)
	if err != nil {
		return Domain{}, err
	}

	return facility, nil
}

func (usecase *FacilityUseCase) GetAll(ctx context.Context) ([]Domain, error) {
	facility, err := usecase.repo.GetAll(ctx)
	if err != nil {
		return []Domain{}, errors.New("tidak ada facility")
	}
	return facility, nil
}

func (usecase *FacilityUseCase) Update(ctx context.Context, id uint, domain Domain) (Domain, error){
	domain.Id = (id)
	facility, err := usecase.repo.Update(ctx, id, domain)
	if err != nil {
		return Domain{}, errors.New("tidak ada facility dengan ID tersebut")
	}
	return facility, nil
}

func (usecase *FacilityUseCase) Delete(ctx context.Context, id uint) error{
	err := usecase.repo.Delete(ctx, id)
	if err != nil {
		return errors.New("tidak ada facility dengan ID tersebut")
	}
	return nil
}