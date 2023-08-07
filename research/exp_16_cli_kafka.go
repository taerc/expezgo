package main

import (
	"context"
	"fmt"

	kafka "github.com/segmentio/kafka-go"
)

func main() {

	// to consume messages
	topic := "message-test"
	partition := 0

	// conn, err := kafka.DialLeader(context.Background(), "tcp", "172.10.50.238:9092", topic, partition)
	// if err != nil {
	// 	log.Fatal("failed to dial leader:", err)
	// }

	r := kafka.NewReader(kafka.ReaderConfig{Brokers: []string{"172.10.50.238:9092"},
		Topic:     topic,
		Partition: partition,
		GroupID:   "group-id",
		MaxBytes:  10e6})

	defer r.Close()

	ctx := context.Background()

	for {
		m, e := r.FetchMessage(ctx)
		if e != nil {
			fmt.Println(e.Error())
			break
		}

		fmt.Println(string(string(m.Value)))

		if e := r.CommitMessages(ctx, m); e != nil {
			fmt.Println(e.Error())
			break
		}

	}
}
