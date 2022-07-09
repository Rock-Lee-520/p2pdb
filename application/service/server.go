package service

import (
	server "github.com/Rock-liyi/p2pdb/infrastructure/core/server"
)

const (
	dbName    = "test"
	tableName = "p2pdb-example"
)

func StartNewService() {
	server.StartNewService(dbName, tableName)
}
