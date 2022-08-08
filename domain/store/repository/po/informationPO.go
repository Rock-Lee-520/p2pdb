package po

// Database  information_schema definition
type DatabaseInfomation struct {
	BaseColumn
	DatabaseId      string `gorm:"column:database_id"`
	DatabaseName    string `gorm:"column:database_name"`
	TableId         string `gorm:"column:table_id"`
	LocalPeerId     string `gorm:"column:local_peer_id"`
	LocalInstanceId string `gorm:"column:local_instance_id"`
}

// Table  information_schema definition
type TableInfomation struct {
	BaseColumn
	TableId         string `gorm:"column:table_id"`
	DatabaseId      string `gorm:"column:database_id"`
	TableName       string `gorm:"column:table_name"`
	LocalPeerId     string `gorm:"column:local_peer_id"`
	LocalInstanceId string `gorm:"column:local_instance_id"`
	ClockTime       int64  `gorm:"column:clock_time"`
	Version         int64  `gorm:"column:version"`
	Schema          int64  `gorm:"column:schema"`
}
