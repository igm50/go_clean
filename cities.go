package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

func buildCityHandler(db *sql.DB) func(http.ResponseWriter, *http.Request) {
	cityDB := newCityDatabase(db)
	repositoy := newCityRepository(cityDB)

	return func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if r := recover(); r != nil {
				w.WriteHeader(http.StatusInternalServerError)
				log.Println(r)
				fmt.Fprintf(w, "{error: %+v}", r)
				return
			}
		}()

		cities := repositoy.fetchAll()
		result := newCityOutput(citiesToDatas(cities))

		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.Header().Set("sample", "sample")
		w.WriteHeader(http.StatusOK)

		b, e := json.Marshal(result)
		if e != nil {
			panic(e)
		}

		w.Write(b)
	}
}
