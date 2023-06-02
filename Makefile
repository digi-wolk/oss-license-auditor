GIT_TAG?= $(shell git describe --always --tags)
BUILD_FLAGS := "-w -s -X 'main.Version=$(GIT_TAG)' -X 'main.GitTag=$(GIT_TAG)' -X 'main.BuildDate=$(BUILD_DATE)'"
BIN=olaudit
CGO_ENABLED = 0
GO_VERSION = 1.20

build-linux-amd64:
	CGO_ENABLED=$(CGO_ENABLED) GOOS=linux GOARCH=amd64 go build -ldflags=$(BUILD_FLAGS) -o build/$(BIN) ./cmd/olaudit/

build-linux-arm64:
	CGO_ENABLED=$(CGO_ENABLED) GOOS=linux GOARCH=arm64 go build -ldflags=$(BUILD_FLAGS) -o build/$(BIN) ./cmd/olaudit/
