all: build
TAG = v0.1.0
PREFIX = onemorehill

deps:
	go get github.com/tools/godep

clean:
	rm -f kubeme-toolkit

build: clean deps 
	GOOS=linux GOARCH=amd64 GCO_ENABLED=0 godep go build -a .

container: build
		docker build -t $(PREFIX)/kubeme-toolkit:$(TAG) .

.PHONY: all deps build clean container 
