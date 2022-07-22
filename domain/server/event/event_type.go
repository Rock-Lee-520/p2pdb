package event

type NodeMessage struct {
	NodeId    string            `json:"node_id"`
	PeerId    string            `json:"peer_id"`
	ClockTime string            `json:"clock_time"`
	Data      []NodeMessageData `json:"data"`
	Link      []NodeMessageLink `json:"link"`
}

type NodeMessageData struct {
}

type NodeMessageLink struct {
}
