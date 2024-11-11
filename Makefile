VERSION := v0.0.1

BIN_NAME = simpleQweather
CONTAINER = woshiaotian/simpleQweather
LDFLAGS = -ldflags "-s -w"

.PHONY: build img clean
build:
	CGO_ENABLED=0 go build $(LDFLAGS) -o $(BIN_NAME)

img:
	docker build --rm -t $(CONTAINER) -f Dockerfile .

