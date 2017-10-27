package sftp_util

import (
	"flag"
	"fmt"
	"path/filepath"
)

func (util *SftpUtil) init() {
	// Define command-line arguments
	flag.BoolVar(&util.NoChmod, "nochmod", false, "Don't attempt to set remote file permissions")
	flag.StringVar(&util.Rdir, "rdir", ".", "Remote directory")
	flag.StringVar(&util.Ldir, "ldir", ".", "Local directory")
	flag.StringVar(&util.Filename, "file", "", "File to transfer")
	flag.StringVar(&util.Type, "type", "", "GET, PUT, RM or LS (required)")
}

// Make sure required arguments are specified
func (util *SftpUtil) ValidateOptions() (err error) {

	if util.Type == "" {
		return fmt.Errorf("type not specified")
	}
	if util.Type == "GET" || util.Type == "PUT" || util.Type == "RM" {
		if util.Filename == "" {
			return fmt.Errorf("file not specified")
		}
	} else if util.Type != "LS" {
		return fmt.Errorf("type must be GET, PUT or LS")
	}

	util.lFilePath = filepath.Join(util.Ldir, util.Filename)
	util.rFilePath = filepath.Join(util.Rdir, util.Filename)

	return
}

func GetCmdLine() (util *SftpUtil, err error) {

	util = &SftpUtil{}
	util.init()

	// Parse commandline flag arguments
	flag.Parse()

	// Validate
	err = util.ValidateOptions()
	if err != nil {
		return nil, fmt.Errorf("-%v", err)
	}

	// Return commandline context
	return util, nil
}
