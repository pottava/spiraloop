PWD := $(shell pwd)

.PHONY: all run start stop deps test build

all: build

run:
	@docker run --rm -t -v "$(PWD)":/go/src/github.com/pottava/spiraloop/api \
			--name spiraloop-api -d -p 8080:9000 \
			-w /go/src/github.com/pottava/spiraloop/api \
			-e HOST=0.0.0.0 -e PORT=9000 -e LOG_LEVEL=debug \
			golang:1.16.4-alpine3.13 go run main.go > /dev/null
	@docker logs -f spiraloop-api

stop:
	@docker stop spiraloop-api

codegen:
	@docker run --rm -t -v "$(PWD)":/go/src/github.com/pottava/spiraloop/api \
			-w /go/src/github.com/pottava/spiraloop/api \
			ghcr.io/supinf/go-swagger:0.25 \
			generate server -f spec.yaml -t generated

deps:
	@docker run --rm -t -v "$(PWD)":/go/src/github.com/pottava/spiraloop/api \
			-w /go/src/github.com/pottava/spiraloop/api \
			golang:1.16.4-alpine3.13 go mod vendor

test:
	@docker run --rm -t -v "$(PWD)":/go/src/github.com/pottava/spiraloop/api \
			-w /go/src/github.com/pottava/spiraloop/api \
			golangci/golangci-lint:v1.40.1-alpine \
			golangci-lint run --config .golangci.yml
	@docker run --rm -t -v "$(PWD)":/go/src/github.com/pottava/spiraloop/api \
			-w /go/src/github.com/pottava/spiraloop/api \
			--entrypoint go golang:1.16.4-alpine3.13 \
			test -vet off $(go list ./...)
