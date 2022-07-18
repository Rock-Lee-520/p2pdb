package store

import (
	"os"
	"path/filepath"

	conf "github.com/Rock-liyi/p2pdb-log/config"
	debug "github.com/favframework/debug"
	"gorm.io/gorm"
)

type BaseInfo struct {
	address string
	port    int64
	account string
	passwd  string
}

//Connect()应该返回的是连接对象,这里做了省略,所以用string来代替

type DBconnect interface {
	Init(address string, port int64, account string, passwd string)
	Create(value interface{}) *gorm.DB
	Updates(value interface{}) *gorm.DB
	Delete(value interface{}, conds ...interface{}) *gorm.DB
	Select(query interface{}, args ...interface{}) *gorm.DB
	Where(query interface{}, args ...interface{}) *gorm.DB
	First(out interface{}, where ...interface{}) *gorm.DB
	Find(out interface{}, where ...interface{}) *gorm.DB
	Connect()
	// HasTable(dst ...interface{}) bool
	// CreateTable(dst ...interface{}) error
	//AutoMigrate(dst ...interface{}) error
	Migrator() gorm.Migrator
}

type CreateDBFactory struct {
}

func (db *CreateDBFactory) CreateDBConnect(db_type string) DBconnect {
	var orm DBconnect
	switch db_type {
	case "mysql":
		orm = &SqliteDB{}
	default:
		orm = &SqliteDB{}
	}
	return orm
}

func (db *CreateDBFactory) InitDB() DBconnect {
	var connect = db.CreateDBConnect("sqlite")
	//init config,get db path
	dataPath := conf.GetDataPath()
	if dataPath != "" {
		dataPath = dataPath + "/"
	}

	//get db name
	dataName := conf.GetDBName()
	debug.Dump(dataName)

	if dataName == "" {
		panic("dbname does not exits in .env file")
	}

	binary, _ := os.Getwd()
	root := filepath.Dir(binary)
	if root != "" && dataPath == "" {
		dataPath = root + "/"
	}

	address := dataPath + dataName + ".db"
	debug.Dump(address)
	connect.Init(address, 0, "", "")
	connect.Connect()

	return connect

}
