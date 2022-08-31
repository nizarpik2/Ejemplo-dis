package main

import (
	"fmt"
	"log"
	amqp "github.com/rabbitmq/amqp091-go"
	pb "github.com/Sistemas-Distribuidos-2022-2/Ejemplo/Proto"
)

type server struct {
	pb.UnimplementedMessageServiceServer
}

func (s *server) Intercambio (_ context.Context, msg *pb.Message) (*pb.Message, error){
	fmt.Println("mesaje sincrono")
	return &pb.Message{body: msg.body,}, nil
}

func main () {
	qName := "Emergencias"
	host := "localhost"
	conn, err := amqp.Dial("amqp://guest:guest@"+host+":5672")

	if err != nil {log.Fatal(err)}
	defer conn.Close()

	ch, err := conn.Channel()
	if err != nil{log.Fatal(err)}
	defer ch.Close()

	chDelivery, err := ch.Consume(qName, "", true, false, false, false, nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Esperando Emergencias")
	noStop := make(chan bool)
	go func () {
		for delivery := range chDelivery {
			fmt.Println("Pedido de ayuda de " + string(delivery.Body))
		}
 	}()
	
	<- noStop

	hostS := "localhost"
	
	connS, grpc.Dial(host + ":50051", grpc.WithInsecure())

	if err != nil {
		panic("No se pudo conectar con el servidor" + err.Error())
	}

	defer connS.Close()

	serviceCliente := pb.NewMessageServiceClient(connS)

	serviceCliente.Intercambio(context.Background(), &pb.Message{
		body: "Equipo listo?",
		}
	)

	if err != nil {
		panic("No se puede crear el mensaje " + err.Error())
	}

	fmt.Println()
}