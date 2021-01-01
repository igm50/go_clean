package main

import (
	"log"
)

type cityRepository interface {
	fetchAll() []city
}

type cityRepositoryImpl struct {
	database *cityDatabase
}

func newCityRepositoryImpl(d *cityDatabase) cityRepository {
	return cityRepositoryImpl{d}
}

func (r cityRepositoryImpl) fetchAll() []city {
	defer func() {
		if r := recover(); r != nil {
			log.Fatalln(r)
			panic(r)
		}
	}()

	return r.database.fetchAll()
}
