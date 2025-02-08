.PHONY: build run clean

build:
	go build -o bin/app cmd/app/main.go

run: build
	./bin/app

clean:
	rm -rf bin/
