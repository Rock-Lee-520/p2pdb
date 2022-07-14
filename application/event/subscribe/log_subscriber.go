package subscribe

import (
	"log"

	"github.com/Rock-liyi/p2pdb/application/event"
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

var topic = "log"
var eventFunc = &event.EventFuncs{}

func init() {
	chanEvent := make(chan event.DataEvent)
	eventFunc.RegisterAsyncEvent(topic, chanEvent)
	log.Printf("subscribe Register is ok, topic is %s", topic)
	go func() {
		for {
			select {
			case data := <-chanEvent:
				go execute(data)
				//=default:  it will be caused endless loop
			}

		}
	}()
}

func execute(data event.DataEvent) {
	go eventFunc.PrintDataAsyncEvent(topic, data)
}
