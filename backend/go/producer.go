package main

import (
    "context"
    "encoding/json"
    "log"

    "github.com/segmentio/kafka-go"
)

func sendFlightUpdate(flight Flight) {
    w := kafka.Writer{
        Addr:     kafka.TCP("localhost:9092"),
        Topic:    "flight-updates",
        Balancer: &kafka.LeastBytes{},
    }

    msg, err := json.Marshal(flight)
    if err != nil {
        log.Println("Error marshalling flight:", err)
        return
    }

    err = w.WriteMessages(context.Background(),
        kafka.Message{
            Value: msg,
        },
    )
    if err != nil {
        log.Println("Error writing message to Kafka:", err)
        return
    }

    w.Close()
}
