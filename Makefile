VERSION	?= $(shell git describe --abbrev=0 --tags)
LDFLAGS := -s -w -X squarecloud/internal/build.Version=$(VERSION)

build:
	go build -ldflags "$(LDFLAGS)" -o ./bin/squarecloud

clean:
	go clean
	rm *.snap