package po

// Node  information_schema definition
type Node struct {
	BaseColumn
	NodeId       string `gorm:"column:node_id"`
	RequestId    string `gorm:"column:request_id"`
	Type         string `gorm:"column:type"`
	LogicalClock string `gorm:"column:logcal_clock"`
	LastNodeId   string `gorm:"column:last_node_Id"`
}

// Data  information_schema definition
type Data struct {
	BaseColumn
	DataId     string `gorm:"column:data_id"`
	NodeId     string `gorm:"column:node_id"`
	Operation  string `gorm:"column:operation"`
	Properties string `gorm:"column:properties"`
}

// Link  information_schema definition
type Link struct {
	BaseColumn
	LinkId     string `gorm:"column:link_id"`
	LastNodeId string `gorm:"column:last_node_id"`
	NodeId     string `gorm:"column:node_id"`
	TableId    string `gorm:"column:table_id;primary_key"`
	DatabaseId string `gorm:"column:database_id"`
	InstanceId string `gorm:"column:local_instance_id"`
}
