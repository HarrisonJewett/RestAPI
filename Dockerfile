FROM golang:1.16.3-alpine as builder

WORKDIR /app/src/

COPY go.mod go.mod
COPY go.sum go.sum

RUN go mod download

COPY mocks mocks
COPY db db
COPY types types
COPY usecase usecase
COPY server server
COPY main.go main.go

RUN go build -o /server

FROM alpine:3.13

COPY --from=builder /server /server

ENTRYPOINT [ "/server" ]
