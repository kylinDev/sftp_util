# sftp_util
Example Go-lang command-line utility for SFTP
* Uses the "[github.com/pkg/sftp](https://godoc.org/github.com/pkg/sftp)" and "[golang.org/x/crypto/ssh](https://godoc.org/github.com/pkg/sftp)" packages
* Implements GET, PUT and LS
* Preserves the unix file permissions (FileInfo.Mode)

### How to use the package
Confugre the settings for your file transfer, as well as server and credentials, in the SftpUtil struct.  The following comments explain the fields to configure:
```go
type SftpUtil struct {
	Rdir      string // Remote directory
	Ldir      string // Local directory
	Filename  string // File to transfer
	Type      string // GET or PUT
	User      string // Username
	Pass      string // Password
	Host      string // Hostname or IP Address
	Port      string // TCP port
	lFilePath string
	rFilePath string
	lFileInfo os.FileInfo
	rFileInfo os.FileInfo
	Client    *sftp.Client
}
```
The routine GetCmdLine() reads comand-line flags to specify the file details and type of transfer. These parameters set the first four fields in the above struct.

### Sample Code
A sample main() routine for making a stand-alone sftp utilty is provided in [sftp_cmd/main.go](https://github.com/DavidSantia/sftp_util/blob/master/sftp_cmd/main.go)

```go
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
			fmt.Printf("GET error: %v\n", err)
			os.Exit(1)
		}
	} else if cmd.Type == "PUT" {
		err = cmd.PutFile()
		if err != nil {
			fmt.Printf("PUT error: %v\n", err)
			os.Exit(1)
		}
	} else {
		err = cmd.LsDir()
		if err != nil {
			fmt.Printf("LS error: %v\n", err)
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
    	File to transfer
  -ldir string
    	Local directory (default ".")
  -rdir string
    	Remote directory (default ".")
  -type string
    	GET, PUT or LS (required)
```
