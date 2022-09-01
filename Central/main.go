package main

import (
	"fmt"
	"log"
	"context"
	"time"
	amqp "github.com/rabbitmq/amqp091-go"
	"google.golang.org/grpc"
	pb "github.com/Kendovvul/Ejemplo/Proto"
)

func main () {
	qName := "Emergencias"
	hostQ := "localhost"
	hostS := "localhost"
	conn, err := amqp.Dial("amqp://guest:guest@"+hostQ+":5672") //Conexion con RabbitMQ

	if err != nil {log.Fatal(err)}
	defer conn.Close()

	ch, err := conn.Channel()
	if err != nil{log.Fatal(err)}
	defer ch.Close()

	fmt.Println("Esperando Emergencias")
	chDelivery, err := ch.Consume(qName, "", true, false, false, false, nil) //obtiene la cola de RabbitMQ
	if err != nil {
		log.Fatal(err)
	}
	
	for delivery := range chDelivery {
		
		fmt.Println("Pedido de ayuda de " + string(delivery.Body)) //obtiene el primer mensaje de la cola
		connS, err := grpc.Dial(hostS + ":50051", grpc.WithInsecure()) //crea la conexion sincrona con el laboratorio

		if err != nil {
			panic("No se pudo conectar con el servidor" + err.Error())
		}
	
		defer connS.Close()
	
		serviceCliente := pb.NewMessageServiceClient(connS)
	
		for {
			//envia el mensaje al laboratorio
			res, err := serviceCliente.Intercambio(context.Background(), 
				&pb.Message{
					Body: "Equipo listo?",
				})
	
			if err != nil {
				panic("No se puede crear el mensaje " + err.Error())
			}

			fmt.Println(res.Body) //respuesta del laboratorio
			time.Sleep(5 * time.Second) //espera de 5 segundos
		}
	}

}