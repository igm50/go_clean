package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type cityData struct {
	ID          int
	Name        string
	CountryCode string
	District    string
	Population  int
}

type cityPresenter struct{}

func newCityPresenter() cityPresenter {
	return cityPresenter{}
}

func (presenter *cityPresenter) write(w http.ResponseWriter, o fetchCityOutput) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusOK)

	type outputData struct {
		Cities []cityData
	}

	var cityDatas []cityData
	for _, city := range o.Cities {
		data := cityData{city.id, city.name, city.countryCode, city.district, city.population}
		cityDatas = append(cityDatas, data)
	}

	output := outputData{cityDatas}

	b, e := json.Marshal(output)
	if e != nil {
		panic(e)
	}

	w.Write(b)
}

type errorPresenter struct{}

func newErrorPresenter() errorPresenter {
	return errorPresenter{}
}

func (presenter *errorPresenter) write(w http.ResponseWriter, e interface{}) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusInternalServerError)

	log.Println(e)
	fmt.Fprintf(w, "{error: %+v}", e)
}
