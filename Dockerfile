# build stage
FROM golang:1.13 as builder

WORKDIR /app

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ./bin/server

# final stage
FROM alpine:3.9
COPY --from=builder /app /app

RUN apk update && \
    apk add ca-certificates && \
    rm -rf /var/cache/apk/*

WORKDIR /app

EXPOSE 7000

ENTRYPOINT ["/app/bin/server"]