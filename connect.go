package sftp_util

import (
	"fmt"
	"log"

	"github.com/pkg/sftp"
	"golang.org/x/crypto/ssh"
	"net"
)

func DoNothing(hostname string, remote net.Addr, key ssh.PublicKey) (err error) {
	// Using Password mode, HostKey mode not supported
	return nil
}

func (util *SftpUtil) Connect() (err error) {
	var sshClient *ssh.Client

	server := util.Host + ":" + util.Port
	log.Printf("Connecting to %s\n", server)

	config := &ssh.ClientConfig{
		User:            util.User,
		Auth:            []ssh.AuthMethod{ssh.Password(util.Pass)},
		HostKeyCallback: DoNothing,
	}

	sshClient, err = ssh.Dial("tcp", server, config)
	if err != nil {
		return fmt.Errorf("Failed to connect: %v", err)
	}

	util.Client, err = sftp.NewClient(sshClient)
	if err != nil {
		return fmt.Errorf("Failed to create SFTP client: ", err)
	}
	return
}
