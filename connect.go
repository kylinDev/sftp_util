package sftp_util

import (
	"fmt"
	"io/ioutil"
	"log"
	"net"

	"github.com/pkg/sftp"
	"golang.org/x/crypto/ssh"
)

func DoNothing(hostname string, remote net.Addr, key ssh.PublicKey) (err error) {
	// Using Password mode, HostKey mode not supported
	return nil
}

func (sftpSettings *SftpSettings) Connect() (err error) {
	var sshClient *ssh.Client
	var key []byte
	var signer ssh.Signer

	server := sftpSettings.Host + ":" + sftpSettings.Port

	if len(sftpSettings.Pass) == 0 && len(sftpSettings.KeyFile) == 0 {
		return fmt.Errorf("Must provide Pass or KeyFile")
	}

	// Parse SSH key if provided
	if len(sftpSettings.KeyFile) > 0 {
		key, err = ioutil.ReadFile(sftpSettings.KeyFile)
		if err != nil {
			return fmt.Errorf("Failure reading private key: %v", err)
		}
		signer, err = ssh.ParsePrivateKey(key)
		if err != nil {
			return fmt.Errorf("Unable to parse private key: %v", err)
		}
	}

	// LS will give directory listing, no other output
	if sftpSettings.Type != "LS" {
		log.Printf("Connecting to %s", server)
	}

	// Configure client for password and/or ssh key
	config := &ssh.ClientConfig{
		User:            sftpSettings.User,
		HostKeyCallback: DoNothing,
	}
	if len(sftpSettings.Pass) > 0 && len(sftpSettings.KeyFile) > 0 {
		config.Auth = []ssh.AuthMethod{ssh.Password(sftpSettings.Pass), ssh.PublicKeys(signer)}
	} else if len(sftpSettings.Pass) > 0 {
		config.Auth = []ssh.AuthMethod{ssh.Password(sftpSettings.Pass)}
	} else {
		config.Auth = []ssh.AuthMethod{ssh.PublicKeys(signer)}
	}

	sshClient, err = ssh.Dial("tcp", server, config)
	if err != nil {
		err = fmt.Errorf("Failed to connect: %v", err)
		return
	}

	sftpSettings.Client, err = sftp.NewClient(sshClient)
	if err != nil {
		err = fmt.Errorf("Failed to create SFTP client: ", err)
	}
	return
}
