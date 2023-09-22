BINARY_NAME=build
.DEFAULT_GOAL := run

build: clean
	@GOARCH=arm64 GOOS=darwin go build -o bin/${BINARY_NAME}/darwin
	@GOARCH=amd64 GOOS=linux go build -o bin/${BINARY_NAME}/linux
	@GOARCH=amd64 GOOS=windows go build -o bin/${BINARY_NAME}/windows

run: build
	@./bin/build

clean:
	@rm -rf ./bin/build

test: 
	go test -v ./... -count=1

test_coverage:
	go test ./... -coverprofile=coverage.out

dep:
	go mod download
