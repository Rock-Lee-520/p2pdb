package pubsub

import (
	"context"
	PS "github.com/Rock-liyi/p2pdb-pubsub"
	"github.com/libp2p/go-libp2p-core/host"
	discovery "github.com/libp2p/go-libp2p-discovery"
	dht "github.com/libp2p/go-libp2p-kad-dht"
)

// func Publish(topic string, data interface{}) {
// 	var PubSub pb.PubSub
// 	PubSub.InitPubSub()
// 	PubSub.Pub(pb.DataMessage{Type: topic, Data: "123"})
// }

var PubSub PS.PubSub

func InitPubSub(ctx context.Context, host host.Host, Routingdiscovery *discovery.RoutingDiscovery, topicType string, dht *dht.IpfsDHT) PS.PubSub {
	return PubSub.InitPubSub(ctx, topicType, host, Routingdiscovery, dht)
}
