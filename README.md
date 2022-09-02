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


### Usuario

Para poder comunicar los laboratorios con RabbitMQ les recomiendo que vean la siguiente documentacion.

- https://www.rabbitmq.com/access-control.html

Pero en resumen, debe crear otro usuario que en el caso de este ejemplo es tanto el nombre como la contraseña son
"test".

    rabbitmqctl add_user 'test' 'test'

Luego le entregan todos los permisos.

    rabbitmqctl set_permissions "test" ".*" ".*" ".*"

