package main

import (
	"fmt"
	"log"
	"os"

	"github.com/DavidSantia/json_configs"
	"github.com/DavidSantia/sftp_util"
)

type Credentials struct {
	SftpUser string `json:"username"`
	SftpPass string `json:"password"`
	SftpHost string `json:"host"`
	SftpPort string `json:"port"`
}

func main() {

	sftp := &sftp_util.SftpSettings{}

	err := sftp.GetCmdLine()
	if err != nil {
		fmt.Printf("Command-line error: %v\n", err)
		os.Exit(1)
	}

	credentials := &Credentials{}
	err = json_configs.ReadConfigFile(credentials, "./settings.json")
	if err != nil {
		fmt.Printf("Configuration file error: %v\n", err)
		os.Exit(1)
	}

	sftp.User = credentials.SftpUser
	sftp.Pass = credentials.SftpPass
	sftp.Host = credentials.SftpHost
	sftp.Port = credentials.SftpPort

	// Open SSH session
	err = sftp.Connect()
	if err != nil {
		fmt.Printf("Connect error: %v\n", err)
		os.Exit(1)
	}
	defer sftp.Client.Close()

	if sftp.Type == "GET" {
		err = sftp.GetFile()
		if err != nil {
			log.Printf("GET error: %v\n", err)
			os.Exit(1)
		}
	} else if sftp.Type == "PUT" {
		err = sftp.PutFile()
		if err != nil {
			log.Printf("PUT error: %v\n", err)
			os.Exit(1)
		}
	} else if sftp.Type == "RM" {
		err = sftp.RmFile()
		if err != nil {
			log.Printf("RM error: %v\n", err)
			os.Exit(1)
		}
	} else {
		err = sftp.LsDir()
		if err != nil {
			log.Printf("LS error: %v\n", err)
			os.Exit(1)
		}
	}

	os.Exit(0)
}
