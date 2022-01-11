package apps

import (
	"io/ioutil"

	"github.com/thetillhoff/eac/internal/config"
	"github.com/thetillhoff/eac/pkg/apps/internal/app"
	"github.com/thetillhoff/eac/pkg/logs"
)

// Returns a list of all apps that are managed via `eac` (== contained in `versions.yaml`).
// Per default one app per line in the format `<app>[==<installedVersion>]`.
// The seperator can be edited with a flag and is returned as the second parameter.
// Created by trying to run the getInstalledVersion script.
func List(conf config.Config) ([]app.App, config.Config) {
	var (
		appNames []string = []string{}
	)

	logs.Verbose = conf.Verbose

	// Get contents of appsDir
	appsDirContents, err := ioutil.ReadDir(conf.AppsDirPath)
	if err != nil {
		logs.Error("There was an error while reading from appsDir at '"+conf.AppsDirPath+"'. Did you run `eac init` already?", err)
	}

	for _, contentItem := range appsDirContents {
		appNames = append(appNames, contentItem.Name())
	}
	apps := parseApps(appNames, conf.VersionsFilePath) // continueOnError set to false, as this the whole command won't work then

	return apps, conf
}
