package main

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"github.com/segmentio/kafka-go"
)

const (
	topic          = "message-log"
	broker1Address = "localhost:29092"
)

func main() {
	ctx := context.Background()

	go produce(ctx)

	consume(ctx)
}

func consume(ctx context.Context) {
	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers: []string{broker1Address},
		Topic:   topic,
		GroupID: "my-group",
	})

	for {
		msg, err := r.ReadMessage(ctx)
		if err != nil {
			panic("could not read message " + err.Error())
		}

		fmt.Println("received: ", string(msg.Value))
	}
}

func produce(ctx context.Context) {
	i := 0

	w := kafka.NewWriter(kafka.WriterConfig{
		Brokers: []string{broker1Address},
		Topic:   topic,
	})

	for {
		err := w.WriteMessages(ctx, kafka.Message{
			Key:   []byte(strconv.Itoa(i)),
			Value: []byte("this is message " + strconv.Itoa(i)),
		})
		if err != nil {
			panic("could not write message " + err.Error())
		}

		fmt.Println("writes: ", i)
		i++

		time.Sleep(time.Second)
	}
}
