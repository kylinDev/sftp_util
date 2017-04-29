# sftp_util
Example Go-lang command-line utility for SFTP
* Uses the "[github.com/pkg/sftp](https://godoc.org/github.com/pkg/sftp)" and "[golang.org/x/crypto/ssh](https://godoc.org/github.com/pkg/sftp)" packages
* Implements GET and PUT
* Preserves the unix file permissions (FileInfo.Mode)

A sample main.go for making a stand-alone sftp utilty is provided in [sftp_cmd/main.go](https://github.com/DavidSantia/sftp_util/blob/master/sftp_cmd/main.go)

```go
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
```

### Building the project
You can buid directly:
```sh
$ cd sftp_cmd
$ go build
```

Alternatively, a build script is provided that builds for linux using a Docker container:
```sh
$ ./build.sh
```

### Running the utility
The utility takes the following flags:
```sh
Usage of ./sftp_cmd:
  -file string
    	file (required)
  -ldir string
    	Local directory (default ".")
  -rdir string
    	Remote directory (default ".")
  -type string
    	GET or PUT (required)

```
