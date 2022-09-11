package main

import (
	"github.com/Rock-liyi/p2pdb/application/service"
)

func main() {
	//ctx := context.Background()
	//// create a new libp2p Host that listens on a random TCP port
	//h, err := libp2p.New(libp2p.ListenAddrStrings("/ip4/0.0.0.0/tcp/0"))
	//if err != nil {
	//	panic(err)
	//}
	//
	//// create a new PubSub service using the GossipSub router
	//ps, err := pubsub.NewGossipSub(ctx, h)
	//if err != nil {
	//	panic(err)
	//}
	//
	//// setup local mDNS discovery
	//if err := setupDiscovery(h); err != nil {
	//	panic(err)
	//}
	//
	//topic, err := ps.Join((topicName("p2pdb")))
	//if err != nil {
	//	panic(err)
	//}
	//
	//for {
	//	time.Sleep(1 * time.Second)
	//	input := DataMessage{
	//		Data: "Publish",
	//		Type: "store_drop_table_event",
	//	}
	//
	//	msgBytes, err := json.Marshal(input)
	//	if err != nil {
	//		panic(err)
	//	}
	//	topic.Publish(ctx, msgBytes)
	//	time.Sleep(1 * time.Second)
	//}

	service.InitPubSub()
	service.Publish("test", "我就是来测试消息的")

}

//
//type DataMessage struct {
//	Type string
//	Data interface{}
//}
//
//// ChatMessage gets converted to/from JSON and sent in the body of pubsub messages.
//// type DataMessage struct {
//// 	Message    string
//// 	SenderID   string
//// 	SenderNick string
//// 	Topic      string
//// }
//
//func topicName(Name string) string {
//	return Name
//}
//
//// discoveryNotifee gets notified when we find a new peer via mDNS discovery
//type discoveryNotifee struct {
//	h host.Host
//}
//
//// setupDiscovery creates an mDNS discovery service and attaches it to the libp2p Host.
//// This lets us automatically discover peers on the same LAN and connect to them.
//func setupDiscovery(h host.Host) error {
//	// setup mDNS discovery to find local peers
//	s := mdns.NewMdnsService(h, DiscoveryServiceTag, &discoveryNotifee{h: h})
//	return s.Start()
//}
//
//// HandlePeerFound connects to peers discovered via mDNS. Once they're connected,
//// the PubSub system will automatically start interacting with them if they also
//// support PubSub.
//func (n *discoveryNotifee) HandlePeerFound(pi peer.AddrInfo) {
//	fmt.Printf("discovered new peer %s\n", pi.ID.Pretty())
//	err := n.h.Connect(context.Background(), pi)
//	if err != nil {
//		fmt.Printf("error connecting to peer %s: %s\n", pi.ID.Pretty(), err)
//	}
//}
