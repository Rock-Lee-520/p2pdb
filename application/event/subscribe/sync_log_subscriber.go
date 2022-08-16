package subscribe

import (
	"github.com/Rock-liyi/p2pdb/application/event"
	"github.com/Rock-liyi/p2pdb/application/service"
	entity "github.com/Rock-liyi/p2pdb/domain/common/entity"
	value_object "github.com/Rock-liyi/p2pdb/domain/common/entity/value_object"
	"github.com/Rock-liyi/p2pdb/infrastructure/util/function"
	"github.com/Rock-liyi/p2pdb/infrastructure/util/log"
)

func init() {
	//	debug.Dump("new event.Driver======")
	//d := new(event.Driver)
	for i := 0; i < len(value_object.StoreEventType); i++ {
		//println(common_event.StoreEventType[i])
		// Call the event, and all of the registered functions with the same name are called
		event.RegisterSyncEvent(value_object.StoreEventType[i], ExecuteLogFunc)

		// Register the global events on the OnSkill again
		//event.RegisterSyncEvent(common_event.StoreEventType[i], event.GlobalSyncEvent)
	}
}

func ExecuteLogFunc(message event.Message) {
	log.Debug("call ExecuteLogFunc, message is ")
	// log.Debug(message.Type)
	// log.Debug(message.Data)
	switch message.Type {
	case value_object.StoreCreateTableEvent:
		var newData entity.Data
		function.JsonDecode(message.Data, &newData)

		service.InitLogTable(newData)
		service.CreateTableByStoreEvent(newData)
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
