package main

import "strconv"

type fetchCityInput struct {
	id *cityID // nilable
}

func newFetchCityInput(v map[string]string) fetchCityInput {
	var cityID *cityID

	idStr, ok := v["id"]
	if !ok {
		return fetchCityInput{cityID}
	}

	id, e := strconv.Atoi(idStr)
	if e != nil {
		panic(e)
	}

	c := newCityID(id)
	cityID = &c

	return fetchCityInput{cityID}
}
