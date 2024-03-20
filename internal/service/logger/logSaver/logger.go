package logSaver

import (
	"awesomeProject/internal/domain/models"
	"context"
	"encoding/json"
	"github.com/nats-io/nats.go"
	"log/slog"
)

type LogSaver struct {
	log    *slog.Logger
	stream nats.JetStreamContext
	Saver
}

func NewLogSaver(l *slog.Logger, s nats.JetStreamContext, _s Saver) *LogSaver {
	return &LogSaver{
		log:    l,
		stream: s,
		Saver:  _s,
	}
}

type Saver interface {
	Log(ctx context.Context, goods models.Goods) error
}

func (l *LogSaver) Save(ctx context.Context) error {
	msgCh := make(chan *nats.Msg)

	sub, err := l.stream.Subscribe("log.subject", func(msg *nats.Msg) {
		msgCh <- msg
	})
	if err != nil {
		l.log.Error("Error subscribing to log.subject:", err)
		return err
	}
	defer sub.Unsubscribe()

	for {
		select {
		case <-ctx.Done():
			return nil
		case msg := <-msgCh:
			var goods models.Goods
			err := json.Unmarshal(msg.Data, &goods)
			if err != nil {
				l.log.Error("Error unmarshalling goods:", err)
				continue
			}

			err = l.Saver.Log(ctx, goods)
			if err != nil {
				l.log.Error("Error saving goods:", err)
				continue
			}
		}
	}
}
