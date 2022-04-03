package goCopy

import (
	"errors"
	"fmt"
	"io"
	"io/fs"
	"os"
	"path"
)

// Copies sourceFile to destinationFile, while preserving permissions
func CopyFile(sourceFile string, destinationFile string) error {
	var (
		err         error
		sourceInfo  fs.FileInfo
		source      *os.File
		destination *os.File
	)

	if Verbose {
		fmt.Println("Copying ", sourceFile, "to", destinationFile)
	}

	source, err = os.Open(sourceFile)
	if err != nil {
		return errors.New("could not open source file at " + sourceFile + ": " + err.Error())
	}
	defer source.Close()

	sourceInfo, err = os.Stat(sourceFile) // Retrieve information about sourceFile (permissions)
	if err != nil {
		return errors.New("could not read permissions of source file at " + sourceFile + ": " + err.Error())
	}

	err = os.MkdirAll(path.Dir(destinationFile), os.ModeDir)
	if err != nil {
		return errors.New("could not create required destination folder/s: " + err.Error())
	}

	destination, err = os.Create(destinationFile)
	if err != nil {
		return errors.New("could not create destination file at " + destinationFile + ": " + err.Error())
	}
	defer destination.Close()

	_, err = io.Copy(destination, source)
	if err != nil {
		return errors.New("could not copy source to destination: " + err.Error())
	}

	err = destination.Sync()
	if err != nil {
		return errors.New("could not finish writing to destination: " + err.Error())
	}

	err = os.Chmod(destinationFile, sourceInfo.Mode())
	if err != nil {
		return errors.New("could not set destination permissions: " + err.Error())
	}

	return nil
}
