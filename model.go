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
