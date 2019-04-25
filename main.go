package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"os/signal"
	"sync"
	"time"

	"github.com/agungdwiprasetyo/go-utils"
	"github.com/agungdwiprasetyo/go-utils/debug"
)

const (
	jokowiAmin   = "21"
	prabowoSandi = "22"
)

var multiError *utils.MultiError
var analytic Analytic
var result []*Result
var m sync.Mutex

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

	if _, err := ioutil.ReadDir("results"); err != nil {
		os.Mkdir("results", 0700)
	}

	var wilayah Wilayah
	wilayah.Provinsi.Kode = provinsi
	wilayah.Kabupaten.Kode = kabupaten
	wilayah.Kecamatan.Kode = kecamatan
	wilayah.Kelurahan.Kode = kelurahan
	wilayah.TPS.Kode = tps

	wilayah.Parse()

	multiError = utils.NewMultiError()
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, os.Interrupt)

	filename := fmt.Sprintf("%d.json", time.Now().Unix())

	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		defer wg.Done()

		exec(wilayah)
		writeToFile(filename)
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()

		for {
			select {
			case <-sig:
				writeToFile(filename)
			}
		}
	}()

	wg.Wait()
}

func exec(wilayah Wilayah) {
	provinsi, kabupaten, kecamatan, kelurahan, tps := wilayah.Provinsi.Kode, wilayah.Kabupaten.Kode, wilayah.Kecamatan.Kode, wilayah.Kelurahan.Kode, wilayah.TPS.Kode

	if provinsi == "" {
		provs := fetchProvinsi()
		debug.PrintJSON(provs)
		for key, value := range provs {
			wilayah.Provinsi.Kode = key
			wilayah.Provinsi.Nama = value
			fmt.Printf("* Memproses data Provinsi %s\n", value)
			exec(wilayah)
		}
	} else if kabupaten == "" {
		kabs := fetchKabupaten(provinsi)
		for key, value := range kabs {
			wilayah.Kabupaten.Kode = key
			wilayah.Kabupaten.Nama = value
			fmt.Printf("* Memproses data Kabupaten %s\n", value)
			exec(wilayah)
		}
	} else if kecamatan == "" {
		kecs := fetchKecamatan(provinsi, kabupaten)
		for key, value := range kecs {
			wilayah.Kecamatan.Kode = key
			wilayah.Kecamatan.Nama = value
			fmt.Printf("* Memproses data Kecamatan %s\n", value)
			exec(wilayah)
		}
	} else if kelurahan == "" {
		kels := fetchKelurahan(provinsi, kabupaten, kecamatan)
		for key, value := range kels {
			wilayah.Kelurahan.Kode = key
			wilayah.Kelurahan.Nama = value
			fmt.Printf("* Memproses data Kelurahan %s\n", value)
			exec(wilayah)
		}
	} else if tps == "" {
		listTps := fetchTps(provinsi, kabupaten, kecamatan, kelurahan)
		debug.PrintJSON(listTps)
		for key, value := range listTps {
			wilayah.TPS.Kode = key
			wilayah.TPS.Nama = value
			fmt.Printf("* Memproses data [%s => %s => %s => %s] TPS: %s\n",
				wilayah.Provinsi.Nama, wilayah.Kabupaten.Nama, wilayah.Kecamatan.Nama, wilayah.Kelurahan.Nama, value)
			exec(wilayah)
		}
	} else {
		analytic.TotalTPS++
		url := fmt.Sprintf("%s/%s/%s/%s/%s.json", provinsi, kabupaten, kecamatan, kelurahan, tps)
		data, err := detailTps(url)
		debug.PrintJSON(data)
		if err != nil {
			errStr := err.Error()
			fmt.Printf("\t%s: %v\n\n", wilayah.TPS.Nama, errStr)
			if errStr != NotFound {
				result = append(result, &Result{
					Provinsi:  wilayah.Provinsi.Nama,
					Kabupaten: wilayah.Kabupaten.Nama,
					Kecamatan: wilayah.Kecamatan.Nama,
					Kelurahan: wilayah.Kelurahan.Nama,
					TPS:       wilayah.TPS.Nama,
					Data:      data,
					Error:     errStr,
				})
				analytic.TotalAnomali++
			} else {
				analytic.TotalBelumTerisi++
			}
			// key := fmt.Sprintf("%s:%s:%s:%s:%s", provinsi, kabupaten, kecamatan, kelurahan, wilayah.TPS.Nama)
			// multiError.Append(key, err)
		} else {
			analytic.TotalValid++
		}
	}
}
