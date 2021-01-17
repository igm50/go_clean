package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func buildCityController(usecase *cityUseCase, cityPresenter *cityPresenter, errorPresenter *errorPresenter) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		params := mux.Vars(r)

		defer func() {
			if e := recover(); e != nil {
				log.Print("Error occurred in city controller. params: ", params, ", error: ", e)
				errorPresenter.write(w, e)
				return
			}
		}()

		input := newFetchCityInput(params)

		u := *usecase
		output := u.fetch(input)

		cityPresenter.write(w, output)
	}
}
