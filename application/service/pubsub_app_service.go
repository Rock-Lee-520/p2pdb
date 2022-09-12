package service

import (
	"context"
	discovery "github.com/Rock-liyi/p2pdb-discovery"
	PS "github.com/Rock-liyi/p2pdb-pubsub"
	"github.com/Rock-liyi/p2pdb/infrastructure/core/pubsub"
	"github.com/Rock-liyi/p2pdb/infrastructure/util/log"
)

var PubSub PS.PubSub

var topicType = "p2pdb"

const Address string = "/ip4/0.0.0.0/tcp/0"

// InitPubSub setup discovery and pubsub
func InitPubSub() string {

	ctx := context.Background()

	host, dht := discovery.InitDiscovery(ctx, Address)

	discovery.BootstrapDHT(ctx, host, dht)

	// Create a peer discovery service using the Kad DHT : 创建一个节点路由发现方式
	routingdiscovery := discovery.NewRoutingDiscovery(dht)
	//discovery.NewDiscoveryFactory().Create();
	PubSub = pubsub.InitPubSub(ctx, host, routingdiscovery, topicType, dht)
	Publish("init", "peerID:"+host.ID().String()+" | start success")
	log.Debug("========================================================")
	log.Debug("========================================================")

	for _, peer := range dht.RoutingTable().ListPeers() {
		log.Debug("peer:" + peer.String())
	}

	log.Debug("========================================================")
	log.Debug("========================================================")
	return host.ID().String()
}

// GetPubSub get pubsub connect
func GetPubSub() PS.PubSub {
	return PubSub
}

// Publish message to other nodes
func Publish(topic string, data interface{}) {
	PubSub.Pub(PS.DataMessage{Type: topic, Data: data})
}
