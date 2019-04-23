package main

import (
	"flag"
	"fmt"
	"os"
	"os/signal"
	"sync"

	"github.com/agungdwiprasetyo/go-utils"
	"github.com/agungdwiprasetyo/go-utils/debug"
)

const (
	jokowiAmin   = "21"
	prabowoSandi = "22"
)

var multiError *utils.MultiError
var result []*Result
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

	if provinsi != "" {
		wilayah.Provinsi.Nama = fetchProvinsi()[provinsi]
		debug.Println(wilayah.Provinsi.Nama)
	}
	if kabupaten != "" {
		wilayah.Kabupaten.Nama = fetchKabupaten(provinsi)[kabupaten]
		debug.Println(wilayah.Kabupaten.Nama)
	}
	if kecamatan != "" {
		wilayah.Kecamatan.Nama = fetchKecamatan(provinsi, kabupaten)[kecamatan]
		debug.Println(wilayah.Kecamatan.Nama)
	}
	if kelurahan != "" {
		wilayah.Kelurahan.Nama = fetchKelurahan(provinsi, kabupaten, kecamatan)[kelurahan]
		debug.Println(wilayah.Kelurahan.Nama)
	}
	if tps != "" {
		wilayah.TPS.Nama = fetchTps(provinsi, kabupaten, kecamatan, kelurahan)[tps]
		debug.Println(wilayah.TPS.Nama)
	}

	multiError = utils.NewMultiError()
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, os.Interrupt)

	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		defer wg.Done()

		exec(wilayah)
		// debug.PrintJSON(multiError.ToMap())
		debug.PrintJSON(result)
		os.Exit(0)
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()

		for {
			select {
			case <-sig:
				// debug.PrintJSON(multiError.ToMap())
				debug.PrintJSON(result)
				os.Exit(0)
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
			fmt.Printf("* Memproses data TPS %s\n", value)
			exec(wilayah)
		}
	} else {
		url := fmt.Sprintf("%s/%s/%s/%s/%s.json", provinsi, kabupaten, kecamatan, kelurahan, tps)
		data, err := detailTps(url)
		if err != nil {
			errStr := err.Error()
			fmt.Printf("\t%s: %v\n", wilayah.TPS.Nama, errStr)
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
			}
			// key := fmt.Sprintf("%s:%s:%s:%s:%s", provinsi, kabupaten, kecamatan, kelurahan, wilayah.TPS.Nama)
			// multiError.Append(key, err)
			return
		}
		debug.PrintJSON(data)
	}
}
