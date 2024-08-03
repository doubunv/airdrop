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
RUN go mod tidy
RUN go build -o ./bin/main

WORKDIR /src/cmd/bin
CMD ["./main", "-f", "etc/mm-test.yaml"]

#docker cp boring_sinoussi:/src/cmd/bin ~/Desktop

#FROM alpine:latest
#COPY --from=builder /src/cmd/bin /app
#WORKDIR /app
#CMD ["./main", "-f", "etc/main.yaml"]


#docker run  --restart=unless-stopped -p 80:80 -p 443:443 --name nginx  --network boot_network -v /home/ec2-user/www/nginx/conf/nginx.conf:/etc/nginx/nginx.conf -v /home/ec2-user/www/nginx/conf/conf.d:/etc/nginx/conf.d -v /home/ec2-user/www/nginx/log:/var/log/nginx -v /home/ec2-user/www/html:/usr/share/nginx/html  -d nginx:latest