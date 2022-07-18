package store

import (
	"testing"

	debug "github.com/favframework/debug"
)

func TestDBFactory(t *testing.T) {

	var db *CreateDBFactory
	orm := db.InitDB()
	var node = &Node{}
	orm.Where("node_id = ?", "nodeid").First(&node)

	debug.Dump(node)
}

func TestDBSelect(t *testing.T) {

	var db *CreateDBFactory
	orm := db.InitDB()
	var node = &Node{}
	orm.Where("node_id = ?", "nodeid").Select(&node)

	debug.Dump(node)
}
