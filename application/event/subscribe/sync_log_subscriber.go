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

	var newData entity.Data
	function.JsonDecode(message.Data, &newData)
	log.Debug(newData)
	switch message.Type {
	case value_object.StoreCreateTableEvent:
		service.InitLogTable(newData)
		service.CreateTableByStoreEvent(newData)
	case value_object.StoreInsertEvent:
		service.InsertStoreLog(newData, value_object.StoreInsertEvent)
	case value_object.StoreDeleteEvent:
		service.InsertStoreLog(newData, value_object.StoreDeleteEvent)
	case value_object.StoreUpdateEvent:
		service.InsertStoreLog(newData, value_object.StoreUpdateEvent)
		log.Debug(newData)
		service.InsertStoreLog(newData, value_object.StoreUpdateEvent)
	}

}
