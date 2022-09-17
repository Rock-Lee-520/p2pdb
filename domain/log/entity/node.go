package entity

type Node struct {
	NodeId       string `json:"nodeId"`
	RequestId    string `json:"requestId"`
	Type         string `json:"type"`
	LogicalClock int64  `json:"logicalClock"`
	LastNodeId   string `json:"lastNodeId"`
	Data         Data   `json:"data"`
	Link         []Link   `json:"Links"`
}

func NewNode() *Node {
	return &Node{}
}

func (i *Node) GetNodeId() string {
	return i.NodeId
}

func (i *Node) SetNodeId(nodeId string) {
	i.NodeId = nodeId
}

func (i *Node) GetLastNodeId() string {
	return i.LastNodeId
}

func (i *Node) SetLastNodeId(lastNodeId string) {
	i.LastNodeId = lastNodeId
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
