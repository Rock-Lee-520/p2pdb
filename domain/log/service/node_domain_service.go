package service

import (
	repository "github.com/Rock-liyi/p2pdb/domain/log/repository"
)

func CreateNodeTable(TableName string, databaseName string) {
	repository.CreateNodeTable()
}

func CreateNode(nodeId string, instanceId string, loggicalClock string, peerId string) {

}
