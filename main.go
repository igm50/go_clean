package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/mux"
)

var (
	addr = os.Getenv("ADDRESS")
)

func main() {

	// build infrastructure
	db := dbConnection()
	cityDatabase := newCityDatabase(db)

	// build usecase layer
	cityRepository := newCityRepositoryImpl(&cityDatabase)
	cityUsecase := newCityInteractor(&cityRepository)

	// build controller layer
	cityPresenter := newCityPresenter()
	errorPresenter := newErrorPresenter()
	cityController := buildCityController(&cityUsecase, &cityPresenter, &errorPresenter)

	// build server
	r := mux.NewRouter()
	r.HandleFunc("/", homeHandler)
	r.HandleFunc("/cities", cityController)

	srv := http.Server{
		Handler:      r,
		Addr:         addr,
		ReadTimeout:  15 * time.Second,
		WriteTimeout: 15 * time.Second,
	}

	// start server
	log.Println("Start server.")
	log.Fatal(srv.ListenAndServe())
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Hello world!")
}
