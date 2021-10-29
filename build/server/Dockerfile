FROM golang:alpine AS builder
RUN apk update && \
    apk add curl \
            git \
            bash \
            make \
            ca-certificates && \
    rm -rf /var/cache/apk/*

WORKDIR /build

COPY . .
RUN go mod download
RUN go mod verify

COPY . .
RUN make build

FROM alpine:latest AS certificates

RUN apk --no-cache add ca-certificates

FROM alpine:latest
WORKDIR /app
ENV PATH=/app/bin/:$PATH
COPY --from=certificates /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt
WORKDIR /app/
COPY --from=builder /build/server ./app/server
ENTRYPOINT ["server"]