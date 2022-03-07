package main

import (
	"fmt"
	"os"

	"gitlab.com/maxmac99/goport/helper/service"
)

func main() {
	command := os.Args[1]

	if command == "build" {
		if err := service.Build(); err != nil {
			fmt.Println("Error: ", err)
			os.Exit(1)
		}
	} else if command == "pull" {
		if err := service.Pull(); err != nil {
			os.Exit(1)
		}
	} else {
		os.Exit(1)
	}
}
