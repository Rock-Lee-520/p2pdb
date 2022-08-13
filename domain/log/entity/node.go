package entity

type NodeEntity struct {
	NodeId       string     `json:"nodeId"`
	RequestId    string     `json:"requestId"`
	Type         string     `json:"type"`
	LogicalClock int        `json:"logicalClock"`
	Data         DataEntity `json:"data"`
	Link         LinkEntity `json:"Links"`
}

func NewNodeEntity() *NodeEntity {
	return &NodeEntity{}
}
