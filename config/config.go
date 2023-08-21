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

var (
	Getwd      = os.Getwd
	Open       = os.Open
	ReadAll    = io.ReadAll
	Unmarshall = json.Unmarshal
)

func InitializeConfig() error {
	currentDir, err := Getwd()
	if err != nil {
		log.Println(err)
		log.Println("error can't get the current directory")
		return err
	}

	jsonConfig, err := Open(fmt.Sprintf("%s%s", currentDir, "/config/config.json"))
	if err != nil {
		log.Println(err)
		log.Println("error can't find config.json file, please create it on config folder")
		return err
	}

	configByte, err := ReadAll(jsonConfig)
	if err != nil {
		log.Println(err)
		log.Println("error on the data io")
		return err
	}

	err = Unmarshall(configByte, &ConfigObj)
	if err != nil {
		log.Println(err)
		log.Println("error reading config.json file, please use a correct json form")
		return err
	}

	return nil
}
