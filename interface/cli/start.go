package main

import (
	_ "github.com/Rock-liyi/p2pdb/application/event/subscribe" //注册事件监听
	"github.com/Rock-liyi/p2pdb/application/service"
	conf "github.com/Rock-liyi/p2pdb/infrastructure/util/config"
	log "github.com/sirupsen/logrus"
)

func init() {
	environment := conf.GetEnv()
	// do something here to set environment depending on an environment variable
	// or command-line flag
	if environment == "production" {
		log.SetLevel(log.InfoLevel)
	} else {
		// The TextFormatter is default, you don't actually have to do this.
		log.SetLevel(log.DebugLevel)
	}

}

func main() {

	//start a discovery server
	service.InitDiscovery()

	//start a mysql server
	service.StartNewService()
}
