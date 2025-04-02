FROM golang:1.22.7-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-s -w" -o /app/main ./main.go

FROM scratch

WORKDIR /app

COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt
ENV TZ Asia/Shanghai

COPY --from=builder /app/main /app/main
COPY templates /app/templates

EXPOSE 8100

CMD ["/app/main"]
