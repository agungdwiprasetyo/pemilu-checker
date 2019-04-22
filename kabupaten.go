package main

import (
	"encoding/json"
	"fmt"
)

func fetchKabupaten(provinsi string) map[string]string {
	body := requestData(fmt.Sprintf("https://pemilu2019.kpu.go.id/static/json/wilayah/%s.json", provinsi))

	var tmp map[string]struct {
		Nama string `json:"nama"`
	}
	json.Unmarshal(body, &tmp)

	var kabupaten = make(map[string]string)
	for kode, data := range tmp {
		kabupaten[kode] = data.Nama
	}

	return kabupaten
}
