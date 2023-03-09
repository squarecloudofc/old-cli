VERSION	?= $(shell git describe --abbrev=0 --tags)
GOOS 	?= $(shell go env GOOS)
LDFLAGS := -s -w -X squarecloud/internal/build.Version=$(VERSION)
OUTPUT 	:= ./bin/squarecloud
MAIN 	:= ./cmd/squarecloud/squarecloud.go

EXE =
ifeq ($(GOOS),windows)
EXE = .exe
endif

.PHONY: build
build:
	go build -ldflags "$(LDFLAGS)" -o $(OUTPUT)$(EXE) $(MAIN)

.PHONY: test
test:
	go test -v ./...

.PHONY: clean
clean:
	go clean