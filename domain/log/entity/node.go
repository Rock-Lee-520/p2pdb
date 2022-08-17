package entity

type NodeEntity struct {
	NodeId       string     `json:"nodeId"`
	RequestId    string     `json:"requestId"`
	Type         string     `json:"type"`
	LogicalClock int64      `json:"logicalClock"`
	Data         DataEntity `json:"data"`
	Link         LinkEntity `json:"Links"`
}

func NewNodeEntity() *NodeEntity {
	return &NodeEntity{}
}

func (i *NodeEntity) GetNodeId() string {
	return i.NodeId
}

func (i *NodeEntity) SetNodeId(nodeId string) {
	i.NodeId = nodeId
}

func (i *NodeEntity) GetRequestId() string {
	return i.RequestId
}

func (i *NodeEntity) SetRequestId(requestId string) {
	i.RequestId = requestId
}

func (i *NodeEntity) GetType() string {
	return i.Type
}

func (i *NodeEntity) SetType(t string) {
	i.Type = t
}

func (i *NodeEntity) GetLogicalClock() int64 {
	return i.LogicalClock
}

func (i *NodeEntity) SetLogicalClock(logicalClock int64) {
	i.LogicalClock = logicalClock
}
