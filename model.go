package main

import (
	"fmt"
	"os"

	"github.com/agungdwiprasetyo/go-utils/debug"
)

type Wilayah struct {
	Provinsi struct {
		Kode, Nama string
	}
	Kabupaten struct {
		Kode, Nama string
	}
	Kecamatan struct {
		Kode, Nama string
	}
	Kelurahan struct {
		Kode, Nama string
	}
	TPS struct {
		Kode, Nama string
	}
}

func (w *Wilayah) Parse() {
	provinsi, kabupaten, kecamatan, kelurahan, tps := w.Provinsi.Kode, w.Kabupaten.Kode, w.Kecamatan.Kode, w.Kelurahan.Kode, w.TPS.Kode
	var ok bool
	if provinsi != "" {
		m := fetchProvinsi()
		w.Provinsi.Nama, ok = m[provinsi]
		if !ok {
			k, v := searchByValue(provinsi, m)
			if k == "" {
				fmt.Println("Kode provinsi tidak valid. Daftar provinsi:")
				debug.PrintJSON(m)
				os.Exit(1)

			}
			w.Provinsi.Kode, w.Provinsi.Nama = k, v
			provinsi = k
		}
		fmt.Printf("Provinsi\t: %s\n", w.Provinsi.Nama)
	}
	if kabupaten != "" {
		m := fetchKabupaten(provinsi)
		w.Kabupaten.Nama, ok = m[w.Kabupaten.Kode]
		if !ok {
			k, v := searchByValue(w.Kabupaten.Kode, m)
			if k == "" {
				fmt.Println("Kode kabupaten tidak valid. Daftar kabupaten:")
				debug.PrintJSON(m)
				os.Exit(1)
			}
			w.Kabupaten.Kode, w.Kabupaten.Nama = k, v
			kabupaten = k
		}
		fmt.Printf("Kabupaten\t: %s\n", w.Kabupaten.Nama)
	}
	if w.Kecamatan.Kode != "" {
		m := fetchKecamatan(provinsi, kabupaten)
		w.Kecamatan.Nama = m[w.Kecamatan.Kode]
		if !ok {
			k, v := searchByValue(kecamatan, m)
			if k == "" {
				fmt.Println("Kode kecamatan tidak valid. Daftar kecamatan:")
				debug.PrintJSON(m)
				os.Exit(1)
			}
			w.Kecamatan.Kode, w.Kecamatan.Nama = k, v
			kecamatan = k
		}
		fmt.Printf("Kecamatan\t: %s\n", w.Kecamatan.Nama)
	}
	if kelurahan != "" {
		m := fetchKelurahan(provinsi, kabupaten, kecamatan)
		w.Kelurahan.Nama = m[kelurahan]
		if !ok {
			k, v := searchByValue(kelurahan, m)
			if k == "" {
				fmt.Println("Kode kelurahan tidak valid. Daftar kelurahan:")
				debug.PrintJSON(m)
				os.Exit(1)
			}
			w.Kelurahan.Kode, w.Kelurahan.Nama = k, v
			kelurahan = k
		}
		fmt.Printf("Kelurahan\t: %s\n", w.Kelurahan.Nama)
	} else if tps != "" {
		w.TPS.Nama = fetchTps(provinsi, kabupaten, kecamatan, kelurahan)[tps]
		fmt.Printf("TPS: %s\n", w.TPS.Nama)
	}
}

type FormulirC1 struct {
	PerolehanSuara struct {
		Jokowi  int `json:"jokowi"`
		Prabowo int `json:"prabowo"`
	} `json:"perolehanSuara"`
	SuaraSah      int `json:"suaraSah"`
	SuaraTidakSah int `json:"suaraTidakSah"`
	SuaraTotal    int `json:"suaraTotal"`
	TotalDPT      int `json:"totalDPT"`
	TotalPengguna int `json:"totalPengguna"`
}

type Result struct {
	Provinsi  string      `json:"provinsi"`
	Kabupaten string      `json:"kabupaten"`
	Kecamatan string      `json:"kecamatan"`
	Kelurahan string      `json:"kelurahan"`
	TPS       string      `json:"tps"`
	Data      *FormulirC1 `json:"data"`
	Error     string      `json:"error"`
}

type Analytic struct {
	TotalTPS         int `json:"totalTPS"`
	TotalValid       int `json:"totalValid"`
	TotalBelumTerisi int `json:"totalBelumTerisi"`
	TotalAnomali     int `json:"totalAnomali"`
}
