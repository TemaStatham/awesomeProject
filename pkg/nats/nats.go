package nats

import "github.com/nats-io/nats.go"

const (
	maxPending = 256
)

func ConnectNATS(url string) (nats.JetStreamContext, error) {
	nc, err := nats.Connect(url)
	if err != nil {
		return nil, err
	}

	js, err := nc.JetStream(nats.PublishAsyncMaxPending(maxPending))
	if err != nil {
		nc.Close()
		return nil, err
	}

	_, err = js.AddStream(&nats.StreamConfig{
		Name:     "log",
		Subjects: []string{"log"},
	})
	if err != nil {
		nc.Close()
		return nil, err
	}

	return js, nil
}
