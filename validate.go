package sftp_util

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/pkg/sftp"
)

func validateFile(filename string) (fileInfo os.FileInfo, err error) {
	var fullpath string

	fullpath, err = filepath.Abs(filename)
	if err != nil {
		err = fmt.Errorf("%v [%s]", err, filename)
		return
	}
	fileInfo, err = os.Stat(fullpath)
	if os.IsNotExist(err) {
		err = fmt.Errorf("invalid, no such file [%s]", filename)
		return
	} else if os.IsPermission(err) {
		err = fmt.Errorf("invalid, permission denied [%s]", filename)
		return
	} else if err != nil {
		err = fmt.Errorf("invalid, %v [%s]", err, filename)
		return
	} else if fileInfo.IsDir() {
		err = fmt.Errorf("invalid, filename is a directory [%s]", filename)
	}
	return
}

func validateDir(dirname string) (err error) {
	var fullpath string
	var dirInfo os.FileInfo

	fullpath, err = filepath.Abs(dirname)
	if err != nil {
		err = fmt.Errorf("%v [%s]", err, dirname)
		return
	}
	dirInfo, err = os.Stat(fullpath)
	if os.IsNotExist(err) {
		err = fmt.Errorf("invalid, no such directory [%s]", dirname)
		return
	} else if os.IsPermission(err) {
		err = fmt.Errorf("invalid, permission denied [%s]", dirname)
		return
	} else if err != nil {
		err = fmt.Errorf("invalid, %v [%s]", err, dirname)
		return
	} else if !dirInfo.IsDir() {
		err = fmt.Errorf("invalid, not a directory [%s]", dirname)
	}
	return
}

func validateRemoteDir(client *sftp.Client, dirname string) (err error) {
	var remoteDirInfo os.FileInfo

	// Validate remote directory exists
	remoteDirInfo, err = client.Stat(dirname)
	if err != nil {
		err = fmt.Errorf("accessing remote directory %s: %v", dirname, err)
		return
	}
	if !remoteDirInfo.IsDir() {
		err = fmt.Errorf("remote path %s is not a directory", dirname)
	}
	return
}

func validateRemoteFile(client *sftp.Client, filename string) (remoteFileInfo os.FileInfo, err error) {

	// Validate remote directory exists
	remoteFileInfo, err = client.Stat(filename)
	if err != nil {
		err = fmt.Errorf("accessing remote file %s: %v", filename, err)
		return
	}
	if remoteFileInfo.IsDir() {
		err = fmt.Errorf("remote path %s is a directory", filename)
	}
	return
}
