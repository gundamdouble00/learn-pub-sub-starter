package main

import (
	"fmt"
	"os"
	"os/signal"

	amqp "github.com/rabbitmq/amqp091-go"
)

func main() {
	connectionStr := "amqp://guest:guest@localhost:5672/"
	conn, err := amqp.Dial(connectionStr)
	if err != nil {
		fmt.Println("Connection failed")
		return
	}

	defer conn.Close()
	fmt.Println("Connection succeeded")
	fmt.Println("Starting Peril server...")
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, os.Interrupt)
	<-signalChan
	fmt.Println("Connection closed")
}
