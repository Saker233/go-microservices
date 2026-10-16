package main

import (
	"fmt"
	"log"
	"math"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"
)

func main() {
	// connect to rabbitmq
	rabbitConn, err := connect()
	if err !=  nil {
		log.Fatal(err)
	}

	defer rabbitConn.Close()
	log.Println("Connected to RabbitMQ!")

	// start listening to messages


	// create consumer
	

	// watch the queue and consume events
}

func connect() (*amqp.Connection, error) {
	var counts int64
	var backOff =  1 * time.Second
	var connection *amqp.Connection

	// dont continue until rabbitmq is ready
	for {
		c, err := amqp.Dial("amqp://guest:guest@localhost")
		if err !=  nil {
			fmt.Println("Rqabbitmq is not yet ready")
			counts++
		} else {
			connection = c
			break
		}

		if counts > 5 {
			fmt.Println(err)
			return nil, err
		}

		backOff = time.Duration(math.Pow(float64(counts), 2)) * time.Second
		log.Println("backing off.....")
		time.Sleep(backOff)
		continue
	}

	return connection, nil
}