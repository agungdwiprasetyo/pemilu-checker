package main

import (
	"encoding/json"
	"fmt"
)

func fetchTps(provinsi, kabupaten, kecamatan, kelurahan string) map[string]string {
	body := requestData(fmt.Sprintf("https://pemilu2019.kpu.go.id/static/json/wilayah/%s/%s/%s/%s.json", provinsi, kabupaten, kecamatan, kelurahan))

	var tmp map[string]struct {
		Nama string `json:"nama"`
	}
	json.Unmarshal(body, &tmp)

	var tps = make(map[string]string)
	for kode, data := range tmp {
		tps[kode] = data.Nama
	}

	return tps
}

func detailTps(prefix string) (*FormulirC1, error) {
	m.Lock()
	defer m.Unlock()

	url := fmt.Sprintf("https://pemilu2019.kpu.go.id/static/json/hhcw/ppwp/%s", prefix)
	body := requestData(url)

	var formC1 FormulirC1
	err := json.Unmarshal(body, &formC1)
	if err != nil || string(body) == "{}" {
		return nil, fmt.Errorf("Data belum tersedia")
	}

	if formC1.Chart.Jokowi+formC1.Chart.Prabowo != formC1.SuaraSah {
		return nil, fmt.Errorf("SALAH. %d+%d bukan %d, bisa ngitungnya gak sih", formC1.Chart.Jokowi, formC1.Chart.Prabowo, formC1.SuaraSah)
	}

	return &formC1, nil
}
