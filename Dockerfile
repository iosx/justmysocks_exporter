FROM golang:1.20.5-bullseye

ENV GOPROXY=https://goproxy.cn,direct

COPY . /app
WORKDIR /app

RUN go mod download && go build

FROM debian:bullseye-slim

WORKDIR /app

COPY --from=0 /app/justmysocks_exporter .

RUN  apt-get update \
 &&  apt-get install -qy ca-certificates \
 &&  rm -rf /var/lib/apt/lists/* /tmp/* /var/tmp/*

EXPOSE 10001

ENV API_ADDRESS="https://justmysocks6.net/members/getbwcounter.php"
ENV SERVICE=""
ENV ID="" 

# docker run -p 10001:10001 -e SERVICE="your-service-number" -e ID="your-uuid" your-image-name
ENTRYPOINT ["/bin/sh", "-c", "./justmysocks_exporter -api-address=$API_ADDRESS -service=$SERVICE -id=$ID"]