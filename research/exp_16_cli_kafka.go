package main

import (
	"context"
	"errors"
	"fmt"

	kafka "github.com/segmentio/kafka-go"
)

type work struct {
	Code int
	E    error
}

func kafkaConsumer() {

	defer func() {

		if r := recover(); r != nil {

			if _, ok := r.(error); ok {
				fmt.Println("Error:OK ")
			} else {
				fmt.Println("Error:NOT OK")
			}

		}

		w := work{
			Code: 0,
			E:    errors.New(" exit go!!!"),
		}
		consumer <- w
	}()

	// to consume messages
	topic := "message-test"
	partition := 0

	// conn, err := kafka.DialLeader(context.Background(), "tcp", "172.10.50.238:9092", topic, partition)
	// if err != nil {
	// 	log.Fatal("failed to dial leader:", err)
	// }

	r := kafka.NewReader(kafka.ReaderConfig{Brokers: []string{"172.10.60.58:9092"},
		Topic:     topic,
		Partition: partition,
		GroupID:   "group-id",
		MaxBytes:  10e6})

	defer r.Close()

	ctx := context.Background()

	for {
		m, e := r.FetchMessage(ctx)
		if e != nil {
			fmt.Println("This is a line")
			fmt.Println(e.Error())
			break
		}

		fmt.Println(string(string(m.Value)))

		if e := r.CommitMessages(ctx, m); e != nil {
			fmt.Println("BBBBBBB")
			fmt.Println(e.Error())
			break
		}

	}
}

var consumer = make(chan work)

func main() {

	go kafkaConsumer()

	for {
		select {
		case w := <-consumer:
			fmt.Println(w.Code)
			fmt.Println(w.E)
			go kafkaConsumer()
		}
	}

}
