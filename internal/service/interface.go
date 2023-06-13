package service

import "context"

type IQueueBroker interface {
	Append(queueName string, message interface{})
	Pull(queueName string, ctx context.Context) (string, bool)
}
