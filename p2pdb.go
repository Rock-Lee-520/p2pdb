package main

// This is a CLI that lets you join a global permissionless CRDT-based
// database using CRDTs and IPFS.

import (
	"bufio"
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"go-libp2p"
	"go-libp2p/p2p/discovery/mdns"
	"log"
	"math/rand"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"

	"github.com/astaxie/beego"
	logging "github.com/ipfs/go-log/v2"

	"github.com/libp2p/go-libp2p-core/host"
	"github.com/libp2p/go-libp2p-core/peer"
	pubsub "github.com/libp2p/go-libp2p-pubsub"

	multiaddr "github.com/multiformats/go-multiaddr"

	_ "github.com/mattn/go-sqlite3"
)

var (
	logger    = logging.Logger("p2pdb")
	listen    = "/ip4/0.0.0.0/tcp/0"
	topicName = "p2pdb-example"
	netTopic  = "p2pdb-example-net"
	config    = "p2pdb-example"
	dbPath    = "./data/"
	dbName    = "p2pdb.db"
	db        *sql.DB // 全局变量
	selfId    peer.ID
)

// DiscoveryInterval is how often we re-publish our mDNS records.
const DiscoveryInterval = time.Hour

// DiscoveryServiceTag is used in our mDNS advertisements to discover other chat peers.
const DiscoveryServiceTag = "p2pdb-example"

// ChatMessage gets converted to/from JSON and sent in the body of pubsub messages.
type P2pdbLog struct {
	Id   int
	Type string
	Sql  string
}

// discoveryNotifee gets notified when we find a new peer via mDNS discovery
type discoveryNotifee struct {
	h host.Host
}

func main() {
	ctx := context.Background()
	// create a new libp2p Host that listens on a random TCP port
	h, err := libp2p.New(libp2p.ListenAddrStrings(listen))
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
	topic, err := ps.Join(topicName)
	if err != nil {
		panic(err)
	}

	//pub(ps, ctx, topic)
	selfId := h.ID()
	sub(ps, ctx, topic, h, selfId)
	run(h, ctx, topic)

}

func pub(ps *pubsub.PubSub, ctx context.Context, topic *pubsub.Topic) {

	go func() {
		for {
			time.Sleep(1 * time.Second)
			input := P2pdbLog{
				Id:   00000,
				Type: "insert",
				Sql:  "insert into p2pdb(id, name) values(1, 'foo'), (2, 'bar'), (3, 'baz');",
			}

			msgBytes, err := json.Marshal(input)
			if err != nil {
				panic(err)
			}
			topic.Publish(ctx, msgBytes)
		}
	}()

}

func publish(Type string, Sql string, ctx context.Context, topic *pubsub.Topic) {
	input := P2pdbLog{
		Id:   rand.Intn(30),
		Type: Type,
		Sql:  Sql,
	}
	msgBytes, err := json.Marshal(input)
	if err != nil {
		panic(err)
	}
	topic.Publish(ctx, msgBytes)
}

func sub(ps *pubsub.PubSub, ctx context.Context, topic *pubsub.Topic, h host.Host, selfId peer.ID) {

	// and subscribe to it
	sub, err := topic.Subscribe()
	if err != nil {
		panic(err)
	}
	go func() {
		for {
			time.Sleep(1 * time.Second)
			msg, err := sub.Next(ctx)
			if err != nil {
				panic(err)
			}
			// beego.Debug("ReceivedFrom:")
			// beego.Debug(msg.ReceivedFrom)

			// beego.Debug("selfId:")
			// beego.Debug(selfId)

			// only forward messages delivered by others
			if msg.ReceivedFrom == selfId {
				continue
			}
			cm := new(P2pdbLog)
			err = json.Unmarshal(msg.Data, cm)
			beego.Debug("debug:")
			beego.Debug(cm.Sql)
			if err != nil {
				continue
			}

			fields := strings.Fields(cm.Sql)
			if len(fields) == 0 {
				fmt.Printf("> ")
				continue
			}

			execute(cm.Type, fields, cm.Sql, ctx, topic, h, false)
			//beego.Debug(cm)
			// send valid messages onto the Messages channel
			//Messages <- cm
		}
	}()

}

func run(h host.Host, ctx context.Context, topic *pubsub.Topic) {
	fmt.Printf(`
	Peer ID: %s
	Listen address: %s
	Topic: %s
	Data Folder: %s
	
	welcome to p2pdb
	
	Commands:
	
	> create               ->  create table 
	> insert <key>         ->  insert data 
	> delete <key> <value> ->  delete data 
	> select               ->  select data
	> update               ->  update data
	`,
		h.ID(), listen, topicName, dbPath,
	)

	if len(os.Args) > 1 && os.Args[1] == "daemon" {
		fmt.Println("Running in daemon mode")
		go func() {
			for {
				fmt.Printf("%s - %d connected peers\n", time.Now().Format(time.Stamp), len(connectedPeers(h)))
				time.Sleep(10 * time.Second)
			}
		}()
		signalChan := make(chan os.Signal, 20)
		signal.Notify(
			signalChan,
			syscall.SIGINT,
			syscall.SIGTERM,
			syscall.SIGHUP,
		)
		<-signalChan
		return
	}

	fmt.Printf("> ")
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		p2pdbSql := scanner.Text()
		fields := strings.Fields(p2pdbSql)
		if len(fields) == 0 {
			fmt.Printf("> ")
			continue
		}
		execute(fields[0], fields, p2pdbSql, ctx, topic, h, true)

		fmt.Printf("> ")
	}
}

