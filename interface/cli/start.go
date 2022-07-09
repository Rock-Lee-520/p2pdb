package main

import (
	conf "github.com/Rock-liyi/p2pdb-server/config"
	"github.com/Rock-liyi/p2pdb/application/service"
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

	service.StartNewService()
}
