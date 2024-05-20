package services

import (
	"archive/tar"
	"archive/zip"
	"compress/gzip"
	"fmt"
	"io"
	"os"
	"path/filepath"

	"github.com/schollz/progressbar/v3"
)

func extractFile(filename, destinationPath string, progressBar []progressbar.Option) error {

	ext := filepath.Ext(filename)
	switch ext {
	case ".zip", ".war":
		return extractZipFile(filename, destinationPath, progressBar)
	case ".gz":
		return extractTarGz(filename, destinationPath)
	default:
		return fmt.Errorf("unsupported file format: %s", ext)
	}
}

func extractZipFile(filePath, destinationPath string, progressBar []progressbar.Option) error {
	// Open the zip file for reading
	r, err := zip.OpenReader(filePath)
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
	bar := progressbar.NewOptions64(totalSize, progressBar...)
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
		path := filepath.Join(destinationPath, file.Name)
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

func extractTarGz(filePath, destinationPath string) error {
	// Open the .tar.gz file
	file, err := os.Open(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	// Create a gzip reader
	gzipReader, err := gzip.NewReader(file)
	if err != nil {
		return err
	}
	defer gzipReader.Close()

	// Create a tar reader
	tarReader := tar.NewReader(gzipReader)

	// Extract each file from the tar archive
	for {
		header, err := tarReader.Next()
		if err == io.EOF {
			break // End of archive
		}
		if err != nil {
			return err
		}

		// Construct the full path for the extracted file
		targetPath := destinationPath + "/" + header.Name

		// Check if the file is a directory or a regular file
		switch header.Typeflag {
		case tar.TypeDir:
			// Create directories
			if err := os.MkdirAll(targetPath, os.ModePerm); err != nil {
				return err
			}
		case tar.TypeReg:
			// Create the file
			outFile, err := os.Create(targetPath)
			if err != nil {
				return err
			}
			defer outFile.Close()

			if err := os.Chmod(targetPath, 0777); err != nil {
				return fmt.Errorf("failed to set file permissions: %w", err)
			}

			// Copy the file contents
			if _, err := io.Copy(outFile, tarReader); err != nil {
				return err
			}
		}
	}

	return nil
}
