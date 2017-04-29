package sftp_util

import (
	"fmt"
	"github.com/pkg/sftp"
	"log"
	"os"
)

func (util *SftpUtil) GetFile() (err error) {
	var lfile *os.File
	var rfile *sftp.File

	err = util.ValidateDirs()
	if err != nil {
		return
	}

	rfile, err = util.Client.Open(util.RFilePath)
	if err != nil {
		return fmt.Errorf("Cannot read remote file: %v", err)
	}

	lfile, err = os.OpenFile(util.LFilePath, os.O_CREATE+os.O_WRONLY, util.RFileInfo.Mode())
	if err != nil {
		return fmt.Errorf("Cannot write local file: %v", err)
	}

	log.Printf("Getting File %s\n", util.RFilePath)
	var b []byte = make([]byte, BUFSIZE)
	var n, m int
	for {
		n, err = rfile.Read(b)
		m, err = lfile.Write(b[:n])
		if err != nil {
			return fmt.Errorf("Problem writing local file: %v", err)
		}
		if n != m {
			return fmt.Errorf("Attempted to write %d bytes, but wrote %d to local file", n, m)
		}

		if n != BUFSIZE {
			lfile.Close()
			rfile.Close()
			break
		}
	}

	return
}
