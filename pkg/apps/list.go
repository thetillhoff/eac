package apps

import (
	"io/ioutil"
	"strings"

	"github.com/thetillhoff/eac/pkg/logs"
)

func List(appsDirPath string, versionsFilePath string, noVersion bool, seperator string, verbose bool) {
	logs.Verbose = verbose
	loadVersions(versionsFilePath, false) // continueOnError set to false, as this whole command won't work then

	files, err := ioutil.ReadDir(appsDirPath)
	if err != nil {
		logs.Err("There was an error while reading from appsDir at '"+appsDirPath+"':", false, err) // continueOnError set to false, as this the whole command won't work then
	}

	items := []string{}
	for _, file := range files {
		if noVersion {
			items = append(items, file.Name())
		} else {
			items = append(items, file.Name()+"=="+getVersion(file.Name()))
		}
		logs.Info(strings.Join(items, seperator))
	}
}
