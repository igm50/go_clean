package main

import (
	"database/sql"
	"log"
)

type cityRepository interface {
	fetchAll() []city
}

type cityRepositoryImpl struct {
	database *cityDatabase
}

func newCityRepository(c *cityDatabase) *cityRepositoryImpl {
	return &cityRepositoryImpl{c}
}

func (c *cityRepositoryImpl) fetchAll() []city {
	defer func() {
		if r := recover(); r != nil {
			log.Fatalln(r)
			panic(r)
		}
	}()

	return c.database.fetchAll()
}

type cityDatabase struct {
	db *sql.DB
}

func newCityDatabase(db *sql.DB) *cityDatabase {
	return &cityDatabase{db}
}

func (c *cityDatabase) fetchAll() []city {
	stmt, sErr := c.db.Prepare("SELECT * FROM city")
	if sErr != nil {
		panic(sErr)
	}

	rows, qErr := stmt.Query()
	if qErr != nil {
		panic(qErr)
	}
	defer rows.Close()

	var result []city
	for rows.Next() {
		var (
			id          int
			name        string
			countryCode string
			district    string
			population  int
		)

		if err := rows.Scan(&id, &name, &countryCode, &district, &population); err != nil {
			panic(err)
		}

		result = append(result, newCity(id, name, countryCode, district, population))
	}

	return result
}
