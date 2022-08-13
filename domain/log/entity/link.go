package entity

type LinkEntity struct {
	LinkId     string `json:"linkId"`
	LastNodeId string `json:"lastNodeId"`
	NodeId     string `json:"nodeId"`
	TableId    string `json:"tableId"`
	DatabaseId string `json:"databaseId"`
	InstanceId string `json:"instanceId"`
}
