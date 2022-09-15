package main

import (
	"fmt"
	"sync"
	"log"
	"context"
	"time"
	amqp "github.com/rabbitmq/amqp091-go"
	"google.golang.org/grpc"
	pb "github.com/Kendovvul/Ejemplo/Proto"
)

// Struct de contador para escuadrones
type SafeCounter struct {
	mu sync.Mutex
	v  map[string]int
}

// Inc increments the counter for the given key.
func (c *SafeCounter) Inc(key string) {
	c.mu.Lock()
	// Lock so only one goroutine at a time can access the map c.v.
	c.v[key]++
	c.mu.Unlock()
}

func (c *SafeCounter) Dec(key string) {
	c.mu.Lock()
	// Lock so only one goroutine at a time can access the map c.v.
	c.v[key]--
	c.mu.Unlock()
}

// Value returns the current value of the counter for the given key.
func (c *SafeCounter) Value(key string) int {
	c.mu.Lock()
	// Lock so only one goroutine at a time can access the map c.v.
	defer c.mu.Unlock()
	return c.v[key]
}

// Inicializa escuadrones
var c = SafeCounter{v: make(map[string]int)}

func central () {
	qName := "Emergencias" //Nombre de la cola
	hostQ := "172.17.0.1"  //Host de RabbitMQ 172.17.0.1
	hostS := "dist098" //Host de un Laboratorio
	connQ, err := amqp.Dial("amqp://guest:guest@"+hostQ+":5672") //Conexion con RabbitMQ

	if err != nil {log.Fatal(err)}
	defer connQ.Close()

	ch, err := connQ.Channel()
	if err != nil{log.Fatal(err)}
	defer ch.Close()

	q, err := ch.QueueDeclare(qName, false, false, false, false, nil) //Se crea la cola en RabbitMQ
	if err != nil {log.Fatal(err)}

	fmt.Println(q)

	fmt.Println("Esperando Emergencias")
	chDelivery, err := ch.Consume(qName, "", true, false, false, false, nil) //obtiene la cola de RabbitMQ
	if err != nil {
		log.Fatal(err)
	}

	for delivery := range chDelivery {
		if (c.Value("a") == 0) && (c.Value("b") == 0) {
			fmt.Println("No hay equipos disponibles!")
			/*
			for cant == 0{
				time.Sleep(5 * time.Second)
			}
			*/
		} else{
			port := "0000"
			// Puerto de la conexion con el laboratorio
			if string(delivery.Body) == "Laboratorio Pripyat"{
				port = ":50051"
			}
			if string(delivery.Body) == "Laboratorio Kampala"{
				port = ":50052"
			}
			if string(delivery.Body) == "Laboratorio Renca"{
				port = ":50053"
			}
			if string(delivery.Body) == "Laboratorio Pohang"{
				port = ":50054"
			}
			// Obtiene el primer mensaje de la cola
			fmt.Println("Mensaje asíncrono de " + string(delivery.Body) + " leído")
			// Crea la conexion sincrona con el laboratorio
			connS, err := grpc.Dial(hostS + port, grpc.WithInsecure())

			if err != nil {
				panic("No se pudo conectar con el servidor" + err.Error())
			}
		
			//defer connS.Close()
		
			serviceCliente := pb.NewMessageServiceClient(connS)
			
			// Ciclo de contención de amenaza
			for {
				//espera de 5 segundos
				time.Sleep(5 * time.Second)
				//envia el mensaje al laboratorio
				res, err := serviceCliente.Intercambio(context.Background(), 
					&pb.Message{
						Body: "Equipo listo?",
					})

				if err != nil {
					panic("No se puede crear el mensaje " + err.Error())
				}
				response := res.Body
				fmt.Println(response)
				if response == "SI"{
					serviceCliente.Intercambio(context.Background(), &pb.Message{Body: "STOP MENACE",})
					break
				}
			}
			connS.Close()
			fmt.Println("Ha terminado la amenaza en " + string(delivery.Body)) //dummy out for lab name
		}
		time.Sleep(1 * time.Second)
	}

}

func main(){
	// Inicializar valor de squad
	c.Inc("a")
	c.Inc("b")
	go central()
	central()
}