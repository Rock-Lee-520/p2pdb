package publish

import (
	"fmt"

	"github.com/Rock-liyi/p2pdb/application/event"
)

type LogPublisherInterface interface {
	PublishAsyncEvent(msg *LogAsyncEventMessage) (bool, error)
	NewAsyncEventMessage(msg *LogAsyncEventMessage) *LogAsyncEventMessage
}

type LogAsyncEvent struct {
	Topic string
}

type LogAsyncEventMessage struct {
	Topic string
	Data  []byte
}

type LogPublisherFactory struct{}

func (m *LogPublisherFactory) NewMessage(topic string, data []byte) *LogAsyncEventMessage {
	return &LogAsyncEventMessage{
		Topic: topic,
		Data:  data,
	}
}

var topic = "log"
var eventFunc = &event.EventFuncs{}

func (m *LogPublisherFactory) PublishAsyncEvent(msg *LogAsyncEventMessage) (bool, error) {
	if msg.Topic == "" {
		return false, fmt.Errorf("empty topic")
	}

	if msg.Topic != topic {
		return false, fmt.Errorf("error topic")
	}

	if len(msg.Data) == 0 {
		return false, fmt.Errorf("empty data")
	}

	eventFunc.PublishAsyncEvent(topic, msg.Data)
	return true, nil
}
