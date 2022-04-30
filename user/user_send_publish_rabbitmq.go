package user

import "fmt"

type publishRabbitMQListener struct {
	data interface{}
}

func NewPublishRabbitMQListener() *publishRabbitMQListener {
	return &publishRabbitMQListener{}
}

func (l *publishRabbitMQListener) SetData(data interface{}) {
	l.data = data
}

func (l *publishRabbitMQListener) Handle() error {
	fmt.Println("Published in RabbitMQ - user: ", l.data.(string))
	return nil
}