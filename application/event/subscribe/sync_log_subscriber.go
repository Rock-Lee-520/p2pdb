package subscribe

import (
	"github.com/Rock-liyi/p2pdb/application/event"
	"github.com/Rock-liyi/p2pdb/application/service"
	commonEvent "github.com/Rock-liyi/p2pdb/domain/common/event"
	StoreEvent "github.com/Rock-liyi/p2pdb/domain/store/event"
	"github.com/Rock-liyi/p2pdb/infrastructure/util/log"
)

func init() {
	//	debug.Dump("new event.Driver======")
	//d := new(event.Driver)
	for i := 0; i < len(commonEvent.StoreEventType); i++ {
		//println(common_event.StoreEventType[i])
		// Call the event, and all of the registered functions with the same name are called
		event.RegisterSyncEvent(commonEvent.StoreEventType[i], ExecuteLogFunc)

		// Register the global events on the OnSkill again
		//event.RegisterSyncEvent(common_event.StoreEventType[i], event.GlobalSyncEvent)
	}
}

func ExecuteLogFunc(message event.Message) {
	log.Debug("call ExecuteLogFunc, message is ")
	log.Debug(message.Type)
	log.Debug(message.Data)
	switch message.Type {
	case StoreEvent.StoreCreateTableEvent:
		service.CreateTableByStoreEvent(message)
	}
	//	event.RegisterSyncEvent(common_event.StoreEventType[i], ExecuteLogFunc)

	// Register the global events on the OnSkill again
	//event.RegisterSyncEvent(common_event.StoreEventType[i], event.GlobalSyncEvent)

}

// func PublishAsyncEvent(eventType string, data interface{}) {

// 	var publisherFactory = &publish.LogPublisherFactory{}
// 	if eventType == "TestPublishAsyncEvent" {
// 		publisherFactory = &publish.LogPublisherFactory{}
// 	}
// 	//data = []byte{8, 0, 0, 0}
// 	message := publisherFactory.NewMessage(eventType, data)
// 	publisherFactory.PublishAsyncEvent(message)
// }
