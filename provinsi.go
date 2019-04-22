package main

import (
	"encoding/json"
)

func fetchWilayah() map[string]string {
	body := requestData("https://pemilu2019.kpu.go.id/static/json/wilayah/0.json")

	var tmp map[string]struct {
		Nama string `json:"nama"`
	}
	json.Unmarshal(body, &tmp)

	var wilayah = make(map[string]string)
	for kode, data := range tmp {
		wilayah[kode] = data.Nama
		// wilayah = append(wilayah, &model.Wilayah{Kode: kode, Nama: data.Nama})
	}

	return wilayah
}
