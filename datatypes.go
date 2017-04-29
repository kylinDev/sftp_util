package sftp_util

import (
	"os"

	"github.com/pkg/sftp"
)

type SftpUtil struct {
	Rdir      string
	Ldir      string
	Filename  string
	LFilePath string
	RFilePath string
	LFileInfo os.FileInfo
	RFileInfo os.FileInfo
	Type      string
	User      string
	Pass      string
	Host      string
	Port      string
	Client    *sftp.Client
}

// Byte buffer for file I/O
const BUFSIZE = 4096