func execute(sqlType string, fields []string, p2pdbSql string, ctx context.Context, topic *pubsub.Topic,
	h host.Host, pubMessage bool) {

	switch sqlType {
	case "exit", "quit":
		return
	case "debug":
		if len(fields) < 2 {
			fmt.Println("debug <on/off/peers>")
		}
		st := fields[1]
		switch st {
		case "on":
			logging.SetLogLevel("globaldb", "debug")
		case "off":
			logging.SetLogLevel("globaldb", "error")
		case "peers": //查看对等节点
			for _, p := range connectedPeers(h) {
				addrs, err := peer.AddrInfoToP2pAddrs(p)
				if err != nil {
					logger.Warn(err)
					continue
				}
				for _, a := range addrs {
					fmt.Println(a)
				}
			}
		}
	case "select":
		db, err := sql.Open("sqlite3", dbPath+dbName)
		if err != nil {
			fmt.Println("sql error-> %s", err)
			return
		}
		defer db.Close()
		rows, err := db.Query(p2pdbSql)
		if err != nil {
			fmt.Println("sql error-> %s", err)
			return
		}
		for rows.Next() {
			var id int
			var name string
			err = rows.Scan(&id, &name)
			fmt.Println(id, name)
		}
	case "insert":
		if len(fields) < 4 {
			fmt.Println("sql error->")
			fmt.Println("insert into p2pdb(id, name) values(1, 'foo'), (2, 'bar'), (3, 'baz');")
			return
		}
		db, err := sql.Open("sqlite3", dbPath+dbName)
		if err != nil {
			fmt.Println("sql error-> %s", err)
			return
		}
		defer db.Close()
		_, err = db.Exec(p2pdbSql)
		if err != nil {
			fmt.Println("sql error-> %s"+p2pdbSql, err)
			return
		}
		if pubMessage == true {
			publish(sqlType, p2pdbSql, ctx, topic)
		}
		fmt.Println("sql execute success-> " + p2pdbSql)
	case "create":
		if len(fields) < 2 {
			fmt.Println("sql error->")
			fmt.Println("create table p2pdb (id integer not null, name text);")
			return
		}
		name := fields[1]
		//	v := strings.Join(fields[2:], " ")
		switch name {
		// case "database":
		// 	if len(fields) < 3 {
		// 		fmt.Println("sql error->")
		// 		fmt.Println("create database p2pdb;")
		// 		continue
		// 	}
		// 	os.Remove(dbPath + dbName)
		// 	db, err := sql.Open("sqlite3", dbPath+dbName)
		// 	if err != nil {
		// 		log.Fatal(err)
		// 	}
		// 	defer db.Close()

		// 	fmt.Printf("databse file is exsit in  %s /n", dbPath+dbName)
		case "table":
			if len(fields) < 3 {
				fmt.Println("sql error->")
				fmt.Println("create table p2pdb (id integer not null, name text);")
				fmt.Printf("> ")
				return
			}
			db, err := sql.Open("sqlite3", dbPath+dbName)
			if err != nil {
				log.Println("sql error->%q: %s\n", err, p2pdbSql)
				return
			}
			defer db.Close()
			_, err = db.Exec(p2pdbSql)
			if err != nil {
				log.Println("sql error->%q: %s\n", err, p2pdbSql)
				return
			}
			if pubMessage == true {
				publish(sqlType, p2pdbSql, ctx, topic)
			}
			fmt.Println("sql execute success-> %s", p2pdbSql)
		}

	}

}

// func openDb(name string) *sql.DB {

// 	db, err := sql.Open("sqlite3", name)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	defer db.Close()
// 	return db
// }

// func Exec(sqlStmt string) {
// 	_, err := db.Exec(sqlStmt, db)
// 	if err != nil {
// 		log.Printf("%q: %s\n", err, sqlStmt)
// 		return
// 	}
// }

// func printErr(err error) {
// 	fmt.Println("error:", err)
// 	fmt.Println("> ")
// }

//对等节点连接，返回对等节点信息
func connectedPeers(h host.Host) []*peer.AddrInfo {
	var pinfos []*peer.AddrInfo
	for _, c := range h.Network().Conns() {
		pinfos = append(pinfos, &peer.AddrInfo{
			ID:    c.RemotePeer(),
			Addrs: []multiaddr.Multiaddr{c.RemoteMultiaddr()},
		})
	}
	return pinfos
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
