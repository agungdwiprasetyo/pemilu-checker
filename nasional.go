package main

func fetchNasional() {
	// req, _ := http.NewRequest("GET", url, nil)

	// client := new(http.Client)
	// client.Transport = &http.Transport{
	// 	TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	// }
	// resp, err := client.Do(req)
	// if err != nil {
	// 	panic(err)
	// }
	// defer resp.Body.Close()

	// body, err := ioutil.ReadAll(resp.Body)
	// if err != nil {
	// 	panic(err)
	// }

	// var data struct {
	// 	Table map[string]interface{} `json:"table"`
	// }
	// json.Unmarshal(body, &data)

	// jokowi, prabowo := 0, 0
	// var kodeWilayah []string
	// for kode, val := range data.Table {
	// 	v := val.(map[string]interface{})
	// 	jokowi += int(v[jokowiAmin].(float64))
	// 	prabowo += int(v[prabowoSandi].(float64))
	// 	kodeWilayah = append(kodeWilayah, kode)
	// }

	// total := jokowi + prabowo
	// debug.Println("01. Jokowi:", jokowi, float64(float64(jokowi)/float64(total))*100)
	// debug.Println("02. Prabowo:", prabowo, float64(float64(prabowo)/float64(total))*100)

	// debug.PrintJSON(data)
	// debug.PrintJSON(fetchWilayah())
}
