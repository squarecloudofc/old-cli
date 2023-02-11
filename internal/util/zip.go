package util

import (
	"archive/zip"
	"io"
	"io/fs"
	"os"
	"path/filepath"
)

func ZipFolder(folder string, destination *os.File) error {
	w := zip.NewWriter(destination)

	if err := addFiles(w, folder, nil); err != nil {
		return err
	}

	err := w.Close()
	if err != nil {
		return err
	}

	return nil
}

func addFiles(w *zip.Writer, basePath string, ignore []string) error {
	return filepath.Walk(basePath, func(path string, info fs.FileInfo, err error) error {
		if info.IsDir() {
			return nil
		}

		file, err := os.Open(path)
		if err != nil {
			return err
		}

		f, err := w.Create(path[len(basePath)+1:])
		if err != nil {
			return err
		}

		if _, err = io.Copy(f, file); err != nil {
			return err
		}

		return nil
	})
}
