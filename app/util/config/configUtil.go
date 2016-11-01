package util

import (
	"encoding/json"
	"io/ioutil"
	"log"

	"./../../model"
)

//LoadConfig is used to Load Configurations from a JSON file
func LoadConfig(path string) *model.Configuration {
	file, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatal("Config File Missing. ", err)
	}

	var config model.Configuration
	err = json.Unmarshal(file, &config)
	if err != nil {
		log.Fatal("Config Parse Error: ", err)
	}

	return &config
}
