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
	flag.StringVar(&util.Filename, "file", "", "file (required)")
	flag.StringVar(&util.Type, "type", "", "GET or PUT (required)")
}

// Make sure required arguments are specified
func (util *SftpUtil) validateFlags() (err error) {
	if util.Filename == "" {
		return fmt.Errorf("-file not specified")
	}

	if util.Type == "" {
		return fmt.Errorf("-type not specified")
	}
	if util.Type != "GET" && util.Type != "PUT" {
		return fmt.Errorf("-type must be GET or PUT")
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

	util.LFilePath = filepath.Join(util.Ldir, util.Filename)
	util.RFilePath = filepath.Join(util.Rdir, util.Filename)

	// Return commandline context
	return util, nil
}
