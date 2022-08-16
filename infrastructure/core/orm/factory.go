package store

import (
	"os"
	"path/filepath"

	conf "github.com/Rock-liyi/p2pdb/infrastructure/util/config"
	"github.com/Rock-liyi/p2pdb/infrastructure/util/log"
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
	Update(column string, value interface{}) *gorm.DB
	Delete(value interface{}, conds ...interface{}) *gorm.DB
	Select(query interface{}, args ...interface{}) *gorm.DB
	Where(query interface{}, args ...interface{}) *gorm.DB
	First(out interface{}, where ...interface{}) *gorm.DB
	Find(out interface{}, where ...interface{}) *gorm.DB
	Connect()
	// HasTable(dst ...interface{}) bool
	// CreateTable(dst ...interface{}) error
	//AutoMigrate(dst ...interface{}) error
	Exec(sql string, values ...interface{}) *gorm.DB
	Raw(sql string, values ...interface{}) *gorm.DB
	Table(name string, args ...interface{}) *gorm.DB
	Migrator() gorm.Migrator
	InsertIgnore(value interface{}) error
}

type CreateDBFactory struct {
	IsInternalStore bool
	DatabaseName    string
}

func (db *CreateDBFactory) SetIsInternalStore(value bool) {
	db.IsInternalStore = value
}

func (db *CreateDBFactory) SetDatabaseName(databaseName string) {
	db.DatabaseName = databaseName
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

	var dataName string
	var dataPath string

	//Is it a internal  database?
	if db.IsInternalStore == true {
		dataName = conf.GetDBInformationName()
		//init config,get db path
		dataPath = conf.GetInternaleDataPath()
		if dataPath != "" {
			dataPath = dataPath + "/"
		}
	} else {
		dataName = conf.GetDBName()
		//init config,get db path
		dataPath = conf.GetDataPath()
		if dataPath != "" {
			dataPath = dataPath + "/"
		}
	}

	//Whether the database name is set dynamically?
	if db.DatabaseName != "" {
		dataName = db.DatabaseName
	}

	if dataName == "" {
		log.Panic("dbname does not exits in .env file")
	}

	binary, _ := os.Getwd()
	root := filepath.Dir(binary)
	if root != "" && dataPath == "" {
		dataPath = root + "/"
	}

	address := dataPath + dataName + ".db"

	log.Info("database address is " + address)
	connect.Init(address, 0, "", "")
	connect.Connect()
	return connect

}
