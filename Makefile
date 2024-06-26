
test:
	go test -v ./...
	go test -v ./modules
	go test -v ./modules/v2
	go test -v ./strings
	go test -v ./validator

build:
	go build -o build/bin/concurrency concurrency/cmd/main.go

clean:
	rm -rf build/