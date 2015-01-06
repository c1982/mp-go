// mp-api project main.go
package main

import (
	"fmt"
)

func main() {

	whoami, err := GetWhoami()

	fmt.Println("Error: ", err)
	fmt.Println("ErrorCode: ", whoami.ErrorCode)
	fmt.Println("StatusCode: ", whoami.StatusCode)
	fmt.Println("StatusCode: ", whoami.Details.Email)
}
