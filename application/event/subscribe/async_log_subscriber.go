package subscribe

import (
	PS "github.com/Rock-liyi/p2pdb-pubsub"
	"github.com/Rock-liyi/p2pdb/application/event"
	"github.com/Rock-liyi/p2pdb/application/service"
	"github.com/Rock-liyi/p2pdb/infrastructure/util/log"
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

var topic = "p2pdb"
var eventFunc = &event.EventFuncs{}

func init() {
	chanEvent := make(chan event.DataEvent)
	eventFunc.RegisterAsyncEvent(topic, chanEvent)
	log.Debug("subscribe Register is ok, topic is %s", topic)
	var pub PS.PubSub

	service.InitPub(pub)
	go func() use(pub) {
		for {
			select {
			case data := <-chanEvent:
				go execute(data)
				//service.Publish(topic, data)
				//=default:  it will be caused endless loop
			}

		}
	}()
}

func execute(data event.DataEvent) {
	log.Debug("call execute method=====")
	go eventFunc.PrintDataAsyncEvent(topic, data)
}
