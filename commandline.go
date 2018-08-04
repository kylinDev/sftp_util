package sftp_util

import (
	"flag"
	"fmt"
)

func (util *SftpSettings) GetCmdLine() (err error) {

	// Define command-line arguments
	flag.BoolVar(&util.NoChmod, "nochmod", false, "Don't attempt to set remote file permissions")
	flag.StringVar(&util.Rdir, "rdir", ".", "Remote directory")
	flag.StringVar(&util.Ldir, "ldir", ".", "Local directory")
	flag.StringVar(&util.Filename, "file", "", "File to transfer")
	flag.StringVar(&util.Type, "type", "", "GET, PUT, RM or LS (required)")

	// Parse commandline flag arguments
	flag.Parse()

	// Validate
	if util.Type == "" {
		err = fmt.Errorf("-type (GET, PUT, RM or LS) required")
		return
	}
	if util.Type == "GET" || util.Type == "PUT" || util.Type == "RM" {
		if util.Filename == "" {
			return fmt.Errorf("-file {filename} required")
		}
	} else if util.Type != "LS" {
		return fmt.Errorf("-type must be GET, PUT, RM or LS")
	}

	return
}
