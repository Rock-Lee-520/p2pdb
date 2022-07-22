package subscribe

import (
	"github.com/Rock-liyi/p2pdb/application/event"
	common_event "github.com/Rock-liyi/p2pdb/domain/common/event"
	debug "github.com/favframework/debug"
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
	debug.Dump("call ExecuteLogFunc, message is ")
	debug.Dump(message)
}
