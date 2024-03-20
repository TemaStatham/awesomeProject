package nats

import "github.com/nats-io/nats.go"

func ConnectNATS() (nats.JetStreamContext, error) {
	// Подключение к серверу NATS
	nc, err := nats.Connect(nats.DefaultURL)
	if err != nil {
		return nil, err
	}

	// Создание контекста для работы с JetStream
	js, err := nc.JetStream(nats.PublishAsyncMaxPending(256))
	if err != nil {
		nc.Close() // Закрываем соединение в случае ошибки
		return nil, err
	}

	return js, nil
}
