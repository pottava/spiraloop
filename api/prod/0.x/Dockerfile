# Spiraloop v0.2
# docker run --rm -it -p 8080:80 spiraloop

FROM golang:1.16.4-alpine3.13 AS builder
RUN apk --no-cache add gcc musl-dev git
WORKDIR /go/src/github.com/pottava
RUN git clone https://github.com/pottava/spiraloop.git
WORKDIR /go/src/github.com/pottava/spiraloop
ENV APP_VERSION=0.x
# RUN git checkout "v${APP_VERSION}" > /dev/null 2>&1
WORKDIR /go/src/github.com/pottava/spiraloop/api
RUN mkdir -p generated
COPY --from=ghcr.io/supinf/go-swagger:0.25 /usr/bin/swagger /usr/bin/swagger
RUN swagger generate server -f spec.yaml -t generated
RUN go mod download
RUN go mod verify
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /app \
    -ldflags "-s -w -X main.ver=${APP_VERSION} -X main.commit=$(git rev-parse --short HEAD 2>/dev/null) -X main.date=$(date +%Y-%m-%d --utc)"
RUN chmod +x /app

FROM alpine:3.13 AS libs
RUN apk --no-cache add ca-certificates

FROM scratch
COPY --from=libs /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
ENV HOST=0.0.0.0 \
    PORT=80
COPY --from=builder /app /spiraloop
ENTRYPOINT ["/spiraloop"]
