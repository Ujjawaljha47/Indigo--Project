
package main

import (
    "context"
    "log"
    "github.com/segmentio/kafka-go"
)

func consumeFlightUpdates() {
    r := kafka.NewReader(kafka.ReaderConfig{
        Brokers:   []string{"localhost:9092"},
        Topic:     "flight-updates",
        Partition: 0,
        MinBytes:  10e3, // 10KB
        MaxBytes:  10e6, // 10MB
    })

    for {
        m, err := r.ReadMessage(context.Background())
        if err != nil {
            log.Println("Error reading message from Kafka:", err)
            continue
        }
        log.Printf("Message at offset %d: %s = %s\n", m.Offset, string(m.Key), string(m.Value))
    }
}
