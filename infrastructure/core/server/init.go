package server

import (
	//"log"

	sqle "github.com/Rock-liyi/p2pdb-server"
	"github.com/Rock-liyi/p2pdb-server/auth"
	ser "github.com/Rock-liyi/p2pdb-server/server"
	"github.com/Rock-liyi/p2pdb-store/memory"
	"github.com/Rock-liyi/p2pdb-store/sql"
	"github.com/Rock-liyi/p2pdb-store/sql/information_schema"
	"github.com/Rock-liyi/p2pdb-store/sqlite"
)

type engine interface {
	CreateSqliteDatabase(dbName string, tableName string) *sqlite.Database

	CreateMemoryDatabase(dbName string, tableName string) *memory.Database
}

func StartNewService(dbName string, tableName string) {

	engine := sqle.NewDefault(
		sql.NewDatabaseProvider(
			CreateMemoryDatabase(dbName, tableName), //choose createMemoryDatabase or createSqliteDatabase
			information_schema.NewInformationSchemaDatabase(),
		))

	config := ser.Config{
		Protocol: "tcp",
		Address:  "localhost:3306",
		Auth:     auth.NewNativeSingle("root", "", auth.AllPermissions),
	}

	s, err := ser.NewDefaultServer(config, engine)
	if err != nil {
		panic(err)
	}
	s.Start()
}
