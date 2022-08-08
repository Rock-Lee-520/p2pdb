package po

// Database  information_schema definition
type DatabaseInfomation struct {
	BaseColumn
	DatabaseId      string `gorm:"column:database_id;primary_key"`
	DatabaseName    string `gorm:"column:database_name"`
	LocalInstanceId string `gorm:"column:local_instance_id"`
}

// Table  information_schema definition
type TableInfomation struct {
	BaseColumn
	TableId         string `gorm:"column:table_id;primary_key"`
	DatabaseId      string `gorm:"column:database_id"`
	TableName       string `gorm:"column:table_name"`
	LocalInstanceId string `gorm:"column:local_instance_id"`
	LogicalClock    int64  `gorm:"column:clock_time"`
	Version         int64  `gorm:"column:version"`
	Schema          int64  `gorm:"column:schema"`
}
