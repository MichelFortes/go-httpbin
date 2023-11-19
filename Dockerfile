FROM golang:1.21-alpine AS builder
WORKDIR /app
COPY go.mod server.go ./
RUN go build -o go-httpbin

FROM scratch
WORKDIR /app
COPY --from=builder /app/go-httpbin .
ENTRYPOINT [ "/app/go-httpbin" ]
EXPOSE 80