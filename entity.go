package main

type city struct {
	id          int
	name        string
	countryCode string
	district    string
	population  int
}

func newCity(id int, name string, countryCode string, district string, population int) city {
	return city{id, name, countryCode, district, population}
}
