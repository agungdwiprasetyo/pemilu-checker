package main

import (
	"crypto/tls"
	"io/ioutil"
	"net/http"
	"strings"
	"time"

	"github.com/agungdwiprasetyo/go-utils/debug"
)

func requestData(url string) []byte {
	req, _ := http.NewRequest("GET", url, nil)

	client := new(http.Client)
	client.Transport = &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client.Timeout = 3 * time.Second
	resp, err := client.Do(req)
	if err != nil {
		debug.Println(err)
		if strings.Contains(err.Error(), Timeout) || strings.Contains(err.Error(), Reset) {
			return requestData(url)
		}
		return nil
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil
	}

	return body
}
