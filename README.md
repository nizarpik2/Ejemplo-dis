# Ejemplo

# Version ejecutable en local con docker o en una sola maquina con docker

### Compilar proto

    protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative Proto/message.proto

### RabbitMQ Docker

    docker run -it --rm --name rabbitmq -p 5670:5672 -p 15670:15672 rabbitmq:3-management

#### Usuario (Puede que sea opcional, hacer solo si le da un error en el usuario y la contraseña)
Para crear el nuevo usuario.

    docker exec rabbitmq rabbitmqctl add_user 'test' 'test'
Para darle los permisos necesarios.

    docker exec rabbitmq rabbitmqctl set_permissions "test" "." "." ".*"

### docker

#### Maquina 1
Para levantar el contenedor

    docker build -t lab1 .
Para ejecutar el contenedor con la Central

    docker run -it --rm -P lab1:latest go run Central/main.go

#### Maquina 2
Para levantar el contenedor
    docker build -t lab1 .
Para ejecutar el contenedero con el Laboratorio
    docker run -it --rm -P -p 50051:50051 lab1:latest go run Laboratorio/main.go
