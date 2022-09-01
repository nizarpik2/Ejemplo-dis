# Ejemplo

## Version Ejecutable en local o en una sola maquina

### Compilar proto
    protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative Proto/message.proto

### GO

Ejecutar cada uno en una consola distinta.

    go run Central/main.go \\
    go run Laboratorio/main.go
