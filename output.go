package main

type cityOutput struct {
	Cities []cityData
}

func newCityOutput(c []cityData) cityOutput {
	return cityOutput{c}
}

type cityData struct {
	ID          int
	Name        string
	CountryCode string
	District    string
	Population  int
}

func newCityData(c city) cityData {
	return cityData{c.id, c.name, c.countryCode, c.district, c.population}
}

func citiesToDatas(cities []city) []cityData {
	var datas []cityData

	for _, city := range cities {
		datas = append(datas, newCityData(city))
	}

	return datas
}
