package sftp_util

import (
	"github.com/pkg/sftp"
	"golang.org/x/crypto/ssh"
)

// Byte buffer for file I/O (128K)
const SFTP_BUFSIZE = 131072

type SftpSettings struct {
	NoChmod  bool
	Rdir     string // Remote directory
	Ldir     string // Local directory
	Filename string // File to transfer
	Type     string // GET or PUT
	User     string // Username
	Pass     string // Password
	KeyFile  string // RSA Key file
	Host     string // Hostname or IP Address
	Port     string // TCP port
	Client   *sftp.Client
	SshClient *ssh.Client
}
