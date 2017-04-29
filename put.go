package sftp_util

import (
	"fmt"
	"github.com/pkg/sftp"
	"log"
	"os"
)

func (util *SftpUtil) PutFile() (err error) {
	var lfile *os.File
	var rfile *sftp.File

	err = util.ValidateDirs()
	if err != nil {
		return
	}

	util.LFileInfo, err = os.Stat(util.LFilePath)
	if err != nil {
		log.Printf("Cannot get local file permissions: %v\n", err)
		return
	}
	lfile, err = os.Open(util.LFilePath)
	if err != nil {
		log.Printf("Cannot open local file: %v\n", err)
		return
	}

	rfile, err = util.Client.Create(util.RFilePath)
	if err != nil {
		return fmt.Errorf("Cannot create remote file: %v", err)
	}

	log.Printf("Putting File %s\n", util.RFilePath)
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

	err = util.Client.Chmod(util.RFilePath, util.LFileInfo.Mode())
	if err != nil {
		return fmt.Errorf("Cannot set remote file permissions: %v", err)
	}

	return
}
