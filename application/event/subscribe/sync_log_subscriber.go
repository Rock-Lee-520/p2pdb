package subscribe

import (
	PS "github.com/Rock-liyi/p2pdb-pubsub"
	"github.com/Rock-liyi/p2pdb-store/entity/value_object"
	"github.com/Rock-liyi/p2pdb/application/event"
	"github.com/Rock-liyi/p2pdb/application/service"
	entity "github.com/Rock-liyi/p2pdb/domain/common/entity"
	"github.com/Rock-liyi/p2pdb/infrastructure/util/function"
	"github.com/Rock-liyi/p2pdb/infrastructure/util/log"
)

var PubSub PS.PubSub

func init() {
	// TODO: 暂时注释掉了，这里应该在pubsub模块初始化完成后，才执行
	//PubSub = service.GetPubSub()

	//	debug.Dump("new event.Driver======")
	//d := new(event.Driver)
	//for i := 0; i < len(value_object.StoreEventType); i++ {
	//	//println(common_event.StoreEventType[i])
	//	// Call the event, and all of the registered functions with the same name are called
	//	event.RegisterSyncEvent(value_object.StoreEventType[i], ExecuteLogFunc)
	//
	//	// Register the global events on the OnSkill again
	//	//event.RegisterSyncEvent(common_event.StoreEventType[i], event.GlobalSyncEvent)
	//}
}

func ExecuteLogFunc(message event.Message) {
	log.Debug("call ExecuteLogFunc, message is ")

	var newData entity.Data
	switch message.Type {
	case value_object.StoreCreateTableEvent:
		function.JsonDecode(message.Data, &newData)
		service.CreateStoteLogTable(newData)
		service.AddTableRecord(newData)
		//service.CreateTableByStoreEvent(newData)
	case value_object.StoreDropTableEvent:
		function.JsonDecode(message.Data, &newData)
		service.DropStoteLogTable(newData)

	case value_object.StoreAlterTableRenameEvent:
		function.JsonDecode(message.Data, &newData)
		service.RenameStoteLogTable(newData)
		service.ChangeTableRecord(newData)

	case value_object.StoreInsertEvent:
		function.JsonDecode(message.Data, &newData)
		service.AddStoreLog(newData, value_object.StoreInsertEvent)

	case value_object.StoreDeleteEvent:
		function.JsonDecode(message.Data, &newData)
		service.AddStoreLog(newData, value_object.StoreDeleteEvent)

	case value_object.StoreUpdateEvent:
		function.JsonDecode(message.Data, &newData)
		service.AddStoreLog(newData, value_object.StoreUpdateEvent)
	}
	log.Debug(newData)

	PubSub.Pub(PS.DataMessage{Type: PubSub.GetType(), Data: message.Data})
}
