.PHONY: build-run
build-run:
	go build -o ./build/api && ./build/api
