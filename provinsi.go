package main

import (
	"crypto/tls"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

func fetchWilayah() map[string]string {
	req, _ := http.NewRequest("GET", "https://pemilu2019.kpu.go.id/static/json/wilayah/0.json", nil)

	client := new(http.Client)
	client.Transport = &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

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
