package main

import (
	"encoding/json"
	"fmt"
)

func fetchKelurahan(provinsi, kabupaten, kecamatan string) map[string]string {
	body := requestData(fmt.Sprintf("https://pemilu2019.kpu.go.id/static/json/wilayah/%s/%s/%s.json", provinsi, kabupaten, kecamatan))

	var tmp map[string]struct {
		Nama string `json:"nama"`
	}
	json.Unmarshal(body, &tmp)

	var kelurahan = make(map[string]string)
	for kode, data := range tmp {
		kelurahan[kode] = data.Nama
	}

	return kelurahan
}
