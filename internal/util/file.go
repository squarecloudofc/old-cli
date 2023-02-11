package util

import "os"

func FileExists(dir string) bool {
	_, err := os.Stat(dir)
	return err == nil
}
