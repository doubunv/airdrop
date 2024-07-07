FROM golang:1.21 AS builder

COPY . /src
WORKDIR /src

RUN make build

FROM debian:bullseye

RUN apt-get update && apt-get install curl -y && apt-get install -y --no-install-recommends \
		ca-certificates  \
        netbase \
        && rm -rf /var/lib/apt/lists/

COPY --from=builder /src/bin /app
WORKDIR /app

EXPOSE 8888
EXPOSE 8888
VOLUME /data/conf

#ENTRYPOINT ["/app/docker-entrypoint.sh"]

CMD ["./listener-service", "-f", "/data/conf/mm-test.yaml"]
