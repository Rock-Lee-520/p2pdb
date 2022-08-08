package service

import (
	"context"
	"encoding/json"
	"time"

	PS "github.com/Rock-liyi/p2pdb-pubsub"
	"github.com/Rock-liyi/p2pdb/application/event"
	"github.com/Rock-liyi/p2pdb/application/event/subscribe"
	DiscoveryService "github.com/Rock-liyi/p2pdb/domain/discovery/service"
	"github.com/Rock-liyi/p2pdb/infrastructure/util/log"
	libpubsub "github.com/libp2p/go-libp2p-pubsub"
)

var Sub PS.PubSub

var topicType = "p2pdb"

func InitSub() {

	Sub.SetType(topicType)
	var subscription, libpubsubTpoic, err = Sub.Sub()
	if err != nil {
		panic(err)
	}
	//resgister a discovery config for the application
	InitDiscovery(Sub.Host.ID().String())
	//Sub.Pub(PS.DataMessage{Type: "InitSub", Data: "test InitSub"})
	InitLogPublisher(libpubsubTpoic)

	subscribe.InitAsyncLogSubscriber()

	go Sub.StartNewSubscribeService(subscription)

}

func InitPub(pub PS.PubSub) {

	pub.SetType("p2pdb")
	pub.InitPub()
}

func InitDiscovery(peerId string) {

	DiscoveryService.InitInstanceInformation(peerId)
	DiscoveryService.InitPeerNodeInfomation()
}

var eventFunc = &event.EventFuncs{}

func InitLogPublisher(libpubsubTpoic *libpubsub.Topic) {
	chanEvent := make(chan event.DataEvent)
	eventFunc.RegisterAsyncEvent(topicType, chanEvent)
	log.Debug("subscribe Register is ok, topic is %s", topicType)
	//execute(event.DataEvent{Data: "p2pdb_start", Topic: "p2pdb_start"}, topicType, libpubsubTpoic)

	go func() {
		for {
			select {
			case data := <-chanEvent:
				log.Debug("chanEvent  is ok, topic is %s", topicType)
				execute(data, topicType, libpubsubTpoic)
				//service.Publish(topic, data)
				//=default:  it will be caused endless loop
			}

		}
	}()
}

func execute(dataEvent event.DataEvent, topicType string, libpubsubTpoic *libpubsub.Topic) {

	log.Debug("call execute method=====")
	log.Debug(dataEvent)
	eventFunc.PrintDataAsyncEvent(topicType, dataEvent)
	ctx := context.Background()

	//for {
	time.Sleep(1 * time.Second)

	msgBytes, err := json.Marshal(dataEvent)
	if err != nil {
		panic(err)
	}
	libpubsubTpoic.Publish(ctx, msgBytes)
	time.Sleep(1 * time.Second)
	//}
}
