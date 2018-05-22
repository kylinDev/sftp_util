package sftp_util

import (
	"fmt"
	"io/ioutil"

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
	var key []byte
	var signer ssh.Signer

	server := util.Host + ":" + util.Port

	if len(util.Pass) == 0 && len(util.KeyFile) == 0 {
		return fmt.Errorf("Must provide Pass or KeyFile")
	}

	// Parse SSH key if provided
	if len(util.KeyFile) > 0 {
		key, err = ioutil.ReadFile(util.KeyFile)
		if err != nil {
			return fmt.Errorf("Failure reading private key: %v", err)
		}
		signer, err = ssh.ParsePrivateKey(key)
		if err != nil {
			return fmt.Errorf("Unable to parse private key: %v", err)
		}
	}

	// LS will give directory listing, no other output
	if util.Type != "LS" {
		util.Message("Connecting to " + server)
	}

	// Configure client for password and/or ssh key
	config := &ssh.ClientConfig{
		User:            util.User,
		HostKeyCallback: DoNothing,
	}
	if len(util.Pass) > 0 && len(util.KeyFile) > 0 {
		config.Auth = []ssh.AuthMethod{ssh.Password(util.Pass), ssh.PublicKeys(signer)}
	} else if len(util.Pass) > 0 {
		config.Auth = []ssh.AuthMethod{ssh.Password(util.Pass)}
	} else {
		config.Auth = []ssh.AuthMethod{ssh.PublicKeys(signer)}
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
