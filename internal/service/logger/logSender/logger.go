package logSender

import (
	"awesomeProject/internal/domain/models"
	"context"
	"encoding/json"
	"fmt"
	"github.com/nats-io/nats.go"
	"log/slog"
)

type LogSender struct {
	log    *slog.Logger
	stream nats.JetStreamContext
	ch     chan models.Goods
}

func NewLogSender(l *slog.Logger, s nats.JetStreamContext) *LogSender {
	return &LogSender{
		log:    l,
		stream: s,
	}
}

func (l *LogSender) Log(ctx context.Context, goods models.Goods) error {
	l.ch <- goods
	return nil
}

func (l *LogSender) send(data []byte) error {
	_, err := l.stream.Publish("log", data)
	if err != nil {
		return fmt.Errorf("error publishing message to NATS: %w", err)
	}
	return nil
}

func (l *LogSender) Start(ctx context.Context) {
	go func() {
		for {
			select {
			case <-ctx.Done():
				return
			case goods := <-l.ch:
				// Преобразуем данные в JSON
				goodsJSON, err := json.Marshal(goods)
				if err != nil {
					l.log.Error("Error marshalling goods to JSON:", err)
					continue
				}

				// Отправляем данные методу send
				if err := l.send(goodsJSON); err != nil {
					l.log.Error("Error sending log message:", err)
					continue
				}
			}
		}
	}()
}
