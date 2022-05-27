package main

import (
	"context"
	"fmt"
	"io/ioutil"

	idp "berty.tech/go-ipfs-log/identityprovider"

	"berty.tech/go-ipfs-log/keystore"
	"github.com/ipfs/go-datastore"
	dssync "github.com/ipfs/go-datastore/sync"
	config "github.com/ipfs/go-ipfs-config"
	ipfs_core "github.com/ipfs/go-ipfs/core"
	"github.com/ipfs/go-ipfs/core/coreapi"
	ipfs_libp2p "github.com/ipfs/go-ipfs/core/node/libp2p"
	ipfs_repo "github.com/ipfs/go-ipfs/repo"
	"github.com/libp2p/go-libp2p"
	"github.com/libp2p/go-libp2p-core/host"
	"github.com/libp2p/go-libp2p-core/peer"
	pstore "github.com/libp2p/go-libp2p-peerstore"

	log "berty.tech/go-ipfs-log"
)

var (
	dbType = "orbitdb"
)

func buildHostOverrideExample(ctx context.Context, id peer.ID, ps pstore.Peerstore, options ...libp2p.Option) (host.Host, error) {
	return ipfs_libp2p.DefaultHostOption(ctx, id, ps, options...)
}

func newRepo() (ipfs_repo.Repo, error) {
	// Generating config
	cfg, err := config.Init(ioutil.Discard, 2048)
	if err != nil {
		return nil, err
	}

	// Listen on local interface only
	cfg.Addresses.Swarm = []string{
		"/ip4/127.0.0.1/tcp/0",
	}

	// Do not bootstrap on ipfs node
	cfg.Bootstrap = []string{
		"/ip4/0.0.0.0/tcp/33123/ipfs/12D3KooWMVdnQXeh97noZrUavoULs7GA2qQYhHTFRueDAmyprRaH",
		//"/ip4/10.209.144.6/tcp/4001/p2p/12D3KooWRtdFu6pJ9kTXTAxzCsmGdoPGzRdm5akw2Bp4wzTz14Gr",
		//"/ip4/10.209.144.5/tcp/4001/p2p/12D3KooWR2Ab6nMFSdAtKEHeoFgJe2TepnCSu6fJforMBLyGubk6",
	}

	return &ipfs_repo.Mock{
		D: dssync.MutexWrap(datastore.NewMapDatastore()),
		C: *cfg,
	}, nil
}

func newRepo2() (ipfs_repo.Repo, error) {
	// Generating config
	cfg, err := config.Init(ioutil.Discard, 2048)
	if err != nil {
		return nil, err
	}

	// Listen on local interface only

	cfg.Addresses.Swarm = []string{
		"/ip4/127.0.0.1/tcp/0",
	}

	// Do not bootstrap on ipfs node
	cfg.Bootstrap = []string{
		"/ip4/0.0.0.0/tcp/33123/ipfs/12D3KooWMVdnQXeh97noZrUavoULs7GA2qQYhHTFRueDAmyprRaH",
		//"/ip4/10.209.144.6/tcp/4001/p2p/12D3KooWRtdFu6pJ9kTXTAxzCsmGdoPGzRdm5akw2Bp4wzTz14Gr",
		//"/ip4/10.209.144.5/tcp/4001/p2p/12D3KooWR2Ab6nMFSdAtKEHeoFgJe2TepnCSu6fJforMBLyGubk6",
	}

	return &ipfs_repo.Mock{
		D: dssync.MutexWrap(datastore.NewMapDatastore()),
		C: *cfg,
	}, nil
}

func buildNode(ctx context.Context) *ipfs_core.IpfsNode {
	r, err := newRepo()
	if err != nil {
		panic(err)
	}

	cfg := &ipfs_core.BuildCfg{
		Online: true,
		Repo:   r,
		Host:   buildHostOverrideExample,
	}

	nodeA, err := ipfs_core.NewNode(ctx, cfg)
	if err != nil {
		panic(err)
	}
	return nodeA
}

// func buildNode2(ctx context.Context) *ipfs_core.IpfsNode {
// 	r, err := newRepo2()
// 	if err != nil {
// 		panic(err)
// 	}

// 	cfg := &ipfs_core.BuildCfg{
// 		Online: true,
// 		Repo:   r,
// 		Host:   buildHostOverrideExample,
// 	}

