package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"time"
)

func writeToFile(filename string) {
	filename = fmt.Sprintf("results/%s", filename)
	var res struct {
		LastChecked string    `json:"lastChecked"`
		Summary     *Analytic `json:"summary"`
		Anomaly     []*Result `json:"anomaly"`
	}
	res.LastChecked = time.Now().Format(time.RFC3339)
	res.Summary = &analytic

	fmt.Printf("\n--------------------------------\n")
	if result == nil {
		fmt.Println("Proses selesai. \x1b[32;1mTidak ada anomali data\x1b[0m")
	} else {
		res.Anomaly = result
		fmt.Printf("Proses selesai. \x1b[31;1mTerdapat %d anomali data\x1b[0m\n", len(result))
	}

	b, _ := json.MarshalIndent(res, "", "\t")
	if err := ioutil.WriteFile(filename, b, 0644); err != nil {
		log.Fatal("Gagal menyimpan ke file")
	}

	fmt.Printf("Hasil tersimpan dalam file => \x1b[33;1m%s\x1b[0m\n\n", filename)
	os.Exit(0)
}
