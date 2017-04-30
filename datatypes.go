package sftp_util

import (
	"os"

	"github.com/pkg/sftp"
)

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

// Byte buffer for file I/O
const BUFSIZE = 4096
