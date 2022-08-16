package service

import "testing"

func TestPublish(t *testing.T) {
	tableName := "test4"
	database := "test5"
	InitLogTable(tableName, database)
}
