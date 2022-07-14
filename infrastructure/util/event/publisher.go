package event

import (
	"github.com/Rock-liyi/p2pdb/application/event/publish"
)

func PublishAsyncEvent(eventType string, data []byte) {

	var publisherFactory = &publish.LogPublisherFactory{}
	if eventType == LogTopic {
		publisherFactory = &publish.LogPublisherFactory{}
	}
	//data = []byte{8, 0, 0, 0}
	message := publisherFactory.NewMessage(eventType, data)
	publisherFactory.PublishAsyncEvent(message)
}
