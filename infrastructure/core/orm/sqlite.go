package store

import (
	"log"
	"os"
	"time"

	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type SqliteDB struct {
	BaseInfo
	OrmDB *gorm.DB
}

func (this *BaseInfo) Init(address string, port int64, account string, passwd string) {
	this.address = address
	this.port = port
	this.account = account
	this.passwd = passwd
}

func (db *SqliteDB) Connect() {
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold:             time.Second,   // Slow SQL threshold
			LogLevel:                  logger.Silent, // Log level
			IgnoreRecordNotFoundError: true,          // Ignore ErrRecordNotFound error for logger
			Colorful:                  false,         // Disable color
		},
	)

	ormDB, err := gorm.Open(sqlite.Open("/Users/rockli/go/src/p2pdb-server/data/test.db"), &gorm.Config{
		Logger: newLogger,
	})
	if err != nil {
		panic("failed to connect database")
	}
	db.OrmDB = ormDB
	// ormDB.Logger
	// ormDB.LogMode(true)
	// ormDB.SingularTable(true)
	//rmDB.SetLogger(log.New(os.Stdout, "\r\n", 0))
}

func (db *SqliteDB) Create(value interface{}) *gorm.DB {

	return db.OrmDB.Create(value)
}
func (db *SqliteDB) Updates(value interface{}) *gorm.DB {
	return db.OrmDB.Updates(value)
}
func (db *SqliteDB) Delete(value interface{}, conds ...interface{}) *gorm.DB {
	return db.OrmDB.Delete(value, conds)
}

func (db *SqliteDB) Select(query interface{}, args ...interface{}) *gorm.DB {
	return db.OrmDB.Select(query, args)
}

func (db *SqliteDB) First(out interface{}, where ...interface{}) *gorm.DB {
	return db.OrmDB.First(out, where)
}

func (db *SqliteDB) Find(out interface{}, where ...interface{}) *gorm.DB {
	return db.OrmDB.Find(out, where)
}

func (db *SqliteDB) Where(query interface{}, args ...interface{}) *gorm.DB {
	return db.OrmDB.Where(query, args)
}

// func (db *SqliteDB) AutoMigrate(dst ...interface{}) error {
// 	type Product struct {
// 		gorm.Model
// 		Code  string
// 		Price uint
// 	}
// 	return db.OrmDB.Migrator().AutoMigrate(&Product{})
// }

func (db *SqliteDB) Migrator() gorm.Migrator {
	type Product struct {
		gorm.Model
		Code  string
		Price uint
	}
	return db.OrmDB.Migrator()
}

//something wrong
// func (db *SqliteDB) CreateTable(dst ...interface{}) error {
// 	return db.OrmDB.Migrator().CreateTable(dst)
// }

// func (db *SqliteDB) HasTable(dst ...interface{}) bool {
// 	return db.OrmDB.Migrator().HasTable(dst)
// }
