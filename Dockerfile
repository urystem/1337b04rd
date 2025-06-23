FROM golang:1.24.4-alpine3.22 AS builder

WORKDIR /app

COPY . .

RUN go build -o 1337b04rd cmd/main.go

# RUN go build -o 1337b04rd redis/1/1.go

FROM alpine

WORKDIR /app

COPY --from=builder /app/1337b04rd .
COPY --from=builder /app/web ./web

ENTRYPOINT [ "./1337b04rd" ]