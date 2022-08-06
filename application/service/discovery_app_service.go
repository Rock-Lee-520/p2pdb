package service

import (
	PS "github.com/Rock-liyi/p2pdb-pubsub"
	DiscoveryService "github.com/Rock-liyi/p2pdb/domain/discovery/service"
)

var Sub PS.PubSub

func InitSub() {

	Sub.SetType("p2pdb")
	var subscription, err = Sub.Sub()
	if err != nil {
		panic(err)
	}
	//resgister a discovery config for the application
	InitDiscovery(Sub.Host.ID().String())
	//Sub.Pub(PS.DataMessage{Type: "InitSub", Data: "test InitSub"})
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
