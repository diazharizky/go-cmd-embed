migrate-up:
	go run main.go migrate-up

migrate-down:
	go run main.go migrate-down

generate:
	go generate -v ./...

build:
	for dir in `find . -name main.go -type f`; do \
		go build -v -o "bin/$$(basename $$(dirname $$dir))" "$$(dirname $$dir)"; \
	done

clear-build:
	rm -rf bin/*
