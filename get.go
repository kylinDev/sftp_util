package sftp_util

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/pkg/sftp"
)

func (sftpSettings *SftpSettings) GetFile() (err error) {
	var localFile *os.File
	var remoteFile *sftp.File
	var remoteFileInfo os.FileInfo
	var localFilename, remoteFilename string

	// Validate directories and files
	err = validateDir(sftpSettings.Ldir)
	if err != nil {
		return
	}
	err = validateRemoteDir(sftpSettings.Client, sftpSettings.Rdir)
	if err != nil {
		return
	}
	remoteFilename = filepath.Join(sftpSettings.Rdir, sftpSettings.Filename)
	remoteFileInfo, err = validateRemoteFile(sftpSettings.Client, remoteFilename)
	if err != nil {
		return
	}
	//localFilename = filepath.Join(sftpSettings.Ldir, sftpSettings.Filename)
	localFilename = filepath.Join("./", sftpSettings.Filename)

	// Open local file to write
	localFile, err = os.OpenFile(localFilename, os.O_CREATE|os.O_WRONLY, remoteFileInfo.Mode())
	if err != nil {
		return fmt.Errorf("cannot write local file: %v", err)
	}

	// Open remote file to read
	remoteFile, err = sftpSettings.Client.Open(remoteFilename)
	if err != nil {
		return fmt.Errorf("cannot read remote file: %v", err)
	}

	log.Printf("Getting %s", remoteFilename)
	var b []byte = make([]byte, SFTP_BUFSIZE)
	var n, m int
	for {
		n, err = remoteFile.Read(b)
		m, err = localFile.Write(b[:n])
		if err != nil {
			return fmt.Errorf("writing local file: %v", err)
		}
		if n != m {
			return fmt.Errorf("attempted to write %d bytes, but wrote %d to local file", n, m)
		}

		if n != SFTP_BUFSIZE {
			localFile.Close()
			remoteFile.Close()
			break
		}
	}
	return
}
