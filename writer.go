package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/agungdwiprasetyo/go-utils/debug"
)

func writeToFile(filename string) {
	if result == nil {
		fmt.Println("Proses selesai. Tidak ada anomali data")
	} else {
		b, _ := json.MarshalIndent(result, "", "\t")
		ioutil.WriteFile(filename, b, 0644)
		fmt.Printf("Proses selesai. Terdapat %d anomali data\n", len(result))
	}

	fmt.Printf("Summary:\n")
	debug.PrintJSON(analytic)
	os.Exit(0)
}
