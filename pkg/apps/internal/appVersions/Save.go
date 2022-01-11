package appVersions

import (
	"os"

	"github.com/thetillhoff/eac/pkg/logs"
	"gopkg.in/yaml.v3"
)

// Saves the current state of Versions to the versionsFile
func Save(versionsFilePath string) {
	bFileContents, err := yaml.Marshal(&versions)
	if err != nil {
		logs.Error("There was an error during the conversion of the updated versionsFile:", err)
	}
	f, err := os.OpenFile(versionsFilePath, os.O_WRONLY, os.ModePerm)
	if err != nil {
		logs.Error("There was an error while opening the versionsFile at '"+versionsFilePath+"':", err)
	}
	defer f.Close()
	_, err = f.WriteString(string(bFileContents))
	if err != nil {
		logs.Error("There was an error while writing to the versionsFile at '"+versionsFilePath+"':", err)
	}
}
