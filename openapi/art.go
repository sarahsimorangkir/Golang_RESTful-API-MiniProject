package openapi

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type Art struct {
	departementId int    `json:"departement_id"`
	displayName   string `json:"display_name"`
}

func GetArt() (Art, error) {
	// akses url
	resp, err := http.Get("https://collectionapi.metmuseum.org/public/collection/v1/departments")

	if err != nil {
		panic(err)
	}

	// ambil data dari response body
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	// masukkan data ke struct
	var res Art
	err = json.Unmarshal(body, &res)
	if err != nil {
		panic(err)
	}

	return res, nil
}
