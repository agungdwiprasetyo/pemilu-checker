package main

import (
	"crypto/tls"
	"io/ioutil"
	"net/http"
	"time"
)

func requestData(url string) []byte {
	req, _ := http.NewRequest("GET", url, nil)

	client := new(http.Client)
	client.Transport = &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client.Timeout = 100 * time.Second
	resp, err := client.Do(req)
	if err != nil {
		return nil
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil
	}

	return body
}
