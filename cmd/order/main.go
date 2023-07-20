package main

import (
	"database/sql"
	"encoding/json"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
	"github.com/nathanSeixeiro/go-intensivo/internal/infra/database"
	usecases "github.com/nathanSeixeiro/go-intensivo/internal/useCases"
	"github.com/nathanSeixeiro/go-intensivo/pkg/rabbitmq"
	amqp "github.com/rabbitmq/amqp091-go"
)

func main() {
	db, err := sql.Open("sqlite3", "db.sqlite3")
	if err != nil {
		panic(err)
	}
	defer db.Close() // waited run all of them and after close connection
	orderRepository := database.NewOrderRepository(db)
	uc := usecases.NewCalculateFinalPrice(orderRepository)
	ch, err := rabbitmq.OpenChannel()
	if err != nil {
		panic(err)
	}
	defer ch.Close()
	msgRabbitmqChannel := make(chan amqp.Delivery)
	go rabbitmq.Consume(ch, msgRabbitmqChannel) // async process (listening queue), put on second thread
	RabbitmqWorker(msgRabbitmqChannel, uc)      // in the first thread save the msgs and execute usecases and the second stay listen the queue
}

func RabbitmqWorker(msgChan chan amqp.Delivery, uc *usecases.CalculateFinalPrice) {
	fmt.Print("Starting Rabbitmq")
	for msg := range msgChan {
		var input usecases.OrderInput
		err := json.Unmarshal(msg.Body, &input)
		if err != nil {
			panic(err)
		}
		out, err := uc.Execute(input)
		if err != nil {
			panic(err)
		}
		msg.Ack(false)
		fmt.Println("message saved on database", out)
	}
}
