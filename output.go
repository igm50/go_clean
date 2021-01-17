package main

type fetchCityOutput struct {
	city city
}

func newFetchCityOutput(c city) fetchCityOutput {
	return fetchCityOutput{c}
}
