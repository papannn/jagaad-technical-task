package config

import (
	"encoding/json"
	"fmt"
	"io"
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
		fmt.Println(err)
		fmt.Println("error opening json file")
	}

	configByte, err := io.ReadAll(jsonConfig)
	if err != nil {
		fmt.Println("error reading json file")
	}

	err = json.Unmarshal(configByte, &ConfigObj)
	if err != nil {
		fmt.Println("error unmarshall config file")
	}
}
