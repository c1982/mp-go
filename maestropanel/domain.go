package api

import (
	"encoding/json"
	"fmt"
)

type Domain struct {
}

func (o *Domain) CretaeDomain(name string, planAlias string, username string,
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

	req, err := webRequest("POST", "Domain/Create", params)

	err = json.Unmarshal(req, &apRes)

	return apRes, err
}
