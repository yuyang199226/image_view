package utils

import (
	"time"
)

func PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}


func CurDatetime() string {
	return time.Now().Format("2006-01-02 15:04:05")
}
