package repository

import (
	"testing"
)

func TestCreateTable(t *testing.T) {
	CreateNodeTable("node5")
	CreateLinkTable("link5")
	CreateDataTable("data5")
}
