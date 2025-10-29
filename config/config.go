package config

import (
	"encoding/json"
	"log"

	"github.com/hello/utils"
)

type HelloMongoConfigType struct {
	DbName        string `json:"db_name"`
	Type          string `json:"type"`
	ConnectionURL string `json:"connection_url"`
}
type databaseConfigType struct {
	Hello HelloMongoConfigType `json:"hello"`
}

type configType struct {
	Service   string             `json:"service"`
	Port      uint               `json:"port"`
	Databases databaseConfigType `json:"databases"`
}

var Config configType

func LoadConfig() {
	fileContent, err := utils.ReadFromFile("config/config.json")
	if err != nil {
		log.Fatal("[err] error in loading file", err)
		panic(err)
	}

	err = json.Unmarshal(fileContent, &Config)
	if err != nil {
		log.Fatal("[err] error in config parsing", err)
		panic(err)
	}
}
