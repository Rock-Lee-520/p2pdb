package pubsub

import (
	discovery "github.com/Rock-liyi/p2pdb-discovery"
	PS "github.com/Rock-liyi/p2pdb-pubsub"
	"github.com/Rock-liyi/p2pdb/infrastructure/util/log"
)

// func Publish(topic string, data interface{}) {
// 	var PubSub pb.PubSub
// 	PubSub.InitPubSub()
// 	PubSub.Pub(pb.DataMessage{Type: topic, Data: "123"})
// }

var PubSub PS.PubSub

var topicType = "p2pdb"

const Address string = "/ip4/0.0.0.0/tcp/0"

func InitPubSub() PS.PubSub {

	host, err := discovery.NewDiscoveryFactory().Create(Address)
	if err != nil {
		panic(err)
	}
	log.Debug("start init PubSub module")
	PubSub = PubSub.InitPubSub(topicType, host)
	return PubSub
}
