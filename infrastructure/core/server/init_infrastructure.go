package server

import (
	log "github.com/Rock-liyi/p2pdb/infrastructure/util/log"

	sqle "github.com/Rock-liyi/p2pdb-server"
	"github.com/Rock-liyi/p2pdb-server/auth"
	ser "github.com/Rock-liyi/p2pdb-server/server"
	"github.com/Rock-liyi/p2pdb-store/memory"
	"github.com/Rock-liyi/p2pdb-store/sql"
	"github.com/Rock-liyi/p2pdb-store/sql/information_schema"
	"github.com/Rock-liyi/p2pdb-store/sqlite"
)

type engine interface {
	CreateSqliteDatabase(DBName string, tableName string) *sqlite.Database

	CreateMemoryDatabase(DBName string, tableName string) *memory.Database
}

func StartNewService(DBName string, tableName string, DBHost string, port string, storageEngine string) {

	var engine *sqle.Engine

	if storageEngine == "file" {
		engine = sqle.NewDefault(
			sql.NewDatabaseProvider(
				CreateSqliteDatabase(DBName, tableName), //choose createMemoryDatabase or createSqliteDatabase
				information_schema.NewInformationSchemaDatabase(),
			))
	} else {
		engine = sqle.NewDefault(
			sql.NewDatabaseProvider(
				CreateMemoryDatabase(DBName, tableName), //choose createMemoryDatabase or createSqliteDatabase
				information_schema.NewInformationSchemaDatabase(),
			))
	}

	config := ser.Config{
		Protocol: "tcp",
		Address:  DBHost + ":" + port,
		Auth:     auth.NewNativeSingle("root", "", auth.AllPermissions),
	}

	log.Info("The server host is: " + DBHost + ",the server port is: " + port)
	s, err := ser.NewDefaultServer(config, engine)
	if err != nil {
		panic(err)
	}
	s.Start()
}
