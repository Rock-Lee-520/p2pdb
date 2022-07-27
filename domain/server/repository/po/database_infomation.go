package persistence

// Database  information_schema definition
type DatabaseInfomation struct {
	BaseColumn
	DatabaseName    string `gorm:"column:database_name"`
	LocalPeerId     string `gorm:"column:local_peer_id"`
	GlobalClockTime int64  `gorm:"column:global_clock_time"`
}
