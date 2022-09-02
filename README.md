# Ejemplo

## Version Ejecutable en local o en una sola maquina

### Compilar proto
    protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative Proto/message.proto

### Docker

    docker run -it --rm --name rabbitmq -p 5672:5672 -p 15672:15672 rabbitmq:3-management

### GO

Ejecutar cada uno en una consola distinta.

Para la central.
    
    go run Central/main.go

Para el Laboratorio

    go run Laboratorio/main.go
