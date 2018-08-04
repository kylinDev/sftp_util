package sftp_util

import (
	"fmt"
	"log"
	"path/filepath"
)

func (sftpSettings *SftpSettings) RmFile() (err error) {
	var remoteFilename string

	// Validate directory
	err = validateRemoteDir(sftpSettings.Client, sftpSettings.Rdir)
	if err != nil {
		return
	}
	remoteFilename = filepath.Join(sftpSettings.Rdir, sftpSettings.Filename)

	log.Printf("Removing %s", remoteFilename)
	err = sftpSettings.Client.Remove(remoteFilename)
	if err != nil {
		err = fmt.Errorf("cannot remove remote file: %v", err)
	}
	return
}
