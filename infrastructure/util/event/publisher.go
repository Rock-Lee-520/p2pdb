package event

import (
	"github.com/Rock-liyi/p2pdb/application/event"
	"github.com/Rock-liyi/p2pdb/application/event/publish"
)

func PublishAsyncEvent(eventType string, data interface{}) {

	var publisherFactory = &publish.LogPublisherFactory{}
	if eventType == "TestPublishAsyncEvent" {
		publisherFactory = &publish.LogPublisherFactory{}
	}
	//data = []byte{8, 0, 0, 0}
	message := publisherFactory.NewMessage(eventType, data)
	publisherFactory.PublishAsyncEvent(message)
}

func PublishSyncEvent(eventType string, data interface{}) {
	// var message event.Message{Type:eventType,Data:Data}

	event.PublishSyncEvent(eventType, event.Message{Type: eventType, Data: data})
}
