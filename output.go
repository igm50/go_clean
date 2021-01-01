package main

type fetchCityOutput struct {
	Cities []city
}

func newFetchCityOutput(c []city) fetchCityOutput {
	return fetchCityOutput{c}
}
