VERSION := v0.0.2

BIN_NAME = simpleQweather
IMAGE = woshiaotian/simple-qweather:$(VERSION)
LDFLAGS = -ldflags "-s -w"

.PHONY: build img clean
build:
	CGO_ENABLED=0 go build $(LDFLAGS) -o $(BIN_NAME)

image:
	docker build --rm -t $(IMAGE) -f Dockerfile .

