# gRPC based API users 


# Stack

- Golang
- Apache Kafka
- Docker, docker-compose
- Redis
- Clickhouse DB
- postgreSQL
- gRPC with .proto


# Usage

docker-compose pull && docker-compose up -d

Environment variables can be edited in app.env


# Additional information

It is assumed that you will use a custom gRPC client, since only the server is implemented in this repository. However, there is a small client for debugging, located at: cmd/client/main.go
