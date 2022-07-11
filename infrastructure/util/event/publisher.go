package event

import (
	"github.com/Rock-liyi/p2pdb/application/event/publish"
)

type Event struct {
	Type string `json:"type"`
	Data string `json:"data"`
}

type EventType struct {
	LOG       string
	DISCOVERY string
}

func New() EventType {
	return EventType{
		LOG: "log", DISCOVERY: "discover"}
}

const (
	LogTopic = "log"
)

func PublishEvent(eventType string, data []byte) {
	var publisherFactory = &publish.LogPublisherFactory{}
	if eventType == LogTopic {
		publisherFactory = &publish.LogPublisherFactory{}
	}
	data = []byte{8, 0, 0, 0}
	message := publisherFactory.NewMessage(eventType, data)
	publisherFactory.Publish(message)
}
