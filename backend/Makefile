#@IgnoreInspection BashAddShebang
export ROOT=$(realpath $(dir $(lastword $(MAKEFILE_LIST))))
export DEBUG=true
export APP=backend
export LDFLAGS="-w -s"

all: build test

build:
	go build -race  .

build-static:
	CGO_ENABLED=0 go build -race -v -o $(APP) -a -installsuffix cgo -ldflags $(LDFLAGS) .

run:
	go run -race .

############################################################
# Test
############################################################

test:
	go test -v -race ./...

container:
	docker build -t $(APP) .

run-container:
	docker run --rm -it $(APP)

.PHONY: build run build-static test container
