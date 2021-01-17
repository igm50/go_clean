package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type cityData struct {
	ID          int
	Name        string
	CountryCode string
	District    string
	Population  int
}

// cityPresenter
type cityPresenter struct{}

func newCityPresenter() cityPresenter {
	return cityPresenter{}
}

func (presenter *cityPresenter) write(w http.ResponseWriter, o fetchCityOutput) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusOK)

	city := o.city
	data := cityData{city.getID(), city.getName(), city.getCountryCode(), city.getDistrict(), city.getPopulation()}

	b, e := json.Marshal(data)
	if e != nil {
		panic(e)
	}

	w.Write(b)
}

// errorPresenter
type errorPresenter struct{}

func newErrorPresenter() errorPresenter {
	return errorPresenter{}
}

func (presenter *errorPresenter) write(w http.ResponseWriter, e interface{}) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusInternalServerError)

	fmt.Fprintf(w, "{error: %+v}", e)
}
