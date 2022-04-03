package goCopy

import (
	"fmt"
	"io/fs"
	"io/ioutil"
	"os"
	"path"
)

// Recursively copy sourceFolder to destinationFolder while preserving permissions
func CopyDir(sourceFolder string, destinationFolder string) error {
	var (
		err             error
		contents        []fs.FileInfo
		sourceInfo      fs.FileInfo
		sourcePath      string
		destinationPath string
	)

	if Verbose {
		fmt.Println("Copying ", sourceFolder, "to", destinationFolder)
	}

	sourceInfo, err = os.Stat(sourceFolder) // Retrieve information about sourceFolder (permissions)
	if err != nil {
		return err
	}

	err = os.MkdirAll(destinationFolder, sourceInfo.Mode()) // Create destinationFolder
	if err != nil {
		return err
	}

	contents, err = ioutil.ReadDir(sourceFolder)
	if err != nil {
		return err
	}

	for _, entry := range contents {
		sourcePath = path.Join(sourceFolder, entry.Name())
		destinationPath = path.Join(destinationFolder, entry.Name())

		sourceInfo, err := os.Stat(sourcePath) // Retrieve information about sourcePath (filetype)
		if err != nil {
			return err
		}

		if sourceInfo.IsDir() { // If sourcePath is Dir
			err = CopyDir(sourcePath, destinationPath)
			if err != nil {
				return err
			}
		} else { // If sourcePath is File
			err = CopyFile(sourcePath, destinationPath)
			if err != nil {
				return err
			}
		}
	}

	return nil
}
