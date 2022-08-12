package config

import (
	"fmt"
	"os"

	log "github.com/Rock-liyi/p2pdb/infrastructure/util/log"
	"github.com/caarlos0/env/v6"

	//"github.com/dolthub/vitess/go/vt/log"

	"github.com/joho/godotenv"
	_ "github.com/joho/godotenv/autoload"
)

type Env struct {
	DEBUG               bool   `env:"DEBUG"`
	ENV                 string `env:"ENV"`
	DB_NAME             string `env:"DB_NAME"`
	DB_TABLE_NAME       string `env:"DB_TABLE_NAME"`
	DATA_PATH           string `env:"DATA_PATH"`
	STORAGE_ENGINE      string `env:"STORAGE_ENGINE"`
	PORT                string `env:"PORT"`
	DB_HOUST            string `env:"DB_HOST"`
	DB_INFORMATION_NAME string `env:"DB_INFORMATION_NAME"`
	INTERNAL_DATA_PATH  string `env:"INTERNAL_DATA_PATH"`
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
		currentWorkDirectory = os.Getenv("GOPATH") + "/src/p2pdb"
		err = godotenv.Load(currentWorkDirectory + `/.env`)
		if err != nil {
			log.Error("p2pdb -> Error loading .env file, the path is" + currentWorkDirectory + `/.env`)
		}
	}
	conf := Env{}
	if err := env.Parse(&conf); err != nil {
		log.Warn(fmt.Printf("%+v\n", err))
		fmt.Printf("%+v\n", conf)
	}
	log.Info(fmt.Printf("%+v\n", conf))
}

func IsDebug() bool {
	conf := Env{}
	if err := env.Parse(&conf); err != nil {
		log.Warn(fmt.Printf("%+v\n", err))
		fmt.Printf("%+v\n", conf)
	}

	//log.Info(conf)
	//fmt.Printf("%+v\n", conf)
	return conf.DEBUG
}

func GetEnv() string {
	conf := Env{}
	if err := env.Parse(&conf); err != nil {
		log.Warn(fmt.Printf("%+v\n", err))
		fmt.Printf("%+v\n", conf)
	}
	return conf.ENV
}

func GetDataPath() string {
	conf := Env{}
	if err := env.Parse(&conf); err != nil {
		log.Warn(fmt.Printf("%+v\n", err))
		fmt.Printf("%+v\n", conf)
	}
	return conf.DATA_PATH
}

func GetInternaleDataPath() string {
	conf := Env{}
	if err := env.Parse(&conf); err != nil {
		log.Warn(fmt.Printf("%+v\n", err))
		fmt.Printf("%+v\n", conf)
	}
	return conf.INTERNAL_DATA_PATH
}

func GetDBTableName() string {
	conf := Env{}
	if err := env.Parse(&conf); err != nil {
		log.Warn(fmt.Printf("%+v\n", err))
		fmt.Printf("%+v\n", conf)
	}
	return conf.DB_TABLE_NAME
}

func GetDBName() string {
	conf := Env{}
	if err := env.Parse(&conf); err != nil {
		log.Warn(fmt.Printf("%+v\n", err))
		fmt.Printf("%+v\n", conf)
	}
	return conf.DB_NAME
}

func GetStorageEngine() string {
	conf := Env{}
	if err := env.Parse(&conf); err != nil {
		log.Warn(fmt.Printf("%+v\n", err))
		fmt.Printf("%+v\n", conf)
	}
	return conf.STORAGE_ENGINE
}

func GetDBHost() string {
	conf := Env{}
	if err := env.Parse(&conf); err != nil {
		log.Warn(fmt.Printf("%+v\n", err))
		fmt.Printf("%+v\n", conf)
	}
	return conf.DB_HOUST
}

func GetPort() string {
	conf := Env{}
	if err := env.Parse(&conf); err != nil {
		log.Warn(fmt.Printf("%+v\n", err))
		fmt.Printf("%+v\n", conf)
	}

	return conf.PORT
}

func GetDBInformationName() string {
	conf := Env{}
	if err := env.Parse(&conf); err != nil {
		log.Warn(fmt.Printf("%+v\n", err))
		fmt.Printf("%+v\n", conf)
	}

	return conf.DB_INFORMATION_NAME
}

func GetDefualtDatabaseAddress() string {
	path := GetDataPath()
	DBname := GetDBName()

	DBPath := path + "/" + DBname + ".db"
	return DBPath
}
