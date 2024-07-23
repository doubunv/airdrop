FROM golang:1.22 AS builder

COPY . /src
WORKDIR /src

RUN go mod tidy
RUN go mod download

RUN cd /src/cmd
RUN mkdir bin
RUN cd ..

WORKDIR /src/cmd

COPY cmd/etc ./bin/etc
COPY cmd/swagger ./bin/swagger
RUN go build -o ./bin/main

WORKDIR /src/cmd/bin
CMD ["./main", "-f", "etc/mm-test.yaml"]

#docker cp jolly_kirch:/src/cmd/bin/main ~/Desktop

#FROM alpine:latest
#COPY --from=builder /src/cmd/bin /app
#WORKDIR /app
#CMD ["./main", "-f", "etc/main.yaml"]
