package main

import (
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func fetchKabupaten(provinsi string) map[string]string {
	req, _ := http.NewRequest("GET", fmt.Sprintf("https://pemilu2019.kpu.go.id/static/json/wilayah/%s.json", provinsi), nil)

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

	var kabupaten = make(map[string]string)
	for kode, data := range tmp {
		kabupaten[kode] = data.Nama
	}

	return kabupaten
}
