package subscribe

import (
	"github.com/Rock-liyi/p2pdb/application/event"
	debug "github.com/favframework/debug"
)

// type selector struct {
// 	name string
// }

// type subscribe struct {
// 	selector
// }
//  func (s *subscribe) String() string {
// 	//     return s.name

// }

func init() {
	debug.Dump("new event.Driver======")
	d := new(event.Driver)

	// Register a callback named OnSkill
	event.RegisterSyncEvent("OnSkill", d.OnEvent)

	// Register the global events on the OnSkill again
	event.RegisterSyncEvent("OnSkill", event.GlobalSyncEvent)

	// Call the event, and all of the registered functions with the same name are called
	event.PublishSyncEvent("OnSkill", 100)

}

func sync_log_execute(data event.DataEvent) {

}
