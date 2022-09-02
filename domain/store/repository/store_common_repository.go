package repository

import (
	orm "github.com/Rock-liyi/p2pdb/infrastructure/core/orm"
)

var ORM *orm.CreateDBFactory = new(orm.CreateDBFactory)
var db orm.DBconnect

func InitDB() orm.DBconnect {
	//settings  use  the default  information database
	ORM.SetIsInternalStore(true)
	db = ORM.InitDB()
	return db
}
