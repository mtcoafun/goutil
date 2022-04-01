// package zip provides zip compression/decompression functions.
package zip

import (
	"archive/zip"
	"io"
	"os"
)

// Zip files to a single zip file.
func Zip(dest string, files ...string) error {
	zipfile, err := os.Create(dest)
	if err != nil {
		return err
	}
	defer zipfile.Close()

	archive := zip.NewWriter(zipfile)
	defer archive.Close()

	for _, file := range files {
		if err := addFileToZip(file, archive); err != nil {
			return err
		}
	}
	return nil
}

func addFileToZip(filename string, archive *zip.Writer) error {
	file, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	info, err := file.Stat()
	if err != nil {
		return err
	}

	header, err := zip.FileInfoHeader(info)
	if err != nil {
		return err
	}

	header.Method = zip.Deflate

	writer, err := archive.CreateHeader(header)
	if err != nil {
		return err
	}
	_, err = io.Copy(writer, file)
	return err
}
