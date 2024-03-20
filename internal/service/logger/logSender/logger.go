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
}

func NewLogSender(l *slog.Logger, s nats.JetStreamContext) *LogSender {
	return &LogSender{
		log:    l,
		stream: s,
	}
}

func (l *LogSender) Log(ctx context.Context, goods models.Goods) error {
	goodsJSON, err := json.Marshal(goods)
	if err != nil {
		l.log.Error("Error marshalling goods to JSON:", err)
		return err
	}

	err = l.send(goodsJSON)
	if err != nil {
		l.log.Error("Error sending log message:", err)
		return err
	}

	return nil
}

func (l *LogSender) send(data []byte) error {
	_, err := l.stream.Publish("log.subject", data)
	if err != nil {
		return fmt.Errorf("error publishing message to NATS: %w", err)
	}

	return nil
}
