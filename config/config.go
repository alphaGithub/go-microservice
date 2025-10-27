package config

import (
	"encoding/json"
	"log"

	"github.com/hello/utils"
)

type configType struct {
	Service string `json:"service"`
	Port    uint   `json:"port"`
}

var Config configType

func LoadConfig() {
	fileContent, err := utils.ReadFromFile("config/config.json")
	if err != nil {
		log.Fatal("[err] error in loading file", err)
		panic(err)
	}

	errParsing := json.Unmarshal(fileContent, &Config)
	if errParsing != nil {
		log.Fatal("[err] error in config parsing", errParsing)
	}
}
