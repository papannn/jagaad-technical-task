package config

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
)

type Config struct {
	BaseURL []string `json:"base_url"`
}

var (
	ConfigObj = Config{}
)

func InitializeConfig() {
	currentDir, _ := os.Getwd()
	jsonConfig, err := os.Open(fmt.Sprintf("%s%s", currentDir, "/config/config.json"))
	if err != nil {
		log.Panic("can't find config.json file, please create it on config folder")
	}

	configByte, err := io.ReadAll(jsonConfig)
	if err != nil {
		log.Panic("error the data io")
	}

	err = json.Unmarshal(configByte, &ConfigObj)
	if err != nil {
		log.Panic("error reading config.json file, please use a correct json form")
	}
}
