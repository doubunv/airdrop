FROM golang:1.22 AS builder

COPY . /src
WORKDIR /src

RUN go mod tidy
RUN go mod download

WORKDIR /src/cmd
RUN cd /src/cmd
RUN mkdir bin

COPY cmd/etc ./bin/etc
COPY cmd/swagger ./bin/swagger
RUN go build -o ./bin/main

FROM alpine:latest
COPY --from=builder /src/cmd/bin /app
WORKDIR /app
CMD ["./main", "-f", "etc/main.yaml"]
