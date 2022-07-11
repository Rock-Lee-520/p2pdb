package publish

import (
	"fmt"

	"github.com/Rock-liyi/p2pdb/application/event"
)

type LogPublisherInterface interface {
	Publish(msg *LogMessage) (bool, error)
	NewMessage(msg *LogMessage) *LogMessage
}

type LogEvent struct {
	Topic string
}

type LogMessage struct {
	Topic string
	Data  []byte
}

type LogPublisherFactory struct{}

func (m *LogPublisherFactory) NewMessage(topic string, data []byte) *LogMessage {
	return &LogMessage{
		Topic: topic,
		Data:  data,
	}
}

var topic = "log"
var eventFunc = &event.EventFuncs{}

func (m *LogPublisherFactory) Publish(msg *LogMessage) (bool, error) {
	if msg.Topic == "" {
		return false, fmt.Errorf("empty topic")
	}

	if msg.Topic != topic {
		return false, fmt.Errorf("error topic")
	}

	if len(msg.Data) == 0 {
		return false, fmt.Errorf("empty data")
	}

	eventFunc.PublishEvent(topic, msg.Data)
	return true, nil
}
