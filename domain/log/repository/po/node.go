package po

// Node  information_schema definition
type Node struct {
	BaseColumn
	NodeId       string `gorm:"column:node_id"`
	RequestId    string `gorm:"column:request_id"`
	Type         string `gorm:"column:type"`
	LogicalClock int64  `gorm:"column:logcal_clock"`
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

func (i *Node) GetLastNodeId() string {
	return i.LastNodeId
}

func (i *Node) SetLastNodeId(lastNodeId string) {
	i.LastNodeId = lastNodeId
}

func (i *Node) GetNodeId() string {
	return i.NodeId
}

func (i *Node) SetNodeId(nodeId string) {
	i.NodeId = nodeId
}

func (i *Node) GetRequestId() string {
	return i.RequestId
}

func (i *Node) SetRequestId(requestId string) {
	i.RequestId = requestId
}

func (i *Node) GetType() string {
	return i.Type
}

func (i *Node) SetType(t string) {
	i.Type = t
}

func (i *Node) GetLogicalClock() int64 {
	return i.LogicalClock
}

func (i *Node) SetLogicalClock(logicalClock int64) {
	i.LogicalClock = logicalClock
}
