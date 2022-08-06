package event

import (
	"github.com/Rock-liyi/p2pdb/infrastructure/util/log"
)

type Driver struct {
}

// Add an event-processing function for function
func (a *Driver) OnEvent(param interface{}) {
	log.Debug("actor event:")
	log.Debug(param)
}

func GlobalSyncEvent(param interface{}) {
	log.Debug("global event:")
	log.Debug(param)
}

func PublishSyncEvent(name string, message Message) {
	// Find the list of events by the name
	list := eventByName[name]

	// Go through all the callbacks for this event
	for _, callback := range list {

		//The incoming parameter calls a callback
		callback(message)
	}

}

// Inantiate a map that slices from a mapping function through a string
var eventByName = make(map[string][]func(Message))

// Register the events, providing the event name and the callback function
func RegisterSyncEvent(name string, callback func(Message)) {
	//debug.Dump("call RegisterSyncEvent======")
	// Find the list of events by the name
	list := eventByName[name]

	// Add functions in the list slice
	list = append(list, callback)

	// Save the modified event list slice back in eventByName
	eventByName[name] = list
}
