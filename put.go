package sftp_util

import (
	"fmt"
	"os"

	"github.com/pkg/sftp"
)

func (util *SftpUtil) PutFile() (err error) {
	var lfile *os.File
	var rfile *sftp.File

	err = util.ValidateDirs()
	if err != nil {
		return
	}

	util.lFileInfo, err = os.Stat(util.lFilePath)
	if err != nil {
		return fmt.Errorf("Cannot get local file permissions: %v", err)
	}
	lfile, err = os.Open(util.lFilePath)
	if err != nil {
		return fmt.Errorf("Cannot open local file: %v", err)
	}

	rfile, err = util.Client.Create(util.rFilePath)
	if err != nil {
		return fmt.Errorf("Cannot create remote file: %v", err)
	}

	util.Message("Putting File " + util.rFilePath)
	var b []byte = make([]byte, BUFSIZE)
	var n, m int
	for {
		n, err = lfile.Read(b)
		m, err = rfile.Write(b[:n])
		if err != nil {
			return fmt.Errorf("Problem writing remote file: %v", err)
		}
		if n != m {
			return fmt.Errorf("Attempted to write %d bytes, but wrote %d to remote file", n, m)
		}

		if n != BUFSIZE {
			lfile.Close()
			rfile.Close()
			break
		}
	}

	err = util.Client.Chmod(util.rFilePath, util.lFileInfo.Mode())
	if err != nil {
		return fmt.Errorf("Cannot set remote file permissions: %v", err)
	}

	return
}
