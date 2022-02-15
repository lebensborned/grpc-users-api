FROM golang:latest
WORKDIR /app
ADD . .
RUN go mod download
CMD go run cmd/server/main.go