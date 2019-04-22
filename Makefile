build:
	go build -o bin

run: build
	./bin --provinsi 22328 --kabupaten 22875 --kecamatan 22962 --kelurahan 22964 --tps 900133331