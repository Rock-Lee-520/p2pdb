package config

import (
	"fmt"
	"os"

	"github.com/caarlos0/env/v6"
	log "github.com/sirupsen/logrus"

	//"github.com/dolthub/vitess/go/vt/log"

	"github.com/joho/godotenv"
	_ "github.com/joho/godotenv/autoload"
)

type Env struct {
	DEBUG          bool   `env:"DEBUG"`
	ENV            string `env:"ENV"`
	DBNAME         string `env:"DBNAME"`
	DBTABLENAME    string `env:"DBTABLENAME"`
	DATAPATH       string `env:"DATAPATH"`
	STORAGE_ENGINE string `env:"STORAGE_ENGINE"`
	PORT           string `env:"PORT"`
	DBHOUST        string `env:"DBHOST"`
}

const projectDirName = "p2pdb-store" // change to relevant project name

func init() {
	// projectName := regexp.MustCompile(`^(.*` + projectDirName + `)`)

	// rootPath := projectName.Find([]byte(currentWorkDirectory))
	// fmt.Printf("%+v\n", rootPath)
	// err := godotenv.Load(string(rootPath) + `/.env`)
	//filenames = godotenv.Load(filenames)

	currentWorkDirectory, _ := os.Getwd()

	err := godotenv.Load(currentWorkDirectory + `/.env`)
	if err != nil {
		log.Error("p2pdb -> Error loading .env file")
	}
}

func IsDebug() bool {
	conf := Env{}
	if err := env.Parse(&conf); err != nil {
		fmt.Printf("%+v\n", err)
	}
	fmt.Printf("%+v\n", conf)
	return conf.DEBUG
}

func GetEnv() string {
	conf := Env{}
	if err := env.Parse(&conf); err != nil {
		fmt.Printf("%+v\n", err)
	}
	fmt.Printf("%+v\n", conf)
	return conf.ENV
}

func GetDataPath() string {
	conf := Env{}
	if err := env.Parse(&conf); err != nil {
		fmt.Printf("%+v\n", err)
	}
	fmt.Printf("%+v\n", conf)
	return conf.DATAPATH
}

func GetDBTableName() string {
	conf := Env{}
	if err := env.Parse(&conf); err != nil {
		fmt.Printf("%+v\n", err)
	}
	fmt.Printf("%+v\n", conf)
	return conf.DBTABLENAME
}

func GetDBName() string {
	conf := Env{}
	if err := env.Parse(&conf); err != nil {
		fmt.Printf("%+v\n", err)
	}
	fmt.Printf("%+v\n", conf)
	return conf.DBNAME
}

func GetStorageEngine() string {
	conf := Env{}
	if err := env.Parse(&conf); err != nil {
		fmt.Printf("%+v\n", err)
	}
	fmt.Printf("%+v\n", conf)
	return conf.STORAGE_ENGINE
}

func GetDBHost() string {
	conf := Env{}
	if err := env.Parse(&conf); err != nil {
		fmt.Printf("%+v\n", err)
	}
	fmt.Printf("%+v\n", conf)
	return conf.DBHOUST
}

func GetPort() string {
	conf := Env{}
	if err := env.Parse(&conf); err != nil {
		fmt.Printf("%+v\n", err)
	}
	fmt.Printf("%+v\n", conf)
	return conf.PORT
}
