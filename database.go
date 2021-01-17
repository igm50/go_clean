package main

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

type cityDatabase struct {
	db *sql.DB
}

func newCityDatabase(db *sql.DB) cityDatabase {
	return cityDatabase{db}
}

func (d *cityDatabase) fetchByID(id int) city {
	row := d.db.QueryRow("SELECT * FROM city WHERE id = ?", id)

	var (
		fetchedID   int
		name        string
		countryCode string
		district    string
		population  int
	)

	if err := row.Scan(&fetchedID, &name, &countryCode, &district, &population); err != nil {
		panic(err)
	}

	return newCity(cityID(fetchedID), name, countryCode, district, population)
}
