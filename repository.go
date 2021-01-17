package main

type cityRepository interface {
	fetchByID(i fetchCityInput) city
}

type cityRepositoryImpl struct {
	database *cityDatabase
}

func newCityRepositoryImpl(d *cityDatabase) cityRepository {
	return cityRepositoryImpl{d}
}

func (r cityRepositoryImpl) fetchByID(i fetchCityInput) city {
	return r.database.fetchByID(int(*i.id))
}
