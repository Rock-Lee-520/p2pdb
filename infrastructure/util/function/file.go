package function

import "os"

func RemoveFile(path string) error {
	return os.Remove(path)
}

func CreateFile(path string) error {

	return nil
}
