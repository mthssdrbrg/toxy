build:
	@ mkdir -p build
	go build -o build/toxy ./cmd/toxy/...

docker-build:
	docker build -t mthssdrbrg/toxy .
