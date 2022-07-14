package event

import "fmt"

// 声明角色的结构体
type Driver struct {
}

// Add an event-processing function for function
func (a *Driver) OnEvent(param interface{}) {

	fmt.Println("actor event:", param)
}

func GlobalEvent(param interface{}) {

	fmt.Println("global event:", param)
}

func CallEvent(name string, param interface{}) {

	// Find the list of events by the name
	list := eventByName[name]

	// Go through all the callbacks for this event
	for _, callback := range list {

		//The incoming parameter calls a callback
		callback(param)
	}

}

// Inantiate a map that slices from a mapping function through a string
var eventByName = make(map[string][]func(interface{}))

// Register the events, providing the event name and the callback function
func RegisterEvent(name string, callback func(interface{})) {

	// Find the list of events by the name
	list := eventByName[name]

	// Add functions in the list slice
	list = append(list, callback)

	// Save the modified event list slice back in eventByName
	eventByName[name] = list
}
