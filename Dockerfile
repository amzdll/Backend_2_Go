FROM golang:1.22.3-alpine3.20 AS builder

COPY . /shop
WORKDIR /shop

RUN go mod download
RUN go build -o ./bin/main cmd/main.go

#FROM alpine:3.20
#
#WORKDIR /root/
#COPY --from=builder /shop/bin/app .
#CMD ["./app"]