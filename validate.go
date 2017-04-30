package sftp_util

import (
	"fmt"
	"os"
)

func (util *SftpUtil) ValidateDirs() (err error) {
	var remoteDirInfo os.FileInfo

	// Validate remote directory exists
	remoteDirInfo, err = util.Client.Stat(util.Rdir)
	if err != nil {
		return fmt.Errorf("Problem accessing remote directory %s: %v", util.Rdir, err)
	}
	if !remoteDirInfo.IsDir() {
		return fmt.Errorf("Remote path %s is not a directory", util.Rdir)
	}
	if util.Type == "LS" {
		// For LS, nothing more to check
		return
	}

	// Validate local directory exists
	_, err = os.Stat(util.Ldir)
	if err != nil {
		return fmt.Errorf("Local directory %q missing: %v", util.Ldir, err)
	}

	// Validate remote file
	util.rFileInfo, err = util.Client.Stat(util.rFilePath)
	if err == nil && util.rFileInfo.IsDir() {
		return fmt.Errorf("Remote file %s is a directory", util.rFilePath)
	}

	if util.Type == "GET" {
		// For Get, remote file must exist
		if err != nil {
			return fmt.Errorf("Problem accessing remote file %s: %v", util.rFilePath, err)
		}
	} else {
		// For Put, warn if it exists
		if err == nil {
			fmt.Printf("Remote file %s already exists, will overwrite", util.rFilePath)
		}
		// File doesn't exist is normal, clear error
		err = nil
	}
	return
}
