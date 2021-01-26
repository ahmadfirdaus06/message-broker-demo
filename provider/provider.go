package main

import (
	"encoding/json"
	"fmt"

	"github.com/ricochet2200/go-disk-usage/du"

	"log"

	"github.com/streadway/amqp"
)

var KB = uint64(1024)

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

type ServerStats struct {
	Free      uint64
	Available uint64
	Size      uint64
	Used      uint64
	Usage     float32
}

func main() {
	usage := du.NewDiskUsage("/")
	fmt.Println("Free:", usage.Free()/(KB*KB))
	fmt.Println("Available:", usage.Available()/(KB*KB))
	fmt.Println("Size:", usage.Size()/(KB*KB))
	fmt.Println("Used:", usage.Used()/(KB*KB))
	fmt.Println("Usage:", usage.Usage()*100, "%")

	var usagePercentage = usage.Usage() * 100

	stats := ServerStats{
		usage.Free() / (KB * KB),
		usage.Available() / (KB * KB),
		usage.Size() / (KB * KB),
		usage.Used() / (KB * KB),
		usagePercentage,
	}

	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/") //point to ip or host of RabbitMQ container
	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	q, err := ch.QueueDeclare(
		"server_stats", // name
		false,          // durable
		false,          // delete when unused
		false,          // exclusive
		false,          // no-wait
		nil,            // arguments
	)
	failOnError(err, "Failed to declare a queue")
	body, _ := json.Marshal(stats)
	err = ch.Publish(
		"",     // exchange
		q.Name, // routing key
		false,  // mandatory
		false,  // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(body),
		})
	failOnError(err, "Failed to publish a message")
	log.Printf(" [x] Sent %s", body)
}
