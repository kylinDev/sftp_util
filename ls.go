package sftp_util

import (
	"fmt"
	"os"
)

func (util *SftpUtil) LsDir() (err error) {
	var dir []os.FileInfo
	var file os.FileInfo

	err = util.ValidateDirs()
	if err != nil {
		return
	}

	dir, err = util.Client.ReadDir(util.Rdir)
	if err != nil {
		return fmt.Errorf("Cannot read remote directory: %v", err)
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
