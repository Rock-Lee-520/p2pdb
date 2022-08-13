package entity

type DataEntity struct {
	DataId     string   `json:"dataId"`
	Operation  string   `json:"operation"` //create  update delete
	Properties []string `json:"properties"`
}
