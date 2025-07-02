package utils

import "os"

func IsExistsFolder(p string) bool {
	info, err := os.Stat(p)
	if os.IsNotExist(err); err != nil {
		return false
	}

	return info.IsDir()
}
