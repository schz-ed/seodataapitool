package utils

import "os"

func EnsureDir(dirName string) error {
	return os.MkdirAll(dirName, 0755)
}
