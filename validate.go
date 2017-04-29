package sftp_util

import (
	"fmt"
	"log"
	"os"
)

func (util *SftpUtil) ValidateDirs() (err error) {
	var remoteDirInfo os.FileInfo

	_, err = os.Stat(util.Ldir)
	if err != nil {
		return fmt.Errorf("Local directory %q missing: %v", util.Ldir, err)
	}

	// Validate remote directory exists
	remoteDirInfo, err = util.Client.Stat(util.Rdir)
	if err != nil {
		return fmt.Errorf("Problem accessing remote directory %s: %v", util.Rdir, err)
	}
	if !remoteDirInfo.IsDir() {
		return fmt.Errorf("Remote path %s is not a directory", util.Rdir)
	}

	// Validate remote file
	util.RFileInfo, err = util.Client.Stat(util.RFilePath)
	if err == nil && util.RFileInfo.IsDir() {
		return fmt.Errorf("Remote file %s is a directory", util.RFilePath)
	}

	if util.Type == "GET" {
		// For Get, remote file must exist
		if err != nil {
			return fmt.Errorf("Problem accessing remote file %s: %v", util.RFilePath, err)
		}
	} else {
		// For Put, warn if it exists
		if err == nil {
			log.Printf("Remote file %s already exists, will overwrite", util.RFilePath)
		}
	}
	return
}
