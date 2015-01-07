package main

import (
	"fmt"
	"mp-go/maestropanel"
)

func main() {

	_user := api.User{}

	whoami, err := _user.GetWhoami()

	fmt.Println("Error: ", err)
	fmt.Println("ErrorCode: ", whoami.ErrorCode)
	fmt.Println("StatusCode: ", whoami.StatusCode)
	fmt.Println("StatusCode: ", whoami.Details.Email)
}
