package buildings

import (
	"context"
	"errors"
	"time"
)

type BuildingUseCase struct {
	repo BuildingRepoInterface
	ctx  time.Duration
}

func NewBuildingUseCase(buildingRepo BuildingRepoInterface, ctx time.Duration) *BuildingUseCase {
	return &BuildingUseCase{
		repo: buildingRepo,
		ctx:  ctx,
	}
}

func (usecase *BuildingUseCase) Add (ctx context.Context, domain Domain) (Domain, error) {
	if domain.Name == "" {
		return Domain{}, errors.New("Silahkan isi nama complex")
	}
	if domain.Complex_ID == 0 {
		return Domain{}, errors.New("Silahkan isi street complex")
	}
	if domain.Photo == "" {
		return Domain{}, errors.New("Silahkan isi city complex")
	}
	if domain.Year == 0 {
		return Domain{}, errors.New("Silahkan isi state complex")
	}
	if domain.Floors == 0 {
		return Domain{}, errors.New("Silahkan isi country complex")
	}
	if domain.Floor_Surface == 0 {
		return Domain{}, errors.New("Silahkan isi postalcode complex")
	}
	if domain.Total_Surface == 0 {
		return Domain{}, errors.New("Silahkan isi latitude complex")
	}

	complex, err := usecase.repo.Add(ctx, domain)
	if err != nil {
		return Domain{}, err
	}
	return complex, nil
}

func (usecase *BuildingUseCase) GetAll (ctx context.Context) ([]Domain, error) {
	building, err := usecase.repo.GetAll(ctx)
	if err != nil {
		return []Domain{}, errors.New("Building Tidak Ada")
	}
	return building, nil
}

func (usecase *BuildingUseCase) GetByID(ctx context.Context, id uint) (Domain, error){
	building, err := usecase.repo.GetByID(ctx, id)
	if err != nil {
		return Domain{}, errors.New("Tidak ada building dengan ID tersebut")
	}
	if id == 0 {
		return Domain{}, errors.New("ID harus diisi")
	}
	return building, nil
}

func (usecase *BuildingUseCase) GetByComplexID(complexid uint, ctx context.Context) ([]Domain, error){
	building, err := usecase.repo.GetByComplexID(complexid, ctx)
	if err != nil {
		return []Domain{}, errors.New("tidak ada complex dengan ID tersebut")
	}
	if complexid == 0 {
		return []Domain{}, errors.New("complexID harus diisi")
	}
	return building, nil
}

func (usecase *BuildingUseCase) Update(ctx context.Context, id uint, domain Domain) (Domain, error){
	domain.Id = (id)
	building, err := usecase.repo.Update(ctx, id, domain)
	if err != nil {
		return Domain{}, errors.New("Tidak ada building dengan ID tersebut")
	}
	return building, nil
}

func (usecase *BuildingUseCase) Delete(ctx context.Context, id uint) error {
	err := usecase.repo.Delete(ctx, id)
	if err != nil {
		return errors.New("Tidak ada building dengan ID tersebut")
	}
	return nil
}