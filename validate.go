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

	if util.Type != "RM" {
		// Validate local directory exists for GET / PUT
		_, err = os.Stat(util.Ldir)
		if err != nil {
			return fmt.Errorf("Local directory %q missing: %v", util.Ldir, err)
		}
	}

	// Validate remote file
	util.rFileInfo, err = util.Client.Stat(util.rFilePath)
	if err == nil && util.rFileInfo.IsDir() {
		return fmt.Errorf("Remote file %s is a directory", util.rFilePath)
	}

	if util.Type == "PUT" {
		// For PUT, warn if it exists
		if err == nil {
			util.Message("Remote file " + util.rFilePath + " already exists, will overwrite")
		}
		// File doesn't exist is normal, clear error
		err = nil
	} else {
		// For GET or RM, remote file must exist
		if err != nil {
			return fmt.Errorf("Problem accessing remote file %s: %v", util.rFilePath, err)
		}
	}

	return
}
