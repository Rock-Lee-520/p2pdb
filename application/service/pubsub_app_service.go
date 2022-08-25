package service

import (
	PS "github.com/Rock-liyi/p2pdb-pubsub"
	pubsub "github.com/Rock-liyi/p2pdb/infrastructure/core/pubsub"
	"github.com/Rock-liyi/p2pdb/infrastructure/util/log"
)

func Publish(topic string, data interface{}) {
	//PubSub.Pub(PS.DataMessage{Type: topic, Data: "123"})
}

var PubSub PS.PubSub

var topicType = "p2pdb"

const Address string = "/ip4/0.0.0.0/tcp/0"

func InitPubSub() string {

	PubSub = pubsub.InitPubSub()
	log.Debug("init PubSub module finished")
	//var subscription, libpubsubTpoic, err = PubSub.Sub()
	//if err != nil {
	//	panic(err)
	//}
	// resgister a discovery config for the application
	var instanceId = InitDiscovery(PubSub.Host.ID().String())
	//PubSub.Pub(PS.DataMessage{Type: "InitPubSub", Data: "test InitPubSub"})
	InitLogPublisher(PubSub.Topic)

	//subscribe.InitAsyncLogSubscriber()

	return instanceId
}

func GetPubSub() PS.PubSub {

	PubSub = pubsub.InitPubSub()

	return PubSub
}
