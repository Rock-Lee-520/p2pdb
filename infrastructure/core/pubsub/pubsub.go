package pubsub

import (
	pb "github.com/Rock-liyi/p2pdb-pubsub"
)

func Publish(topic string, data interface{}) {
	var PubSub pb.PubSub
	PubSub.InitPubSub()
	PubSub.Pub(pb.DataMessage{Type: topic, Data: "123"})
}
