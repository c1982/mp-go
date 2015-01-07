package api

import (
	"encoding/json"
)

type User struct{}

func (User) GetWhoami() (result Whoami, err error) {
	var apiRes Whoami
	req, err := webRequest("GET", "User/Whoami", nil)

	err = json.Unmarshal(req, &apiRes)

	return apiRes, err
}
