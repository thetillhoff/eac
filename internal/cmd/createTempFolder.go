package cmd

import (
	"io/ioutil"
	"log"

	"github.com/thetillhoff/eac/pkg/logs"
)

func createTempFolder() string {
	// Creating tempfolder
	tempFolder, err := ioutil.TempDir("", "eac-*")
	if err != nil {
		log.Fatal(err)
	}
	if err != nil {
		logs.Err("Failed to create temporary folder:", err)
	}
	logs.Info("Created temporary folder at " + tempFolder)

	return tempFolder
}
