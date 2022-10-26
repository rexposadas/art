package util

import "os"

func EnsureDir(dirName string) error {

	err := os.MkdirAll(dirName, os.ModePerm)

	if err == nil || os.IsExist(err) {
		return nil
	} else {
		return err
	}
}
