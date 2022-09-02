package po

// Database  information_schema definition
type DatabaseInfomation struct {
	BaseColumn
	DatabaseId      string `gorm:"column:database_id;primary_key"`
	DatabaseName    string `gorm:"column:database_name"`
	LogicalClock    int64  `gorm:"column:logical_clock"`
	LocalInstanceId string `gorm:"column:local_instance_id"`
}

func (instance *DatabaseInfomation) GetDatabasePrimary() string {
	return "database_id"
}

func (database *DatabaseInfomation) GetDatabaseName() string {
	return database.DatabaseName
}

func (database *DatabaseInfomation) SetDatabaseName(databaseName string) {
	database.DatabaseName = databaseName
}

func (database *DatabaseInfomation) GetDatabaseId() string {
	return database.DatabaseId
}

func (database *DatabaseInfomation) SetDatabaseId(databaseId string) {
	database.DatabaseId = databaseId
}

func (database *DatabaseInfomation) SetLocalInstanceId(localInstanceId string) {
	database.LocalInstanceId = localInstanceId
}

func (database *DatabaseInfomation) GetLocalInstanceId() string {
	return database.LocalInstanceId
}

func (database *DatabaseInfomation) GetLogicalClock() int64 {
	return database.LogicalClock
}

func (database *DatabaseInfomation) SetLogicalClock(logicalClock int64) {
	database.LogicalClock = logicalClock
}
