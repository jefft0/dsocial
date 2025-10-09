FROM golang:1.24-alpine AS builder

COPY . /app

WORKDIR /app

RUN go build -o notification .


FROM alpine

COPY --from=builder /app/notification /usr/local/bin/notification

ENTRYPOINT [ "/usr/local/bin/notification" ]
