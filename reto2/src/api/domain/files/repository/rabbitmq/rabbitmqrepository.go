package rabbitmq

import (
	"github.com/ksossa/Topicos-Telematica/reto2/src/api/infrastructure/dependencies"
	"github.com/streadway/amqp"
	"strings"
)

type RabbitRQRepository interface {
	ReadAllRabbit() ([]string, error)
}

type rabbitRQRepository struct {
	RabbitRQ *amqp.Channel
}

func NewRabbitRQRepository(container *dependencies.Container) RabbitRQRepository {
	return &rabbitRQRepository{
		RabbitRQ: container.RabbitMQ,
	}
}

func (repository *rabbitRQRepository) ReadAllRabbit() ([]string, error) {
	err := repository.RabbitRQ.Publish("", "readAll", false, false,
		amqp.Publishing{
			Headers:     nil,
			ContentType: "text/plain",
		})

	if err != nil {
		return nil, err
	}
	chDelivery, err := repository.RabbitRQ.Consume(
		"readAllAns",
		"",
		true,
		false,
		false,
		false, nil)
	for ans := range chDelivery {
		return strings.Split(string(ans.Body), ", "), nil
	}
	return nil, nil
}
