package service

import (
	ps "github.com/Rock-liyi/p2pdb-pubsub"
)

func InitSub() {

	var sub ps.PubSub
	sub.SetType("p2pdb")
	go sub.Sub()
}
