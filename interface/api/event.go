package api

import (
	"github.com/Rock-liyi/p2pdb/application/event"
	"github.com/Rock-liyi/p2pdb/application/event/publish"
)

type EventApiInterface interface {
	PublishAsyncEvent(eventType string, data interface{})
	PublishSyncEvent(eventType string, data interface{})
}

type EventApi struct {
}

func (e *EventApi) PublishAsyncEvent(eventType string, data interface{}) {

	var publisherFactory = &publish.LogPublisherFactory{}
	if eventType == "TestPublishAsyncEvent" {
		publisherFactory = &publish.LogPublisherFactory{}
	}
	//data = []byte{8, 0, 0, 0}
	message := publisherFactory.NewMessage(eventType, data)
	publisherFactory.PublishAsyncEvent(message)
}

func (e *EventApi) PublishSyncEvent(eventType string, data []byte) {
	// var message event.Message{Type:eventType,Data:Data}

	event.PublishSyncEvent(eventType, event.Message{Type: eventType, Data: data})
}
