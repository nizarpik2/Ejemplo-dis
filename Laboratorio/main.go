// commit simple

package main

import (
	"fmt"
	"math/rand" // consultar
	"log"
	"context"
	"net"
	amqp "github.com/rabbitmq/amqp091-go"
	"google.golang.org/grpc"
	pb "github.com/Kendovvul/Ejemplo/Proto"
)

type server struct {
	pb.UnimplementedMessageServiceServer
}

func (s *server) Intercambio (ctx context.Context, msg *pb.Message) (*pb.Message, error){
	fmt.Println(msg.Body)
	fmt.Println("hola")
	return &pb.Message{Body: "SI",}, nil
}

func main() {
	LabName := "Laboratiorio Pripyat" //nombre del laboratorio
	qName := "Emergencias" //nombre de la cola
	hostQ := "dist097" //ip del servidor de RabbitMQ 172.17.0.1
	connQ, err := amqp.Dial("amqp://guest:guest@"+hostQ+":5670") //conexion con RabbitMQ
	
	if err != nil {log.Fatal(err)}
	defer connQ.Close()

	ch, err := connQ.Channel()
	if err != nil{log.Fatal(err)}
	defer ch.Close()

	//Mensaje enviado a la cola de RabbitMQ (Llamado de emergencia)
	err = ch.Publish("", qName, false, false,
		amqp.Publishing{
			Headers: nil,
			ContentType: "text/plain",
			Body: []byte(LabName),  //Contenido del mensaje
		})

	if err != nil{
		log.Fatal(err)
	}

	fmt.Println(LabName)
	port := ":50051"
	listener, err := net.Listen("tcp", port) //conexion sincrona
	if err != nil {
		panic("La conexion no se pudo crear" + err.Error())
	}

	serv := grpc.NewServer()
	for {
		// must dummy out
		prob := rand.Intn(6)
		fmt.Println(prob)
		if (prob == 1 || prob == 2) {
			fmt.Println("NO")
		} else{
			pb.RegisterMessageServiceServer(serv, &server{})
			if err = serv.Serve(listener); err != nil {
				panic("El server no se pudo iniciar" + err.Error())
			}
		}
	}
}