// 	nodeA, err := ipfs_core.NewNode(ctx, cfg)
// 	if err != nil {
// 		panic(err)
// 	}
// 	return nodeA
// }

// func buildNode3(ctx context.Context) *ipfs_core.IpfsNode {
// 	r, err := newRepo2()
// 	if err != nil {
// 		panic(err)
// 	}

// 	cfg := &ipfs_core.BuildCfg{
// 		Online: true,
// 		Repo:   r,
// 		Host:   buildHostOverrideExample,
// 	}

// 	nodeA, err := ipfs_core.NewNode(ctx, cfg)
// 	if err != nil {
// 		panic(err)
// 	}
// 	return nodeA
// }

func main() {
	fmt.Println("test log merge")
	//初始化上下文
	ctx := context.Background()
	// Build Ipfs Node A and Node B
	nodeA := buildNode(ctx)
	nodeB := buildNode(ctx)
	//nodeC := buildNode(ctx)
	// Fill up datastore with identities
	ds := dssync.MutexWrap(datastore.NewMapDatastore())
	ks, err := keystore.NewKeystore(ds)
	if err != nil {
		panic(err)
	}
	// Create identity A
	identityA, err := idp.CreateIdentity(&idp.CreateIdentityOptions{
		Keystore: ks,
		ID:       "dbA",
		Type:     dbType,
	})
	if err != nil {
		panic(fmt.Errorf("coreapi error: %s", err))
	}

	identityB, err := idp.CreateIdentity(&idp.CreateIdentityOptions{
		Keystore: ks,
		ID:       "dbB",
		Type:     dbType,
	})
	if err != nil {
		panic(fmt.Errorf("coreapi error: %s", err))
	}

	// identityC, err := idp.CreateIdentity(&idp.CreateIdentityOptions{
	// 	Keystore: ks,
	// 	ID:       "dbC",
	// 	Type:     dbType,
	// })
	// if err != nil {
	// 	panic(fmt.Errorf("coreapi error: %s", err))
	// }

	serviceA, err := coreapi.NewCoreAPI(nodeA)
	if err != nil {
		panic(fmt.Errorf("coreapi error: %s", err))
	}

	serviceB, err := coreapi.NewCoreAPI(nodeB)
	if err != nil {
		panic(fmt.Errorf("coreapi error: %s", err))
	}

	// serviceC, err := coreapi.NewCoreAPI(nodeC)
	// if err != nil {
	// 	panic(fmt.Errorf("coreapi error: %s", err))
	// }

	// creating log
	log1, err := log.NewLog(serviceA, identityA, &log.LogOptions{ID: "A"})
	if err != nil {
		panic(err)
	}

	// creating log
	log2, err := log.NewLog(serviceB, identityB, &log.LogOptions{ID: "A"})
	if err != nil {
		panic(err)
	}

	// creating log
	// log3, err := log.NewLog(serviceC, identityC, &log.LogOptions{ID: "A"})
	// if err != nil {
	// 	panic(err)
	// }

	_, err = log1.Append(ctx, []byte("Aone"), nil)
	if err != nil {
		panic(fmt.Errorf("append error: %s", err))
	}
	_, err = log2.Append(ctx, []byte("Bone"), nil)
	if err != nil {
		panic(fmt.Errorf("append error: %s", err))
	}

	_, err = log1.Append(ctx, []byte("Atwo"), nil)
	if err != nil {
		panic(fmt.Errorf("append error: %s", err))
	}
	_, err = log2.Append(ctx, []byte("Btwo"), nil)
	if err != nil {
		panic(fmt.Errorf("append error: %s", err))
	}

	// Join the logs
	//log3.Join(log1, 1)
	log1.Join(log2, 2)

	_, err = log1.Append(ctx, []byte("four"), nil)
	if err != nil {
		panic(fmt.Errorf("append error: %s", err))
	}

	fmt.Println(log1.Values())

	h, err := log1.ToMultihash(ctx)
	if err != nil {
		panic(fmt.Errorf("ToMultihash error: %s", err))
	}

	res, err := log.NewFromMultihash(ctx, serviceA, identityA, h, &log.LogOptions{}, &log.FetchOptions{})
	if err != nil {
		panic(fmt.Errorf("NewFromMultihash error: %s", err))
	}
	fmt.Println("echo result====")
	// nodeB lookup logA
	fmt.Println(res.ToString(nil))

}
