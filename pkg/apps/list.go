package apps

import (
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/thetillhoff/eac/internal/config"
	"github.com/thetillhoff/eac/pkg/logs"
)

func List(conf config.Config) {
	var (
		appNames []string = []string{}
	)

	logs.Verbose = conf.Verbose

	// Get contents of appsDir
	appsDirContents, err := ioutil.ReadDir(conf.AppsDirPath)
	if err != nil {
		logs.Err("There was an error while reading from appsDir at '"+conf.AppsDirPath+"'. Did you run `eac init` already?", err)
	}

	for _, contentItem := range appsDirContents {
		appNames = append(appNames, contentItem.Name())
	}
	apps := apps(appNames, conf.VersionsFilePath) // continueOnError set to false, as this the whole command won't work then

	items := []string{}
	for _, appItem := range apps {
		if conf.NoVersion || appItem.LocalVersion(conf.AppsDirPath) == "" { // if no version should be displayed or no version is installed
			items = append(items, appItem.Name)
		} else {
			items = append(items, appItem.Name+"=="+appItem.LocalVersion(conf.AppsDirPath))
		}
	}
	fmt.Println(strings.Join(items, conf.Seperator))
}
