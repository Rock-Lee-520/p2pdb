package subscribe

import (
	"github.com/Rock-liyi/p2pdb/application/event"
	"github.com/Rock-liyi/p2pdb/application/event/publish"
	common_event "github.com/Rock-liyi/p2pdb/domain/common/event"
	"github.com/Rock-liyi/p2pdb/infrastructure/util/log"
)

func init() {
	//	debug.Dump("new event.Driver======")
	//d := new(event.Driver)
	for i := 0; i < len(common_event.StoreEventType); i++ {
		//println(common_event.StoreEventType[i])
		// Register a callback named OnSkill
		event.RegisterSyncEvent(common_event.StoreEventType[i], ExecuteLogFunc)

		// Register the global events on the OnSkill again
		//event.RegisterSyncEvent(common_event.StoreEventType[i], event.GlobalSyncEvent)
	}

	// Call the event, and all of the registered functions with the same name are called
	//	event.PublishSyncEvent("OnSkill", 100)

}

func ExecuteLogFunc(message event.Message) {
	log.Debug("call ExecuteLogFunc, message is ")
	log.Debug(message)

	//service.Publish(message.Type, message.Data)

	//PublishAsyncEvent(message.Type, message.Data)
}

func PublishAsyncEvent(eventType string, data interface{}) {

	var publisherFactory = &publish.LogPublisherFactory{}
	if eventType == "TestPublishAsyncEvent" {
		publisherFactory = &publish.LogPublisherFactory{}
	}
	//data = []byte{8, 0, 0, 0}
	message := publisherFactory.NewMessage(eventType, data)
	publisherFactory.PublishAsyncEvent(message)
}
