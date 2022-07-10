package config

import (
	"testing"

	debug "github.com/favframework/debug"
	_ "github.com/joho/godotenv/autoload"
)

func TestDatabase_Name(t *testing.T) {
	//require := require.New(t)
	path := GetDataPath()
	env := GetEnv()
	isdebu := IsDebug()
	debug.Dump(path)
	debug.Dump(env)
	debug.Dump(isdebu)
	//require.Emptyf(path)
	// val, ex := os.LookupEnv("DATAPATH")
	// debug.Dump(val)
	// debug.Dump(ex)
}
