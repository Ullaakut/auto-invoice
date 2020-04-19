package workhours

import (
	"io/ioutil"
	"encoding/json"
)

func Get(_ Source) (Report, error) {
	var rep Report

	contents, err := ioutil.ReadFile("test_assets/response.json")
	if err != nil {
		return rep, err
	}

	err = json.Unmarshal(contents, &rep)
	if err != nil {
		return rep, err
	}

	return rep, nil
}
