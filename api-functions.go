package main

import (
	"encoding/json"
	"fmt"
)

func GetWhoami() (result Whoami, err error) {
	var apiRes Whoami
	req, err := WebRequest("GET", "User/Whoami", nil)

	err = json.Unmarshal(req, &apiRes)

	return apiRes, err
}

func CretaeDomain(name string, planAlias string, username string,
	password string, activedomainuser bool, firstname string,
	lastname string, email string, expiration string) (result DomainResult, err error) {

	var apRes DomainResult

	params := map[string]string{"name": name,
		"planAlias":        planAlias,
		"username":         username,
		"password":         password,
		"activedomainuser": fmt.Sprintf("%v", activedomainuser),
		"firstname":        firstname,
		"lastname":         lastname,
		"email":            email,
		"expiration":       expiration}

	req, err := WebRequest("POST", "Domain/Create", params)

	err = json.Unmarshal(req, &apRes)

	return apRes, err
}
