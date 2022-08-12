package entity

import (
	"log"

	function "github.com/Rock-liyi/p2pdb/infrastructure/util/function"
)

func GetNewDatabaseId(databaseName string) string {
	cid, err := function.GetCidString(databaseName)
	if err != nil {
		log.Fatal(err.Error())
	}
	return cid
}
