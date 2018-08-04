package sftp_util

import (
	"fmt"
	"log"
	"os"

	"github.com/pkg/sftp"
	"path/filepath"
)

func (sftpSettings *SftpSettings) PutFile() (err error) {
	var localFile *os.File
	var localFileInfo os.FileInfo
	var remoteFile *sftp.File
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
	localFilename = filepath.Join(sftpSettings.Ldir, sftpSettings.Filename)
	localFileInfo, err = validateFile(localFilename)
	if err != nil {
		return
	}

	// Open local file to read
	localFile, err = os.Open(localFilename)
	if err != nil {
		return fmt.Errorf("cannot open local file: %v", err)
	}

	// Open remote file to write
	remoteFilename = filepath.Join(sftpSettings.Rdir, sftpSettings.Filename)
	remoteFile, err = sftpSettings.Client.Create(remoteFilename)
	if err != nil {
		return fmt.Errorf("Cannot create remote file: %v", err)
	}

	log.Printf("Putting %s", remoteFilename)
	var b []byte = make([]byte, SFTP_BUFSIZE)
	var n, m int
	for {
		n, err = localFile.Read(b)
		m, err = remoteFile.Write(b[:n])
		if err != nil {
			return fmt.Errorf("writing remote file: %v", err)
		}
		if n != m {
			return fmt.Errorf("attempted to write %d bytes, but wrote %d to remote file", n, m)
		}

		if n != SFTP_BUFSIZE {
			remoteFile.Close()
			localFile.Close()
			break
		}
	}

	if !sftpSettings.NoChmod {
		err = sftpSettings.Client.Chmod(remoteFilename, localFileInfo.Mode())
		if err != nil {
			return fmt.Errorf("cannot set remote file permissions: %v", err)
		}
	}
	return
}
