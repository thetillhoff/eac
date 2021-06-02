package app

import (
	"os"
	"path"
	"strconv"
	"time"

	"github.com/thetillhoff/eac/pkg/logs"
)

func createTmpFolder(continueOnError bool) string {
	tmpFolder := path.Join(os.TempDir(), strconv.FormatInt(time.Now().UnixNano(), 10))

	err := os.Mkdir(tmpFolder, os.ModePerm)
	if err != nil {
		logs.Err("Failed to create temporary folder:", continueOnError, err)
	}
	logs.Info("Created temporary folder at " + tmpFolder)

	return tmpFolder
}

func deleteTmpFolder(tmpFolder string, continueOnError bool) {
	err := os.RemoveAll(tmpFolder)
	if err != nil {
		logs.Err("Failed to delete temporary folder:", continueOnError, err)
	}
	logs.Info("Deleted temporary folder at " + tmpFolder)
}
