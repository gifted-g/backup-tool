package backup

import (
	"archive/zip"
	"io"
	"os"
	"path/filepath"
)

func RestoreBackup(zipPath, destination string) error {
	reader, err := zip.OpenReader(zipPath)
	if err != nil {
		return err
	}
	defer reader.Close()

	for _, file := range reader.File {
		fPath := filepath.Join(destination, file.Name)

		if file.FileInfo().IsDir() {
			err = os.MkdirAll(fPath, os.ModePerm)
			if err != nil {
				return err
			}
			continue
		}

		err = os.MkdirAll(filepath.Dir(fPath), os.ModePerm)
		if err != nil {
			return err
		}

		dstFile, err := os.Create(fPath)
		if err != nil {
			return err
		}
		defer dstFile.Close()

		srcFile, err := file.Open()
		if err != nil {
			return err
		}
		defer srcFile.Close()

		_, err = io.Copy(dstFile, srcFile)
		if err != nil {
			return err
		}
	}

	return nil
}
