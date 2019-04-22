package main

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

type FormulirC1 struct {
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
