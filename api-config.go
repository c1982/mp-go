package main

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

type ApiAccess struct {
	Host  string
	Port  int
	Key   string
	Https bool
}

func GetConfig() (result ApiAccess, err error) {
	var cnf ApiAccess

	data, err := ioutil.ReadFile("config.yaml")

	if err != nil {
		return cnf, err
	}

	err = yaml.Unmarshal(data, &cnf)
	if err != nil {
		return cnf, err
	}

	return cnf, err
}
