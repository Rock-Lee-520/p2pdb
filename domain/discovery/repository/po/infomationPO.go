package persistence

// Database  information_schema definition
type InstanceInformation struct {
	BaseColumn
	InstanceId      string `gorm:"column:instance_id;primary_key"`
	LocalPeerId     string `gorm:"column:local_peer_id"`
	GlobalClockTime int64  `gorm:"column:global_clock_time"`
}

// Database  information_schema definition
type PeerNodeInfomation struct {
	BaseColumn
	PeerNodeId         string `gorm:"column:peer_node_id;primary_key"`
	DatabaseId         string `gorm:"column:database_id"`
	TableId            string `gorm:"column:table_id"`
	LocalPeerId        string `gorm:"column:local_peer_id"`
	LocalInstanceId    string `gorm:"column:local_instance_id"`
	LocalHost          string `gorm:"column:local_host"`
	RemoteHost         string `gorm:"column:remote_host"`
	RemotePeerId       string `gorm:"column:remote_peer_id"`
	RemoteInstanceId   string `gorm:"column:remote_instance_id"`
	LastDagMergeNodeId string `gorm:"column:last_node_id"`
	ProtocolId         string `gorm:"column:protocol_id"` //defualt /ipfs/id/1.0.0
}
