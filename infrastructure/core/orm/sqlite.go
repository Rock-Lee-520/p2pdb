package store

import (
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
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
	//echo all of sql
	// logger.Default.LogMode(logger.Info)

	// newLogger := logger.New(
	// 	log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
	// 	logger.Config{
	// 		SlowThreshold:             time.Second,   // Slow SQL threshold
	// 		LogLevel:                  logger.Silent, // Log level
	// 		IgnoreRecordNotFoundError: true,          // Ignore ErrRecordNotFound error for logger
	// 		Colorful:                  true,          // Disable color
	// 	},
	// )
	//logger.Default.LogMode(logger.Info)

	ormDB, err := gorm.Open(sqlite.Open("/Users/rockli/go/src/p2pdb-server/data/test.db"), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		panic("failed to connect database")
	}

	// ormDB.LogMode(true)
	// ormDB.SingularTable(true)
	// ormDB.SetLogger(log.New(os.Stdout, "\r\n", 0))
	db.OrmDB = ormDB
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
	return db.OrmDB.Migrator()
}

func (db *SqliteDB) InsertIgnore(value interface{}) error {
	//return nil
	return db.OrmDB.Clauses(clause.Insert{Modifier: " OR IGNORE"}).Create(value).Error
}

//something wrong
// func (db *SqliteDB) CreateTable(dst ...interface{}) error {
// 	return db.OrmDB.Migrator().CreateTable(dst)
// }

// func (db *SqliteDB) HasTable(dst ...interface{}) bool {
// 	return db.OrmDB.Migrator().HasTable(dst)
// }
