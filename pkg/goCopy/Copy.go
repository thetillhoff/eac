package goCopy

import (
	"io/fs"
	"os"
)

var (
	Verbose = false
)

// Copies source to destination.
// Can copy files and/or folders.
// Folders are copied recursively.
// Permissions are preserved.
func Copy(sourcePath string, destinationPath string) error {
	var (
		err        error
		sourceInfo fs.FileInfo
	)

	sourceInfo, err = os.Stat(sourcePath) // Retrieve information about sourcePath
	if err != nil {
		return err
	}

	if sourceInfo.IsDir() {
		return CopyDir(sourcePath, destinationPath)
	} else {
		return CopyFile(sourcePath, destinationPath)
	}
}
