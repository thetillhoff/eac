package cmd

import (
	"os"

	"github.com/thetillhoff/eac/pkg/logs"
)

func deleteTempFolder(tempFolder string) {
	err := os.RemoveAll(tempFolder)
	if err != nil {
		logs.Error("Failed to delete temporary folder:", err)
	}
	logs.Info("Deleted temporary folder at '" + tempFolder + "'.")
}
