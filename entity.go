package main

type cityID int

func newCityID(id int) cityID {
	return cityID(id)
}

type city struct {
	id          cityID
	name        string
	countryCode string
	district    string
	population  int
}

func newCity(id cityID, name string, countryCode string, district string, population int) city {
	return city{id, name, countryCode, district, population}
}

func (c city) getID() int {
	return int(c.id)
}

func (c city) getName() string {
	return c.name
}

func (c city) getCountryCode() string {
	return c.countryCode
}

func (c city) getDistrict() string {
	return c.district
}

func (c city) getPopulation() int {
	return c.population
}
