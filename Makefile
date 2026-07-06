.PHONY: build install test clean

VERSION := 0.2.0
BINARY  := blink

build:
	go build -ldflags "-X main.Version=$(VERSION)" -o $(BINARY) ./cmd/blink

install:
	go install -ldflags "-X main.Version=$(VERSION)" ./cmd/blink

test:
	go test ./...

clean:
	rm -f $(BINARY)
