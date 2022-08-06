package service

import (
	ps "github.com/Rock-liyi/p2pdb-pubsub"
	DiscoveryService "github.com/Rock-liyi/p2pdb/domain/discovery/service"
)

func InitSub() {

	var sub ps.PubSub
	sub.SetType("p2pdb")
	var subscription, err = sub.Sub()
	if err != nil {
		panic(err)
	}
	//resgister a discovery config for the application
	InitDiscovery(sub.Host.ID().String())

	go sub.StartNewSubscribeService(subscription)
}

func InitPub() {
	var pub ps.PubSub
	pub.SetType("p2pdb")
	pub.InitPub()
}

func InitDiscovery(peerId string) {

	DiscoveryService.InitInstanceInformation(peerId)
	DiscoveryService.InitPeerNodeInfomation()
}
