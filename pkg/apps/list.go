package apps

import (
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/thetillhoff/eac/pkg/logs"
)

func List(appsDirPath string, versionsFilePath string, noVersion bool, seperator string, verbose bool) {
	logs.Verbose = verbose
	appsDirContents, err := ioutil.ReadDir(appsDirPath)
	if err != nil {
		logs.Err("There was an error while reading from appsDir at '"+appsDirPath+"':", false, err) // continueOnError set to false, as this the whole command won't work then
	}
	appNames := []string{}
	for _, contentItem := range appsDirContents {
		appNames = append(appNames, contentItem.Name())
	}
	apps := apps(appNames, "", versionsFilePath) // continueOnError set to false, as this the whole command won't work then

	items := []string{}
	for _, appItem := range apps {
		if noVersion || appItem.LocalVersion(appsDirPath) == "" { // if no version should be displayed or no version is installed
			items = append(items, appItem.Name)
		} else {
			items = append(items, appItem.Name+"=="+appItem.LocalVersion(appsDirPath))
		}
	}
	fmt.Println(strings.Join(items, seperator))
}
