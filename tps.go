package main

import (
	"encoding/json"
	"fmt"
)

func fetchTps(provinsi, kabupaten, kecamatan, kelurahan string) map[string]string {
	m.Lock()
	defer m.Unlock()

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

	var tmp struct {
		Chart struct {
			Jokowi  int `json:"21"`
			Prabowo int `json:"22"`
		} `json:"chart"`
		TotalDPT      int `json:"pemilih_j"`
		TotalPengguna int `json:"pengguna_j"`
		SuaraSah      int `json:"suara_sah"`
		SuaraTidakSah int `json:"suara_tidak_sah"`
		SuaraTotal    int `json:"suara_total"`
	}
	err := json.Unmarshal(body, &tmp)
	if err != nil || string(body) == "{}" {
		return nil, fmt.Errorf(NotFound)
	}

	var formC1 FormulirC1
	formC1.PerolehanSuara.Jokowi = tmp.Chart.Jokowi
	formC1.PerolehanSuara.Prabowo = tmp.Chart.Prabowo
	formC1.TotalDPT = tmp.TotalDPT
	formC1.TotalPengguna = tmp.TotalPengguna
	formC1.SuaraSah = tmp.SuaraSah
	formC1.SuaraTidakSah = tmp.SuaraTidakSah
	formC1.SuaraTotal = tmp.SuaraTotal

	if formC1.PerolehanSuara.Jokowi+formC1.PerolehanSuara.Prabowo != formC1.SuaraSah {
		return &formC1, fmt.Errorf("SALAH. %d+%d bukan %d, bisa ngitungnya gak sih",
			formC1.PerolehanSuara.Jokowi, formC1.PerolehanSuara.Prabowo, formC1.SuaraSah)
	}

	return &formC1, nil
}
