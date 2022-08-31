package main

import (
	"fmt"
	"log"
	amqp "github.com/rabbitmq/amqp091-go"
	"google.golang.org/grpc"
	pb "github.com/Sistemas-Distribuidos-2022-2/Ejemplo/Proto"
)

func main() {
	LabName := "Laboratiorio Pripyat"
	qName := "Emergencias"
	hostQ := "localhost"
	connQ, err := amqp.Dial("amqp://guest:guest@"+host+":5672")
	
	if err != nil {
		log.Fatal(err)
	}

	defer connQ.Close()

	ch, err := connQ.Channel()

	if err != nil{
		log.Fatal(err)
	}
	defer ch.Close()
	q, err := ch.QueueDeclare(qName, false, false, false, false, nil)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(q)

	err = ch.Publish("", qName, false, false,
		amqp.Publishing{
			Headers: nil,
			ContentType: "text/plain",
			Body: []byte(LabName),
		})

	if err != nil{
		log.Fatal(err)
	}

	fmt.Println(LabName)

	listener, err := net.Listen("tcp", ":50051")
	if err != nil {
		panic("La conexion no se pudo crear" + err.Error())
	}

	serv := grpc.NewServer()
	pb.RegisterMessageServiceServer(serv, &server{})

	if err = serv.Serve(listener); err != nil {
		panic("El server no se pudo iniciar" + err.Error())
	}
}