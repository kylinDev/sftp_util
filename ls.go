package sftp_util

import (
	"fmt"
	"os"
)

func (sftpSettings *SftpSettings) LsDir() (err error) {
	var dir []os.FileInfo
	var file os.FileInfo

	// Validate directories and files
	err = validateRemoteDir(sftpSettings.Client, sftpSettings.Rdir)
	if err != nil {
		return
	}

	dir, err = sftpSettings.Client.ReadDir(sftpSettings.Rdir)
	if err != nil {
		return fmt.Errorf("cannot read remote directory: %v", err)
	}
	for _, file = range dir {
		if file.IsDir() {
			// Only list files
			continue
		}
		fmt.Println(file.Name())
	}
	return
}
