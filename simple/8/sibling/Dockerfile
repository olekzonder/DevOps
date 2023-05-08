## Build

FROM golang:1.17-alpine as builder

RUN apk update

WORKDIR /app

COPY go.mod .
COPY go.sum .
RUN go mod download 

COPY . .

RUN go build -o build-increment

FROM alpine:3.17 as main
WORKDIR /app/
COPY --from=builder /app/build-increment .
ENTRYPOINT ["/app/build-increment"]
