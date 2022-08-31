package main

import (
	"fmt"
	amqp "github.com/rabbitmq/amqp091-go"
	"log"
)

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

}