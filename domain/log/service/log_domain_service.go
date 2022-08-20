package service

import (
	"github.com/Rock-liyi/p2pdb/domain/log/entity"
	"github.com/Rock-liyi/p2pdb/domain/log/repository"
	PO "github.com/Rock-liyi/p2pdb/domain/log/repository/po"
	"github.com/Rock-liyi/p2pdb/infrastructure/util/function"
)

func InitLogTable(tableName string, databaseName string) {
	CreateNodeTable(tableName, databaseName)
	CreateDataTable(tableName, databaseName)
	CreateLinkTable(tableName, databaseName)
}

func CreateNodeTable(tableName string, databaseName string) {
	tableName = "node_" + tableName
	repository.CreateNodeTable(tableName, databaseName)

}

func CreateDataTable(tableName string, databaseName string) {
	tableName = "data_" + tableName
	repository.CreateDataTable(tableName, databaseName)

}

func CreateLinkTable(tableName string, databaseName string) {
	tableName = "link_" + tableName
	repository.CreateLinkTable(tableName, databaseName)

}

func InsertDataLog(data *entity.Data, db *entity.DB) {
	var DataPO = PO.NewData()
	DataPO.SetDataId(data.GetDataId())
	DataPO.SetNodeId(data.GetNodeId())
	DataPO.SetOperation(data.GetOperation())
	var properties = function.JsonEncode(data.GetProperties())
	DataPO.SetProperties(function.ToString(properties))
	tableName := "data_" + db.GetTableName()
	repository.InsertDataLog(DataPO, tableName, db.GetTableName())
}

func InsertNodeLog(node *entity.Node, db *entity.DB) {
	var NodePO = PO.NewNode()
	NodePO.SetNodeId(node.GetNodeId())
	NodePO.SetRequestId(node.GetRequestId())
	NodePO.SetLogicalClock(node.GetLogicalClock())
	NodePO.SetLastNodeId(node.GetLastNodeId())
	NodePO.SetType(node.GetType())
	tableName := "node_" + db.GetTableName()
	repository.InsertNodeLog(NodePO, tableName, db.GetTableName())
}

func InsertLinkLog(link *entity.Link, db *entity.DB) {
	var LinkPO = PO.NewLink()
	LinkPO.SetDatabaseId(link.GetDatabaseId())
	LinkPO.SetLastNodeId(link.GetLastNodeId())
	LinkPO.SetInstanceId(link.GetInstanceId())
	LinkPO.SetLinkId(link.GetLinkId())
	LinkPO.SetNodeId(link.GetNodeId())
	LinkPO.SetTableId(link.GetTableId())
	tableName := "link_" + db.GetTableName()
	repository.InsertLinkLog(LinkPO, tableName, db.GetTableName())
}
