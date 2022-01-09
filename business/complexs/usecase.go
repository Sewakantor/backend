package complexs

import (
	"context"
	"errors"
	"time"
)

type ComplexUseCase struct {
	repo ComplexRepoInterface
	ctx  time.Duration
}

func NewComplexUseCase(complexRepo ComplexRepoInterface, ctx time.Duration) *ComplexUseCase {
	return &ComplexUseCase{
		repo: complexRepo,
		ctx: ctx,
	}
}

func (usecase *ComplexUseCase) Add (ctx context.Context, domain Domain) (Domain, error) {
	if domain.Name == "" {
		return Domain{}, errors.New("Silahkan isi nama complex")
	}
	if domain.Street == "" {
		return Domain{}, errors.New("Silahkan isi street complex")
	}
	if domain.City == "" {
		return Domain{}, errors.New("Silahkan isi city complex")
	}
	if domain.State == "" {
		return Domain{}, errors.New("Silahkan isi state complex")
	}
	if domain.Country == "" {
		return Domain{}, errors.New("Silahkan isi country complex")
	}
	if domain.PostalCode == "" {
		return Domain{}, errors.New("Silahkan isi postalcode complex")
	}
	if domain.Latitude == 0 {
		return Domain{}, errors.New("Silahkan isi latitude complex")
	}
	// if domain.Longtitude == 0 {
	// 	return Domain{}, errors.New("Silahkan isi longtitude complex")
	// }
	complex, err := usecase.repo.Add(ctx, domain)
	if err != nil {
		return Domain{}, err
	}
	return complex, nil
}

func (usecase *ComplexUseCase) GetAll (ctx context.Context) ([]Domain, error) {
	complex, err := usecase.repo.GetAll(ctx)
	if err != nil {
		return []Domain{}, errors.New("Complex Tidak Ada")
	}
	return complex, nil
}

func (usecase *ComplexUseCase) GetByID(ctx context.Context, id uint) (Domain, error){
	complex, err := usecase.repo.GetByID(ctx, id)
	if err != nil {
		return Domain{}, errors.New("Tidak ada complex dengan ID tersebut")
	}
	if id == 0 {
		return Domain{}, errors.New("ID harus diisi")
	}
	return complex, nil
}

func (usecase *ComplexUseCase) Update(ctx context.Context, id uint, domain Domain) (Domain, error){
	domain.Id = (id)
	complex, err := usecase.repo.Update(ctx, id, domain)
	if err != nil {
		return Domain{}, errors.New("Tidak ada complex dengan ID tersebut")
	}
	return complex, nil
}

func (usecase *ComplexUseCase) Delete(ctx context.Context, id uint) error {
	err := usecase.repo.Delete(ctx, id)
	if err != nil {
		return errors.New("Tidak ada Complex dengan ID tersebut")
	}
	return nil
}