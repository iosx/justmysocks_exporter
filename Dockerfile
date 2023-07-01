FROM golang:1.20.5-bullseye

ENV GOPROXY=https://goproxy.cn,direct

WORKDIR /app

RUN apt-get update && apt-get install -y git && git clone https://github.com/wulabing/justmysocks_exporter.git .

RUN go mod download && go build -v -o justmysocks_exporter . && chmod +x justmysocks_exporter
FROM debian:bullseye-slim

WORKDIR /app

COPY --from=0 /app/justmysocks_exporter .

RUN apt-get update && apt-get install -y ca-certificates

EXPOSE 10001

ENV API_ADDRESS="https://justmysocks5.net/members/getbwcounter.php" \
    SERVICE="" \
    ID="" \
    UPDATE_INTERVAL="5m"


# docker run -p 10001:10001 -e SERVICE="your-service-number" -e ID="your-uuid" your-image-name
#CMD ["./justmysocks_exporter", \
#     "-api-address=${API_ADDRESS}", \
#     "-service=${SERVICE}", \
#     "-id=${ID}", \
#     "-update-interval=${UPDATE_INTERVAL}"]
CMD ./justmysocks_exporter \
     -api-address=$API_ADDRESS \
     -service=$SERVICE \
     -id=$ID \
     -update-interval=$UPDATE_INTERVAL