package main

import (
	"encoding/json"
	"fmt"
)

func fetchKecamatan(provinsi, kabupaten string) map[string]string {
	m.Lock()
	defer m.Unlock()

	body := requestData(fmt.Sprintf("https://pemilu2019.kpu.go.id/static/json/wilayah/%s/%s.json", provinsi, kabupaten))

	var tmp map[string]struct {
		Nama string `json:"nama"`
	}
	json.Unmarshal(body, &tmp)

	var kecamatan = make(map[string]string)
	for kode, data := range tmp {
		kecamatan[kode] = data.Nama
	}

	return kecamatan
}
