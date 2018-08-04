# sftp_util
Example Go-lang command-line utility for SFTP
* Uses the "[github.com/pkg/sftp](https://godoc.org/github.com/pkg/sftp)" and "[golang.org/x/crypto/ssh](https://godoc.org/github.com/pkg/sftp)" packages
* Implements GET, PUT, RM and LS
* Preserves the unix file permissions (FileInfo.Mode)

### How to use the package
Confugre the settings for your file transfer, as well as server and credentials, in the SftpUtil struct.  The following comments explain the fields to configure:
```go
type SftpUtil struct {
	Log       bool
	Rdir      string // Remote directory
	Ldir      string // Local directory
	Filename  string // File to transfer
	Type      string // GET or PUT
	User      string // Username
	Pass      string // Password
	Host      string // Hostname or IP Address
	Port      string // TCP port
	Client    *sftp.Client
}
```
The routine GetCmdLine() reads comand-line flags to specify the file details and type of transfer. These parameters set the first four fields in the above struct.

### Sample Code
A sample main() routine for making a stand-alone sftp utilty is provided in [sftp-cmd/sftp-cmd.go](https://github.com/DavidSantia/sftp_util/blob/master/sftp-cmd/sftp-cmd.go)

```go
package main

import (
	"fmt"
	"github.com/DavidSantia/sftp_util"
)

func main() {
	sftp := &sftp_util.SftpSettings{}
	err := sftp.GetCmdLine()
	if err != nil {
		fmt.Printf("Command-line error: %v\n", err)
		return
	}
	sftp.User = "USER"
	sftp.Pass = "PASSWORD"
	sftp.Host = "HOST"
	sftp.Port = "22"

	// Open SSH session
	err = sftp.Connect()
	if err != nil {
		fmt.Printf("Connect error: %v\n", err)
		return
	}
	defer sftp.Client.Close()

	if sftp.Type == "GET" {
		err = sftp.GetFile()
		if err != nil {
			fmt.Printf("GET error: %v\n", err)
			return
		}
	} else if sftp.Type == "PUT" {
		err = sftp.PutFile()
		if err != nil {
			fmt.Printf("PUT error: %v\n", err)
			return
		}
	} else if sftp.Type == "RM" {
		err = sftp.RmFile()
		if err != nil {
			fmt.Printf("RM error: %v\n", err)
			return
		}
	} else {
		err = sftp.LsDir()
		if err != nil {
			fmt.Printf("LS error: %v\n", err)
			return
		}
	}
}
```

### Building the project
You can buid directly:
```sh
$ cd sftp-cmd
$ go build
```

Alternatively, a build script is provided that builds for linux using a Docker container:
```sh
$ ./build.sh
```

### Running the utility
The utility takes the following flags:
```sh
Usage of ./sftp-cmd:
  -file string
    	File to transfer
  -ldir string
    	Local directory (default ".")
  -nochmod
    	Don't attempt to set remote file permissions
  -rdir string
    	Remote directory (default ".")
  -type string
    	GET, PUT, RM or LS (required)
```
