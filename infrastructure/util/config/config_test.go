package config

import (
	"testing"

	log "github.com/Rock-liyi/p2pdb/infrastructure/util/log"
	_ "github.com/joho/godotenv/autoload"
)

func TestDatabase_Name(t *testing.T) {
	//require := require.New(t)
	// path := GetDataPath()
	// env := GetEnv()
	isdebug := IsDebug()
	// log.Debug(path)
	// log.Debug(env)
	log.Debug(isdebug)

}
