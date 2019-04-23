package main

import (
	"encoding/json"
)

func fetchProvinsi() map[string]string {
	m.Lock()
	defer m.Unlock()

	body := requestData("https://pemilu2019.kpu.go.id/static/json/wilayah/0.json")

	var tmp map[string]struct {
		Nama string `json:"nama"`
	}
	json.Unmarshal(body, &tmp)

	var prov = make(map[string]string)
	for kode, data := range tmp {
		prov[kode] = data.Nama
	}

	return prov
}
