package service

import (
	PS "github.com/Rock-liyi/p2pdb-pubsub"
)

func Publish(topic string, data interface{}) {
	Sub.Pub(PS.DataMessage{Type: topic, Data: "123"})
}
