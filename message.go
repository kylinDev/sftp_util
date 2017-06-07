package sftp_util

import (
	"fmt"
	"log"
)

func (util *SftpUtil) Message(msg string) {
	if util.Log {
		log.Println(msg)
	} else {
		fmt.Println(msg)
	}
}
