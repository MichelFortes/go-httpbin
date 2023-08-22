FROM golang:1.21-alpine
WORKDIR /app
COPY go.mod server.go ./
RUN go build -o go-httpbin
CMD [ "/app/go-httpbin" ]