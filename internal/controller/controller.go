package controller

import "messageBroker/internal/service"

type Controller struct {
	svc service.IQueueBroker
}

func New(s service.IQueueBroker) *Controller {
	return &Controller{svc: s}
}
