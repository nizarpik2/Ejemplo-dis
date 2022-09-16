package main

import (
	"fmt"
	"log"
	"context"
	"time"
	//"io/ioutil"
	"os"
	amqp "github.com/rabbitmq/amqp091-go"
	"google.golang.org/grpc"
	pb "github.com/Kendovvul/Ejemplo/Proto"
)

func central (squad string) {
	qName := "Emergencias" //Nombre de la cola
	hostQ := "172.17.0.1"  //Host de RabbitMQ 172.17.0.1
	connQ, err := amqp.Dial("amqp://guest:guest@"+hostQ+":5670") //Conexion con RabbitMQ

	if err != nil {log.Fatal(err)}
	defer connQ.Close()

	ch, err := connQ.Channel()
	if err != nil{log.Fatal(err)}
	defer ch.Close()

	q, err := ch.QueueDeclare(qName, false, false, false, false, nil) //Se crea la cola en RabbitMQ
	if err != nil {log.Fatal(err)}

	fmt.Println(q)

	fmt.Println("Esperando Emergencias")
	chDelivery, err := ch.Consume(qName, "", false, false, false, false, nil) //obtiene la cola de RabbitMQ AutoACK fols 3er parametro
	if err != nil {
		log.Fatal(err)
	}

	for delivery := range chDelivery {
		if (1 == 2) {
			fmt.Println("No hay equipos disponibles!")
			/*
			for cant == 0{
				time.Sleep(5 * time.Second)
			}
			*/
		} else{
			port := "0000"
			hostS := "default"
			// Puerto de la conexion con el laboratorio
			if string(delivery.Body) == "Laboratorio Pripyat"{
				hostS = "dist097" //Host de un Laboratorio
				port = ":50051"
			}
			if string(delivery.Body) == "Laboratorio Kampala"{
				hostS = "dist098" //Host de un Laboratorio
				port = ":50052"
			}
			if string(delivery.Body) == "Laboratorio Renca"{
				hostS = "dist099" //Host de un Laboratorio
				port = ":50053"
			}
			if string(delivery.Body) == "Laboratorio Pohang"{
				hostS = "dist100" //Host de un Laboratorio
				port = ":50054"
			}
			// Obtiene el primer mensaje de la cola
			fmt.Println("Mensaje asíncrono de " + string(delivery.Body) + " leído")
			// Crea la conexion sincrona con el laboratorio
			connS, err := grpc.Dial(hostS + port, grpc.WithInsecure())

			if err != nil {
				panic("No se pudo conectar con el servidor" + err.Error())
			}
			
			if err != nil {
				panic("No se puede crear el mensaje " + err.Error())
			}
		
			//defer connS.Close()
		
			serviceCliente := pb.NewMessageServiceClient(connS)

			serviceCliente.Intercambio(context.Background(), &pb.Message{Body: squad,})

			fmt.Println("Se envía escuadra " + squad + " a " + string(delivery.Body) + ".")
			
			var consultas int = 0

			// Ciclo de contención de amenaza
			for {
				//espera de 5 segundos
				time.Sleep(5 * time.Second)
				//envia el mensaje al laboratorio
				res, err := serviceCliente.Intercambio(context.Background(), 
					&pb.Message{
						Body: "Equipo listo?",
					})

					consultas += 1

				if err != nil {
					panic("No se puede crear el mensaje " + err.Error())
				}
				response := res.Body
				fmt.Println("Status " + squad + ": " + response)
				if response == "SI"{
					serviceCliente.Intercambio(context.Background(), &pb.Message{Body: "STOP MENACE",})
					var escrito string
					escrito = string(delivery.Body) + "; " + string(consultas) + "\n" //formato (NombreLab;CantidadDeConsultas)
					f, err := os.OpenFile("SOLICITUDES.txt",
					os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
					if err != nil {
						log.Println(err)
					}
					defer f.Close()
					if _, err := f.WriteString([]byte(escrito)); err != nil {
						log.Println(err)
					}
					break
				}
			}
			connS.Close()
			delivery.Ack(false) //ACK cuando se resuelve la amenaza
			fmt.Println("Retorno a central " + squad + ", conexión " + string(delivery.Body) + " cerrada.") //dummy out for lab name
		}
		time.Sleep(1 * time.Second)
	}

}

func main(){
	go central("SQUAD A")
	central("SQUAD B")
}