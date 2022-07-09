package service

import (
	//"log"

	sqle "github.com/Rock-liyi/p2pdb-server"
	"github.com/Rock-liyi/p2pdb-server/auth"
	conf "github.com/Rock-liyi/p2pdb-server/config"
	"github.com/Rock-liyi/p2pdb-server/server"
	"github.com/Rock-liyi/p2pdb-store/sql"
	"github.com/Rock-liyi/p2pdb-store/sql/information_schema"
	log "github.com/sirupsen/logrus"
)

const (
	dbName    = "test"
	tableName = "p2pdb-example"
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

func startNewService() {
	engine := sqle.NewDefault(
		sql.NewDatabaseProvider(
			createMemoryDatabase(), //choose createMemoryDatabase or createSqliteDatabase
			information_schema.NewInformationSchemaDatabase(),
		))

	config := server.Config{
		Protocol: "tcp",
		Address:  "localhost:3306",
		Auth:     auth.NewNativeSingle("root", "", auth.AllPermissions),
	}

	s, err := server.NewDefaultServer(config, engine)
	if err != nil {
		panic(err)
	}

	log.Debug("main function finish =====")
	s.Start()
}
