package main

import (
	"fmt"
	"os"

	"github.com/DavidSantia/sftp_util"
)

func main() {

	cmd, err := sftp_util.GetCmdLine()
	if err != nil {
		fmt.Printf("Command-line error: %v\n", err)
		os.Exit(1)
	}

	cmd.User = "USER"
	cmd.Pass = "PASSWORD"
	cmd.Host = "HOST"
	cmd.Port = "22"

	// Open SSH session
	err = cmd.Connect()
	if err != nil {
		fmt.Printf("Connect error: %v\n", err)
		os.Exit(1)
	}
	defer cmd.Client.Close()

	if cmd.Type == "GET" {
		err = cmd.GetFile()
		if err != nil {
			fmt.Printf("Get error: %v\n", err)
			os.Exit(1)
		}
	} else {
		err = cmd.PutFile()
		if err != nil {
			fmt.Printf("Put error: %v\n", err)
			os.Exit(1)
		}
	}

	os.Exit(0)
}
