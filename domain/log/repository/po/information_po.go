package po

/*
* TODO maybe it will caused variable  value mixed with others
 */
var NODE_TABLE_NAME string = "node"
var DATA_TABLE_NAME string = "data"
var LINK_TABLE_NAME string = "link"

// Node  information_schema definition
type Node struct {
	BaseColumn
	NodeId       string `gorm:"column:node_id"`
	RequestId    string `gorm:"column:request_id"`
	Type         string `gorm:"column:type"`
	LogicalClock string `gorm:"column:logcal_clock"`
	LastNodeId   string `gorm:"column:last_node_Id"`
	TabName      string `gorm:"-"`
}

func NewNode() *Node {
	return &Node{}
}

func (n *Node) SetTabName(TabName string) {
	//n.TabName = TabName
	NODE_TABLE_NAME = TabName
}

func (n *Node) TableName() string {
	return NODE_TABLE_NAME
}

// Data  information_schema definition
type Data struct {
	BaseColumn
	DataId     string `gorm:"column:data_id"`
	NodeId     string `gorm:"column:node_id"`
	Operation  string `gorm:"column:operation"`
	Properties string `gorm:"column:properties"`
	TabName    string `gorm:"-"`
}

func NewData() *Data {
	return &Data{}
}

func (n *Data) SetTabName(TabName string) {
	//n.TabName = TabName
	DATA_TABLE_NAME = TabName
}

func (n *Data) TableName() string {
	return DATA_TABLE_NAME
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
	TabName    string `gorm:"-"`
}

func NewLink() *Link {
	return &Link{}
}

func (n *Link) SetTabName(TabName string) {
	LINK_TABLE_NAME = TabName
}

func (n *Link) TableName() string {
	return LINK_TABLE_NAME
}
