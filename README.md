# Ejemplo

# Version ejecutable en local con docker o en una sola maquina con docker

### Compilar proto

    protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative Proto/message.proto

### RabbitMQ Docker

    docker run -it --rm --name rabbitmq -p 5670:5672 -p 15670:15672 rabbitmq:3-management

### Programa principal en Docker
Ejecutar en dos consolas distintas.

Para la Central

    docker run -it --rm -P lab1:latest go run Central/main.go
Para el Laboratorio

    docker run -it --rm -P -p 50051:50051 lab1:latest go run Laboratorio/main.go
