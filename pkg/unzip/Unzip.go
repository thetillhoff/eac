package unzip

import (
	"archive/zip"
	"fmt"
	"io"
	"os"
	"path"
)

var (
	Verbose bool = false
)

func Unzip(sourcefile string) error {

	var (
		err              error
		file             *zip.File
		zipReader        *zip.ReadCloser
		destination      = path.Dir(sourcefile)
		destinationFile  *os.File
		sourceFileReader io.ReadCloser
	)

	if Verbose {
		fmt.Println("Unzipping", sourcefile, "...")
	}

	zipReader, err = zip.OpenReader(sourcefile)
	if err != nil {
		return err
	}
	defer zipReader.Close()

	// Extracting zipped files
	if Verbose {
		fmt.Println("Extracting...")
	}
	for _, file = range zipReader.File {
		// Get the filepath and set the destination directory
		filepath := path.Join(destination, file.Name)

		if file.FileInfo().IsDir() { // Handle folders
			if Verbose {
				fmt.Println("Creating folder", filepath, "...")
			}

			err = os.MkdirAll(filepath, os.ModePerm)
			if err != nil {
				return err
			}

			if Verbose {
				fmt.Println("Folder creation finished.")
			}
		} else { // Handle files
			if Verbose {
				fmt.Println("Creating containing folder for file", filepath, "...")
			}

			err = os.MkdirAll(path.Dir(filepath), os.ModePerm)
			if err != nil {
				return err
			}

			if Verbose {
				fmt.Println("Folder creation finished.")
				fmt.Println("Creating file", filepath, "...")
			}

			destinationFile, err = os.OpenFile(filepath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, file.Mode())
			if err != nil {
				return err
			}
			defer destinationFile.Close()

			sourceFileReader, err = file.Open()
			if err != nil {
				return err
			}
			defer sourceFileReader.Close()

			_, err = io.Copy(destinationFile, sourceFileReader)
			if err != nil {
				return err
			}

			if Verbose {
				fmt.Println("File creation finished.")
			}
		}
	}
	if Verbose {
		fmt.Println("Extraction finished.")
	}

	return nil
}
