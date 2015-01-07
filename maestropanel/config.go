package api

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"os"
)

type apiAccess struct {
	Host  string
	Port  int
	Key   string
	Https bool
}

func getConfig() (result apiAccess, err error) {
	var cnf apiAccess

	data, err := readConfigFile()

	err = yaml.Unmarshal(data, &cnf)
	if err != nil {
		return cnf, err
	}

	return cnf, err
}

func readConfigFile() (result []byte, err error) {
	configFile := "config.yaml"

	if !fileExists(configFile) {
		currentPath, _ := os.Getwd()

		log.Fatal("config.yaml not found in ", currentPath)
	}

	data, err := ioutil.ReadFile(configFile)

	return data, err
}

func fileExists(name string) bool {
	isExists := false

	if _, err := os.Stat(name); err == nil {
		isExists = true
	}

	return isExists
}
