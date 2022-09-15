package main

import (
	"fmt"
	"math/rand" // consultar
	"log"
	"context"
	"net"
	"time"
	amqp "github.com/rabbitmq/amqp091-go"
	"google.golang.org/grpc"
	pb "github.com/Kendovvul/Ejemplo/Proto"
)

type server struct {
	pb.UnimplementedMessageServiceServer
}

var serv *grpc.Server

func (s *server) Intercambio (ctx context.Context, msg *pb.Message) (*pb.Message, error){
	if msg.Body == "STOP MENACE"{
		serv.Stop()
		return &pb.Message{Body: "",}, nil
	}

	if msg.Body == "Equipo listo?"{
		fmt.Print("Revisando estado Escuadrón: [LISTO / NO LISTO] ")
		prob := rand.Intn(5)
		if (prob == 0 || prob == 1) {
			fmt.Print("NO LISTO\n")
			return &pb.Message{Body: "NO",}, nil
		} else{
			fmt.Print("LISTO\n")
			return &pb.Message{Body: "SI",}, nil
		}
	}else{
		return &pb.Message{Body: "pues nada",}, nil
	}

}

func main() {
	LabName := "Laboratorio Kampala" //nombre del laboratorio
	qName := "Emergencias" //nombre de la cola
	hostQ := "dist097" //ip del servidor de RabbitMQ 172.17.0.1
	connQ, err := amqp.Dial("amqp://guest:guest@"+hostQ+":5670") //conexion con RabbitMQ
	
	if err != nil {log.Fatal(err)}
	defer connQ.Close()

	ch, err := connQ.Channel()
	if err != nil{log.Fatal(err)}
	defer ch.Close()

	rand.Seed(time.Now().UnixNano())

	// Ciclo de llamados de emergencia
	for{
		fmt.Print("Analizando estado Laboratorio: [ ESTALLIDO / OK ] ")
		time.Sleep(5 * time.Second) //espera de 5 segundos
		prob := rand.Intn(5)
		if (prob != 0) {
			fmt.Print("ESTALLIDO\n")
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
			fmt.Print("SOS Enviado a Central. Esperando respuesta...\n")
			listener, err := net.Listen("tcp", ":50052") //conexion sincrona
			if err != nil {
				panic("La conexion no se pudo crear" + err.Error())
			}

			serv = grpc.NewServer()
			pb.RegisterMessageServiceServer(serv, &server{})
			if err = serv.Serve(listener); err != nil {
				panic("El server no se pudo iniciar" + err.Error())
			}

			fmt.Println("Acabaron las amenazas en este laboratorio!")
		} else{
			fmt.Println("OK")
		}
	}
}