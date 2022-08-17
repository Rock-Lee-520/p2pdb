package service

import (
	"github.com/Rock-liyi/p2pdb/domain/log/entity"
	"github.com/Rock-liyi/p2pdb/domain/log/repository"
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

func InsertDataLog(data *entity.DataEntity) {

}

func InsertNodeLog(node *entity.NodeEntity) {

}

func InsertLinkLog(link *entity.LinkEntity) {

}
