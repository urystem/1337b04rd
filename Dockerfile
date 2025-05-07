FROM golang:1.24.2-alpine3.21 AS builder

WORKDIR /app

COPY . .

RUN go build -o 1337b04rd cmd/main.go


FROM alpine

WORKDIR /app

COPY --from=builder /app/1337b04rd .

ENTRYPOINT [ "./1337b04rd" ]