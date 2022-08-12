package entity

type Node struct {
	NodeId       string   `json:"nodeId"`
	RequestId    string   `json:"requestId"`
	Type         string   `json:"type"`
	LogicalClock int      `json:"logicalClock"`
	LastNodeId   string   `json:"lastNodeId"`
	Data         Data     `json:"data"`
	Link         []string `json:"Links"`
}

type Data struct {
	DataId     string   `json:"dataId"`
	Operation  string   `json:"operation"`
	Properties []string `json:"properties"`
}

type Link struct {
	LinkId     string `json:"linkId"`
	LastNodeId string `json:"lastNodeId"`
	NodeId     string `json:"nodeId"`
	TableId    string `json:"tableId"`
	DatabaseId string `json:"databaseId"`
	InstanceId string `json:"instanceId"`
}
