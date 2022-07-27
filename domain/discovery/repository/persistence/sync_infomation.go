package persistence

// Database  information_schema definition
type SyncInfomation struct {
	BaseColumn
	DatabaseName string `gorm:"column:database_name"`
	TableName    string `gorm:"column:table_name"`
	LocalPeerId  string `gorm:"column:local_peer_id"`
	LocalHost    string `gorm:"column:local_host"`
	RemoteHost   string `gorm:"column:remote_host"`
	RemotePeerId string `gorm:"column:remote_peer_id"`
	LastNodeId   string `gorm:"column:last_node_id"`
	ProtocolId   string `gorm:"column:protocol_id"` //defualt /ipfs/id/1.0.0
}
