PWD := $(shell pwd)

.PHONY: all run start stop deps test build

all: build

start:
	@docker run --rm -t -v "$(PWD)":/go/src/github.com/pottava/spiraloop/cli \
			-w /go/src/github.com/pottava/spiraloop/cli \
			-e APP_DEBUG=0 \
			-e API_ENDPOINT=localhost:8080 \
			golang:1.16.4-alpine3.13 \
			go run cmd/spiraloop/main.go \
			start --key foo

success:
	@docker run --rm -t -v "$(PWD)":/go/src/github.com/pottava/spiraloop/cli \
			-w /go/src/github.com/pottava/spiraloop/cli \
			-e APP_DEBUG=0 \
			-e API_ENDPOINT=localhost:8080 \
			golang:1.16.4-alpine3.13 \
			go run cmd/spiraloop/main.go \
			success --key foo

deps:
	@docker run --rm -t -v "$(PWD)":/go/src/github.com/pottava/spiraloop/cli \
			-w /go/src/github.com/pottava/spiraloop/cli \
			golang:1.16.4-alpine3.13 go mod vendor

test:
	@docker run --rm -t -v "$(PWD)":/go/src/github.com/pottava/spiraloop/cli \
			-w /go/src/github.com/pottava/spiraloop/cli \
			golangci/golangci-lint:v1.40.1-alpine \
			golangci-lint run --config .golangci.yml
	@docker run --rm -t -v "$(PWD)":/go/src/github.com/pottava/spiraloop/cli \
			-w /go/src/github.com/pottava/spiraloop/cli \
			--entrypoint go golang:1.16.4-alpine3.13 \
			test -vet off $(go list ./...)
