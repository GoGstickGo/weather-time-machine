package utils

import (
	"fmt"
	"os"
)

func LogFile() (file *os.File, err error) {
	file, err = os.OpenFile("/var/log/wtm_logs", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		return nil, fmt.Errorf("can't create log file")
	}
	return file, nil
}
