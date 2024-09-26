package main

import (
	"context"
	"github.com/segmentio/kafka-go"
	"log"
)

func main() {
	// Configuração do escritor Kafka
	writer := kafka.NewWriter(kafka.WriterConfig{
		Brokers:  []string{"fullcycle-kafka-kafka-1:9092"}, // Endereço do broker Kafka
		Topic:    "teste",                                  // Nome do tópico
		Balancer: &kafka.LeastBytes{},
	})
	// Publicando uma mensagem no tópico
	err := writer.WriteMessages(context.Background(),
		kafka.Message{
			Key:   []byte("Key-A"),
			Value: []byte("mensagem"),
		},
	)
	if err != nil {
		log.Fatal("Erro ao enviar mensagem:", err)
	}
	log.Println("Mensagem enviada com sucesso!")
	writer.Close()
}
