package service

import (
	"container/list"
	"context"
	"sync"
)

type QueueBroker struct {
	queues map[string]*list.List
	mutex  sync.Mutex
}

func NewQueueBroker() *QueueBroker {
	return &QueueBroker{
		queues: make(map[string]*list.List),
	}
}

func (qb *QueueBroker) Append(queueName string, message interface{}) {
	qb.mutex.Lock()
	defer qb.mutex.Unlock()

	if _, ok := qb.queues[queueName]; !ok {
		qb.queues[queueName] = list.New()
	}

	qb.queues[queueName].PushBack(message)
}

func (qb *QueueBroker) Pull(queueName string, ctx context.Context) (string, bool) {
	for {
		select {
		case <-ctx.Done():
			return "", false
		default:
			if q, ok := qb.queues[queueName]; ok && q.Len() > 0 {
				el := q.Front()
				if el != nil {
					q.Remove(el)
					return el.Value.(string), true
				}
			}
		}
	}
}
