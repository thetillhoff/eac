package untar

import (
	"archive/tar"
	"compress/gzip"
	"errors"
	"fmt"
	"io"
	"os"
	"path"
)

var (
	Verbose bool = false
)

func Untar(sourcefile string) error {

	var (
		err         error
		destination = path.Dir(sourcefile)
		file        *os.File
	)

	if Verbose {
		fmt.Println("Untarring", sourcefile, "...")
	}

	file, err = os.Open(sourcefile)
	if err != nil {
		return err
	}
	defer file.Close()

	var fileReader io.ReadCloser = file

	// In case we are reading a .tar.gz file, add a filter to handle gzipped file
	if path.Ext(sourcefile) == ".gz" {
		if Verbose {
			fmt.Println("Detected gzipped file, adjusting...")
		}
		if fileReader, err = gzip.NewReader(file); err != nil {
			return err
		}
		defer fileReader.Close()
		if Verbose {
			fmt.Println("Adjustment finished.")
		}
	}

	tarBallReader := tar.NewReader(fileReader)

	// Extracting tarred files
	if Verbose {
		fmt.Println("Extracting...")
	}
	for {
		header, err := tarBallReader.Next()
		if err != nil {
			if err == io.EOF {
				break
			}
			return err
		}

		// Get the filepath and set the destination directory
		filepath := path.Join(destination, header.Name)

		switch header.Typeflag {
		case tar.TypeDir: // Handle folders
			if Verbose {
				fmt.Println("Creating folder", filepath, "...")
			}
			err = os.MkdirAll(filepath, os.FileMode(header.Mode))
			if err != nil {
				return err
			}

		case tar.TypeReg: // Handle files
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
			writer, err := os.Create(filepath)
			if err != nil {
				return err
			}

			io.Copy(writer, tarBallReader)

			err = os.Chmod(filepath, os.FileMode(header.Mode))
			if err != nil {
				return err
			}

			writer.Close()
		default:
			return errors.New("unable to untar type '" + string(header.Typeflag) + "' in file '" + filepath + "'")
		}
	}
	if Verbose {
		fmt.Println("Extraction finished.")
	}

	return nil
}
