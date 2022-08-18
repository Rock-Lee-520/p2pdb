package repository

import (
	"testing"
)

func TestCreateTable(t *testing.T) {
	CreateNodeTable("node5", "test5")
	CreateLinkTable("link5", "test5")
	CreateDataTable("data5", "test5")
}
