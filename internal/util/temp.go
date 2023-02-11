package util

import (
	"os"
	"path/filepath"
)

func GetTempDir() (string, error) {
	tempDir := os.TempDir()
	tempFolder := filepath.Join(tempDir, "SquareTemp")
	if _, err := os.Stat(tempFolder); err != nil {
		if err := os.Mkdir(tempFolder, os.ModePerm); err != nil {
			return "", err
		}
	}

	return tempFolder, nil
}
