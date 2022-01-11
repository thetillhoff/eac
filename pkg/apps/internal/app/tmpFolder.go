package app

import (
	"os"
	"path"
	"strconv"
	"time"

	"github.com/thetillhoff/eac/pkg/logs"
)

func createTmpFolder() string {
	tmpFolder := path.Join(os.TempDir(), strconv.FormatInt(time.Now().UnixNano(), 10))

	err := os.Mkdir(tmpFolder, os.ModePerm)
	if err != nil {
		logs.Error("Failed to create temporary folder:", err)
	}
	logs.Info("Created temporary folder at " + tmpFolder)

	return tmpFolder
}

func deleteTmpFolder(tmpFolder string) {
	err := os.RemoveAll(tmpFolder)
	if err != nil {
		logs.Error("Failed to delete temporary folder:", err)
	}
	logs.Info("Deleted temporary folder at " + tmpFolder)
}
