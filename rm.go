package sftp_util

import (
	"fmt"
)

func (util *SftpUtil) RmFile() (err error) {
	err = util.ValidateDirs()
	if err != nil {
		return
	}

	err = util.Client.Remove(util.rFilePath)
	if err != nil {
		return fmt.Errorf("Cannot remove remote file: %v", err)
	}
	fmt.Printf("Removed File %s\n", util.rFilePath)

	return
}
