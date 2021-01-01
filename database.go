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

func (d *cityDatabase) fetchAll() []city {
	stmt, sErr := d.db.Prepare("SELECT * FROM city")
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
