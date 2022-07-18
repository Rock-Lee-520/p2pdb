package persistence

// Database  information_schema definition
type SyncInfomation struct {
	BaseColumn
	DatabaseName string `gorm:"column:database_name"`
	TableName    string `gorm:"column:table_name"`
	LocalPeerId  string `gorm:"column:local_peer_id"`
	RemotePeerId string `gorm:"column:remote_peer_id"`
	LastNodeId   string `gorm:"column:last_node_id"`
}
