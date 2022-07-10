package service

import (
	server "github.com/Rock-liyi/p2pdb/infrastructure/core/server"
	config "github.com/Rock-liyi/p2pdb/infrastructure/util/config"
)

func StartNewService() {
	DBHost := config.GetDBHost()
	if DBHost == "" {
		DBHost = "127.0.0.1"
	}

	port := config.GetPort()
	if port == "" {
		port = "3306"
	}

	DBName := config.GetDBName()
	if DBName == "" {
		DBName = "p2pdb"
	}

	DBTableName := config.GetDBTableName()
	if DBTableName == "" {
		DBTableName = "test"
	}

	server.StartNewService(DBName, DBTableName, DBHost, port)
}
