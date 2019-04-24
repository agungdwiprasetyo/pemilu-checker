package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

func writeToFile(filename string) {
	if result == nil {
		fmt.Println("Proses selesai. Tidak ada anomali data")
	} else {
		b, _ := json.MarshalIndent(result, "", "\t")
		ioutil.WriteFile(filename, b, 0644)
	}
	os.Exit(0)
}
