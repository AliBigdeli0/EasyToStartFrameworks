package pkg

import (
	"os"
	"path"
)

func IsFileExist(file_name string) (bool, error) {
	file_info, err := os.Stat(file_name)
	if err != nil {
		return false, err
	}

	return !file_info.IsDir(), nil
}

func IsDirExist(path string) (bool, error) {
	dir_info, err := os.Stat(path)
	if err != nil {
		return false, err
	}
	return dir_info.IsDir(), nil
}

func GetConfigPath(args *IArgs) string {
	return path.Join((*args).GetBasePath(), (*args).GetConfigPath())
}

func GetLogConfigPath(file_name string, args *IArgs) string {
	return path.Join((*args).GetBasePath(), file_name)
}
