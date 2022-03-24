.PHONY: migrate-up migrate-down generate build

migrate-up:
	go run main.go migrate-up

migrate-down:
	go run main.go migrate-down

generate:
	go generate -v ./... && go get -v ./...

build:
	for dir in `find . -name main.go -type f`; do \
		go build -v -o "bin/$$(basename $$(dirname $$dir))" "$$(dirname $$dir)"; \
	done
