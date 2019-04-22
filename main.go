package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"sync"

	"github.com/agungdwiprasetyo/go-utils"
	"github.com/agungdwiprasetyo/go-utils/debug"
)

const (
	url          = "https://pemilu2019.kpu.go.id/static/json/hhcw/ppwp/22328.json"
	jokowiAmin   = "21"
	prabowoSandi = "22"
)

var multiError *utils.MultiError
var m sync.Mutex

// flag.StringVar(&provinsi, "provinsi", "22328", "kode provinsi")
// 	flag.StringVar(&kabupaten, "kabupaten", "22875", "kode kabupaten")
// 	flag.StringVar(&kecamatan, "kecamatan", "22962", "kode kecamatan")
// 	flag.StringVar(&kelurahan, "kelurahan", "22964", "kode kelurahan")
// 	flag.StringVar(&tps, "tps", "900133331", "kode nomor tps")

func main() {
	defer func() {
		if r := recover(); r != nil {
			debug.Println(r)
			os.Exit(1)
		}
	}()

	var (
		provinsi, kabupaten, kecamatan, kelurahan, tps string
	)

	flag.StringVar(&provinsi, "provinsi", "", "kode provinsi")
	flag.StringVar(&kabupaten, "kabupaten", "", "kode kabupaten")
	flag.StringVar(&kecamatan, "kecamatan", "", "kode kecamatan")
	flag.StringVar(&kelurahan, "kelurahan", "", "kode kelurahan")
	flag.StringVar(&tps, "tps", "", "kode nomor tps")

	flag.Parse()

	var wilayah Wilayah
	wilayah.Provinsi.Kode = provinsi
	wilayah.Kabupaten.Kode = kabupaten
	wilayah.Kecamatan.Kode = kecamatan
	wilayah.Kelurahan.Kode = kelurahan
	wilayah.TPS.Kode = tps

	multiError = utils.NewMultiError()
	exec(wilayah)

	debug.PrintJSON(multiError.ToMap())
}

var (
	tt string
)

func exec(wilayah Wilayah) {
	provinsi, kabupaten, kecamatan, kelurahan, tps := wilayah.Provinsi.Kode, wilayah.Kabupaten.Kode, wilayah.Kecamatan.Kode, wilayah.Kelurahan.Kode, wilayah.TPS.Kode

	if provinsi == "" {
		log.Fatal("Provinsi tidak boleh kosong")
	} else if kabupaten == "" {
		kabs := fetchKabupaten(provinsi)
		for key, value := range kabs {
			wilayah.Kabupaten.Kode = key
			wilayah.Kabupaten.Nama = value
			exec(wilayah)
		}
	} else if kecamatan == "" {
		kecs := fetchKecamatan(provinsi, kabupaten)
		for key, value := range kecs {
			wilayah.Kecamatan.Kode = key
			wilayah.Kecamatan.Nama = value
			exec(wilayah)
		}
	} else if kelurahan == "" {
		kels := fetchKelurahan(provinsi, kabupaten, kecamatan)
		for key, value := range kels {
			wilayah.Kelurahan.Kode = key
			wilayah.Kelurahan.Nama = value
			exec(wilayah)
		}
	} else if tps == "" {
		listTps := fetchTps(provinsi, kabupaten, kecamatan, kelurahan)
		debug.PrintJSON(listTps)
		for key, value := range listTps {
			wilayah.TPS.Kode = key
			wilayah.TPS.Nama = value
			exec(wilayah)
		}
	} else {
		url := fmt.Sprintf("%s/%s/%s/%s/%s.json", provinsi, kabupaten, kecamatan, kelurahan, tps)
		data, err := detailTps(url)
		if err != nil {
			key := fmt.Sprintf("%s:%s:%s:%s:%s", provinsi, kabupaten, kecamatan, kelurahan, wilayah.TPS.Nama)
			multiError.Append(key, err)
			return
		}
		debug.PrintJSON(data)
	}
}
