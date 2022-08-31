package main

import (
	"fmt"
	amqp "github.com/rabbitmq/amqp091-go"
	"log"
)

func main() {
	LabName := "Laboratiorio Pripyat"
	qName := "Emergencias"
	host := "localhost"
	conn, err := amqp.Dial("amqp://guest:guest@"+host+":5672")
	
	if err != nil {
		log.Fatal(err)
	}

	defer conn.Close()

	ch, err := conn.Channel()

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
}