package units

import (
	"context"
	"errors"
	"time"
)

type UnitUseCase struct {
	repo UnitRepoInterface
	ctx  time.Duration
}

func NewUnitUseCase(unitRepo UnitRepoInterface, ctx time.Duration) *UnitUseCase {
	return &UnitUseCase{
		repo: unitRepo,
		ctx:  ctx,
	}
}

func (usecase *UnitUseCase) Add (ctx context.Context, domain Domain) (Domain, error) {
	if domain.Name == "" {
		return Domain{}, errors.New("Silahkan isi nama unit")
	}
	if domain.Category == "" {
		return Domain{}, errors.New("Silahkan isi category unit")
	}
	if domain.Building_ID == 0 {
		return Domain{}, errors.New("Silahkan isi building id")
	}
	if domain.Surface == 0 {
		return Domain{}, errors.New("Silahkan isi surface unit")
	}
	if domain.Capacity == 0 {
		return Domain{}, errors.New("Silahkan isi capacity unit")
	}

	unit, err := usecase.repo.Add(ctx, domain)
	if err != nil {
		return Domain{}, err
	}
	return unit, nil
}

func (usecase *UnitUseCase) GetAll (ctx context.Context) ([]Domain, error) {
	unit, err := usecase.repo.GetAll(ctx)
	if err != nil {
		return []Domain{}, errors.New("Unit Tidak Ada")
	}
	return unit, nil
}

func (usecase *UnitUseCase) GetByID(ctx context.Context, id uint) (Domain, error){
	unit, err := usecase.repo.GetByID(ctx, id)
	if err != nil {
		return Domain{}, errors.New("Tidak ada unit dengan ID tersebut")
	}
	if id == 0 {
		return Domain{}, errors.New("ID harus diisi")
	}
	return unit, nil
}

func (usecase *UnitUseCase) Update(ctx context.Context, id uint, domain Domain) (Domain, error){
	domain.Id = (id)
	unit, err := usecase.repo.Update(ctx, id, domain)
	if err != nil {
		return Domain{}, errors.New("Tidak ada unit dengan ID tersebut")
	}
	return unit, nil
}

func (usecase *UnitUseCase) Delete(ctx context.Context, id uint) error {
	err := usecase.repo.Delete(ctx, id)
	if err != nil {
		return errors.New("Tidak ada unit dengan ID tersebut")
	}
	return nil
}