FROM golang:1.21.3-alpine3.18 as builder
RUN apk update && \
    apk upgrade && \
    apk add --no-cache bash git openssh make && \
    apk add --no-cache git ca-certificates tzdata && \
    cp /usr/share/zoneinfo/Asia/Jakarta /etc/localtime && \
    echo "Asia/Jakarta" > /etc/timezone

ADD .. /app/gandiwa

WORKDIR /app/gandiwa

RUN go mod tidy
RUN go mod download
RUN go mod vendor

RUN go clean -modcache

RUN make build

FROM alpine:3.16

RUN mkdir -p /app/gandiwa/resources

COPY --from=builder /app/gandiwa/resources/ /app/gandiwa/resources/
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt
COPY --from=builder /app/gandiwa/bin /app/gandiwa/
COPY --from=builder /app/gandiwa/logs /app/gandiwa/logs
COPY --from=builder /etc/localtime /etc/localtime
COPY --from=builder /etc/timezone /etc/timezone
COPY --from=builder /usr/share/zoneinfo /usr/share/zoneinfo

WORKDIR /app/gandiwa
EXPOSE 80
CMD ["./gandiwa"]
