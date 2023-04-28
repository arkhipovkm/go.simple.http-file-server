FROM golang:1.18.5-alpine3.16 AS builder

WORKDIR /var/lib/app

COPY main.go .
ENV GO111MODULE=auto
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags '-s' -o main

FROM scratch

COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /var/lib/app/main /http-server

ENTRYPOINT [ "/http-server" ]