package main

import (
	"flag"
	"messageBroker/internal/controller"
	"messageBroker/internal/server"
	"messageBroker/internal/service"
)

var port string

func main() {
	// для старта приложения потребуется указать параметр --port={port} в консоли
	// если порт не указан через консоль, то приложение по умолчанию берет порт 8080
	flag.Func("port", "port", func(s string) error {
		port = s
		return nil
	})
	flag.Parse()

	// инициализация сервиса очереди
	svc := service.NewQueueBroker()

	// инициализация обработчиков
	_controller := controller.New(svc)

	// инициализация сервера
	srv := server.New(_controller)

	// запуск сервера
	srv.Run(port)

}
