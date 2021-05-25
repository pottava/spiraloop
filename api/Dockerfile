FROM golang:1.16.4-alpine3.13
ENV HOST=0.0.0.0 \
    PORT=80
RUN apk --no-cache add ca-certificates
COPY api /go/src/github.com/pottava/spiraloop/api
WORKDIR /go/src/github.com/pottava/spiraloop/api
RUN go mod download
RUN go mod verify
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /app \
    -ldflags "-s -w -X main.ver=${APP_VERSION} -X main.commit=$(git rev-parse --short HEAD 2>/dev/null) -X main.date=$(date +%Y-%m-%d --utc)"
RUN chmod +x /app
ENTRYPOINT ["/app"]
