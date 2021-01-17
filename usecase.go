package main

type cityUseCase interface {
	fetch(i fetchCityInput) fetchCityOutput
}

type cityInteractor struct {
	repository *cityRepository
}

func newCityInteractor(r *cityRepository) cityUseCase {
	return cityInteractor{r}
}

func (i cityInteractor) fetch(f fetchCityInput) fetchCityOutput {
	r := *i.repository
	city := r.fetchByID(f)

	return newFetchCityOutput(city)
}
