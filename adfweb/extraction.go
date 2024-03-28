package adfweb

import (
	// "archive/tar"
	// "compress/gzip"
	// "fmt"
	// "io"
	"archive/zip"
	"io"
	"os"
	"path/filepath"

	"github.com/schollz/progressbar/v3"
)

func extractZipFile(zipPath, destPath string) error {
	// Open the zip file for reading
	r, err := zip.OpenReader(zipPath)
	if err != nil {
		return err
	}
	defer r.Close()

	// Calculate the total size of all the files in the zip
	var totalSize int64
	for _, file := range r.File {
		totalSize += file.FileInfo().Size()
	}

	// Create the progress bar
	// bar := progressbar.DefaultBytes(
	// 	totalSize,
	// 	"Extraindo arquivos",
	// )
	bar := progressbar.NewOptions64(totalSize, ProgressBarOptions...)
	bar.Describe("Extraindo arquivos")

	// Extract each file in the zip
	for _, file := range r.File {
		// Open the file inside the zip
		rc, err := file.Open()
		if err != nil {
			return err
		}
		defer rc.Close()

		// Create the corresponding file in the destination path
		path := filepath.Join(destPath, file.Name)
		if file.FileInfo().IsDir() {
			// Create directory if it doesn't exist
			err = os.MkdirAll(path, os.ModePerm)
			if err != nil {
				return err
			}
			continue
		}
		f, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, file.Mode())
		if err != nil {
			return err
		}

		// Copy the contents of the file to the destination
		_, err = io.Copy(f, rc)
		f.Close()
		if err != nil {
			return err
		}

		// Update the progress bar with the file size
		bar.Add64(file.FileInfo().Size())
	}

	// Finish the progress bar
	bar.Finish()

	return nil
}
