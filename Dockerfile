FROM golang:1.22 AS builder

COPY . /src
WORKDIR /src

RUN go mod tidy
RUN go mod download && go build -o bin/main
COPY ./etc bin/etc
COPY ./swagger bin/swagger

FROM alpine:latest
COPY --from=builder /www/bin /app
WORKDIR /app
CMD ["./main", "-f", "etc/main.yaml"]

#
#RUN apt-get update && apt-get install curl -y && apt-get install -y --no-install-recommends \
#		ca-certificates  \
#        netbase \
#        && rm -rf /var/lib/apt/lists/

#COPY --from=builder /src/bin /app
#WORKDIR /app
#
#EXPOSE 8080
#VOLUME /data/conf

#ENTRYPOINT ["/app/docker-entrypoint.sh"]

#CMD ["./listener-service", "-f", "/data/conf/mm-test.yaml"]
