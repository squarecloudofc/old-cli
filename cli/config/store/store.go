package store

import (
	"os"
	"path/filepath"
	"runtime"
)

func GetDataDir() string {
	var path string

	if a := os.Getenv("AppData"); runtime.GOOS == "windows" && a != "" {
		path = filepath.Join(a, "Square Cloud CLI")
	} else if b := os.Getenv("XDG_CONFIG_HOME"); b != "" {
		path = filepath.Join(b, "squarecloud")
	} else {
		dir, err := os.UserHomeDir()
		if err != nil {
			panic("cannot get a valid dir for square cloud config")
		}
		path = filepath.Join(dir, ".config", "squarecloud")
	}

	if _, err := os.Stat(path); err != nil {
		os.MkdirAll(path, 0777)
	}

	return path
}
