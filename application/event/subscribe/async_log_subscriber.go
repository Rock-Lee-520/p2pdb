package subscribe

import (
	"github.com/Rock-liyi/p2pdb/application/event"
	common_event "github.com/Rock-liyi/p2pdb/domain/common/event"
	"github.com/Rock-liyi/p2pdb/infrastructure/util/log"
)

var topic = "p2pdb"
var eventFunc = &event.EventFuncs{}

func init() {
	chanEvent := make(chan event.DataEvent)
	eventFunc.RegisterAsyncEvent(topic, chanEvent)
	log.Debug("subscribe Register is ok, topic is %s", topic)

	for i := 0; i < len(common_event.StoreEventType); i++ {
		eventFunc.RegisterAsyncEvent(topic, chanEvent)
	}

	// service.InitPub(pub)
	go func() {
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
