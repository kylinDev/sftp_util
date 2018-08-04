package sftp_util

import (
	"os"

	"github.com/pkg/sftp"
)

// Byte buffer for file I/O (128K)
const SFTP_BUFSIZE = 131072

type SftpSettings struct {
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
	lFileInfo os.FileInfo
	rFileInfo os.FileInfo
	Client    *sftp.Client
}
