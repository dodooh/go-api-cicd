FROM golang:1.23.1 AS builder

WORKDIR /app

COPY go.mod ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -mod=mod -o myapi

FROM alpine:latest

RUN apk --no-cache add ca-certificates
WORKDIR /app
COPY --from=builder /app/myapi .

EXPOSE 8081
CMD ["./myapi"]