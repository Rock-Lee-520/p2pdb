package pubsub

import (
	"context"
	PS "github.com/Rock-liyi/p2pdb-pubsub"
	"github.com/libp2p/go-libp2p-core/host"
	discovery "github.com/libp2p/go-libp2p-discovery"
)

// func Publish(topic string, data interface{}) {
// 	var PubSub pb.PubSub
// 	PubSub.InitPubSub()
// 	PubSub.Pub(pb.DataMessage{Type: topic, Data: "123"})
// }

var PubSub PS.PubSub

var topicType = "p2pdb"

const Address string = "/ip4/0.0.0.0/tcp/0"

func InitPubSub(ctx context.Context, host host.Host, Routingdiscovery *discovery.RoutingDiscovery) PS.PubSub {
	return PubSub.InitPubSub(ctx, topicType, host, Routingdiscovery)
}
