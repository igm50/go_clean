package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

func buildCityController(usecase *cityUseCase, cityPresenter *cityPresenter, errorPresenter *errorPresenter) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if e := recover(); e != nil {
				errorPresenter.write(w, e)
				return
			}
		}()

		input := newFetchCityInput(mux.Vars(r))

		u := *usecase
		output := u.fetch(input)

		cityPresenter.write(w, output)
	}
}
