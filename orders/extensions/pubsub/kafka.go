package pubsub

import (
	"context"
	"github.com/IBM/sarama"
	"io"
	"log"
	"time"
)

type kafkaImp struct {
	producer sarama.SyncProducer
}

func NewKafkaPub(host string) IPubSub {
	kConfig := sarama.NewConfig()
	kConfig.Producer.Return.Successes = true
	kConfig.Producer.Return.Errors = true

	producer, err := sarama.NewSyncProducer([]string{host}, kConfig)
	if err != nil {
		log.Fatalf("failed to initialize NewSyncProducer, err: %v", err)
	}
	return &kafkaImp{
		producer: producer,
	}
}

func (k *kafkaImp) Produce(_ context.Context, topic string, data io.Reader) error {
	var dataBuf []byte
	if _, err := data.Read(dataBuf); err != nil {
		return err
	}

	msg := &sarama.ProducerMessage{
		Topic:     topic,
		Key:       nil,
		Value:     sarama.ByteEncoder(dataBuf),
		Timestamp: time.Now(),
	}

	partition, offset, err := k.producer.SendMessage(msg)
	if err != nil {
		return err
	}

	log.Printf("[producer] partition id: %d; offser: %d; value: %s", partition, offset, string(dataBuf))

	return nil
}
