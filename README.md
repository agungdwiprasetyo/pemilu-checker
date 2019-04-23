# Pemilu Data Checker

## A HTTP tools for checking result of Indonesian Election 2019 Recursively

### Use
* Install Golang & dependencies
```sh
$ brew install golang
$ glide install
```

* Build tools
```sh
$ make build
```

* Run tools
```sh
$ ./bin --provinsi [kode provinsi] --kabupaten [kode kabupaten] --kecamatan [kode kecamatan] --kelurahan [kode kelurahan] --tps [kode tps]
```

### Flag
* ```--provinsi``` : kode provinsi, if this flag is empty program will run in all TPS in all province in Indonesia
* ```--kabupaten``` : kode kabupaten, if this flag is empty program will run in all kabupaten in given provinsi recursively
* ```--kecamatan``` : kode kecamatan, if this flag is empty program will run in all kecamatan in given kabupaten recursively
* ```--kelurahan``` : kode kelurahan, if this flag is empty program will run in all kelurahan in given kecamatan recursively
* ```--tps``` : kode tps, if this flag is empty program will run in all tps in given kelurahan recursively

### Example
Sample data (last update 2019-04-23 15:39:00):
* Provinsi: **Aceh** (kode: `1`)
* Kabupaten: **Aceh Timur** (kode: `671`)
* Kecamatan: **Julok** (kode: `718`)
* Kelurahan: **Blang Pauh Dua** (kode: `739`)

Command:
```sh
$ ./pemilu-checker --provinsi 1 --kabupaten 671 --kecamatan 718 --kelurahan 739
```

**Output:**

List TPS:
```json
{
     "900001424": "TPS 1",
     "900001425": "TPS 2",
     "900001426": "TPS 3"
}
```

Error:
```json
[
     {
          "provinsi": "ACEH",
          "kabupaten": "ACEH TIMUR",
          "kecamatan": "JULOK",
          "kelurahan": "BLANG PAUH DUA",
          "tps": "TPS 3",
          "data": {
               "perolehanSuara": {
                    "jokowi": 16,
                    "prabowo": 14
               },
               "suaraSah": 157,
               "suaraTidakSah": 2,
               "suaraTotal": 159,
               "totalDPT": 227,
               "totalPengguna": 159
          },
          "error": "SALAH. 16+14 bukan 157, bisa ngitungnya gak sih"
     }
]
```

### Download
Download binary for Linux:
```sh
$ wget https://storage.googleapis.com/agungdp/bin/pemilu-checker/bin_linux && chmod 777 bin_linux
```

Download binary for MacOS:
```sh
$ wget https://storage.googleapis.com/agungdp/bin/pemilu-checker/bin_mac_os && chmod 777 bin_mac_os
```

### TODO
Implement OpenCV for image processing from C1 form photo