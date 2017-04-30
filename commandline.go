package sftp_util

import (
	"flag"
	"fmt"
	"path/filepath"
)

func (util *SftpUtil) init() {
	// Define command-line arguments
	flag.StringVar(&util.Rdir, "rdir", ".", "Remote directory")
	flag.StringVar(&util.Ldir, "ldir", ".", "Local directory")
	flag.StringVar(&util.Filename, "file", "", "File to transfer")
	flag.StringVar(&util.Type, "type", "", "GET, PUT, RM or LS (required)")
}

// Make sure required arguments are specified
func (util *SftpUtil) validateFlags() (err error) {

	if util.Type == "" {
		return fmt.Errorf("-type not specified")
	}
	if util.Type == "GET" || util.Type == "PUT" || util.Type == "RM" {
		if util.Filename == "" {
			return fmt.Errorf("-file not specified")
		}
	} else if util.Type != "LS" {
		return fmt.Errorf("-type must be GET, PUT or LS")
	}

	return
}

func GetCmdLine() (util *SftpUtil, err error) {

	util = &SftpUtil{}
	util.init()

	// Parse commandline flag arguments
	flag.Parse()

	// Validate
	err = util.validateFlags()
	if err != nil {
		return nil, err
	}

	util.lFilePath = filepath.Join(util.Ldir, util.Filename)
	util.rFilePath = filepath.Join(util.Rdir, util.Filename)

	// Return commandline context
	return util, nil
}
