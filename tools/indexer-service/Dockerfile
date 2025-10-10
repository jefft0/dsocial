FROM golang:1.24-alpine AS builder

COPY . /app

WORKDIR /app

RUN go build -o indexer .


FROM alpine

COPY --from=builder /app/indexer /usr/local/bin/indexer

ENTRYPOINT [ "/usr/local/bin/indexer" ]
