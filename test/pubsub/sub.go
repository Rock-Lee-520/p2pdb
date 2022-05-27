package main

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/astaxie/beego"
	"github.com/libp2p/go-libp2p"
	"github.com/libp2p/go-libp2p-core/host"
	"github.com/libp2p/go-libp2p-core/peer"
	pubsub "github.com/libp2p/go-libp2p-pubsub"
	"github.com/libp2p/go-libp2p/p2p/discovery/mdns"
)

// DiscoveryInterval is how often we re-publish our mDNS records.
const DiscoveryInterval = time.Hour

// DiscoveryServiceTag is used in our mDNS advertisements to discover other chat peers.
const DiscoveryServiceTag = "p2pdb-example"

func main() {
	ctx := context.Background()
	// create a new libp2p Host that listens on a random TCP port
	h, err := libp2p.New(libp2p.ListenAddrStrings("/ip4/0.0.0.0/tcp/0"))
	if err != nil {
		panic(err)
	}

	// create a new PubSub service using the GossipSub router
	ps, err := pubsub.NewGossipSub(ctx, h)
	if err != nil {
		panic(err)
	}

	// setup local mDNS discovery
	if err := setupDiscovery(h); err != nil {
		panic(err)
	}

	topic, err := ps.Join((topicName("p2pdb")))
	if err != nil {
		panic(err)
	}

	// and subscribe to it
	sub, err := topic.Subscribe()
	if err != nil {
		panic(err)
	}

	for {
		time.Sleep(1 * time.Second)
		msg, err := sub.Next(ctx)
		if err != nil {
			panic(err)
			return
		}
		// only forward messages delivered by others

		cm := new(DataMessage)
		err = json.Unmarshal(msg.Data, cm)
		beego.Debug("debug:")
		beego.Debug(cm.Message)
		if err != nil {
			continue
		}
		beego.Debug(cm)
		// send valid messages onto the Messages channel
		//Messages <- cm
	}
}

// ChatMessage gets converted to/from JSON and sent in the body of pubsub messages.
type DataMessage struct {
	Message    string
	SenderID   string
	SenderNick string
}

func topicName(Name string) string {
	return "chat-room:" + Name
}

func sub() {

}

func pub() {

}

// discoveryNotifee gets notified when we find a new peer via mDNS discovery
type discoveryNotifee struct {
	h host.Host
}

// setupDiscovery creates an mDNS discovery service and attaches it to the libp2p Host.
// This lets us automatically discover peers on the same LAN and connect to them.
func setupDiscovery(h host.Host) error {
	// setup mDNS discovery to find local peers
	s := mdns.NewMdnsService(h, DiscoveryServiceTag, &discoveryNotifee{h: h})
	return s.Start()
}

// HandlePeerFound connects to peers discovered via mDNS. Once they're connected,
// the PubSub system will automatically start interacting with them if they also
// support PubSub.
func (n *discoveryNotifee) HandlePeerFound(pi peer.AddrInfo) {
	fmt.Printf("discovered new peer %s\n", pi.ID.Pretty())
	err := n.h.Connect(context.Background(), pi)
	if err != nil {
		fmt.Printf("error connecting to peer %s: %s\n", pi.ID.Pretty(), err)
	}
}
