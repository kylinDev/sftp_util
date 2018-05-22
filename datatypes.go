package sftp_util

import (
	"os"

	"github.com/pkg/sftp"
)

type SftpUtil struct {
	Log       bool
	NoChmod   bool
	Rdir      string // Remote directory
	Ldir      string // Local directory
	Filename  string // File to transfer
	Type      string // GET or PUT
	User      string // Username
	Pass      string // Password
	KeyFile   string // RSA Key file
	Host      string // Hostname or IP Address
	Port      string // TCP port
	lFilePath string
	rFilePath string
	lFileInfo os.FileInfo
	rFileInfo os.FileInfo
	Client    *sftp.Client
}

// Byte buffer for file I/O (128K)
const BUFSIZE = 131072
