package appVersions

import (
	"fmt"
	"runtime"

	"github.com/eiannone/keyboard"
	"github.com/thetillhoff/eac/pkg/apps/internal/app"
	"github.com/thetillhoff/eac/pkg/logs"
)

func Update(app app.App) app.App {
	logs.Info("Checking latest " + app.Name + " version ...")
	latestVersion := app.LatestVersion(appsDirPath, runtime.GOOS) // retrieve latest Version
	if latestVersion == "" {
		logs.Error("Retrieval of latest version for app '"+app.Name+"' failed.", app)
	}
	logs.Info("Latest version for app '" + app.Name + "' retrieved:  'v" + latestVersion + "'.")

	if app.WantedVersion != "" { // if VersionsFile contains a version for app
		if app.WantedVersion == latestVersion { // if version is already latest
			logs.Info("Version of app '" + app.Name + "' is already set to latest version (v" + app.WantedVersion + ").")
		} else { // if version is set but not latest: ask to upgrade
			fmt.Println("Do you want to update the version of app '" + app.Name + "' from 'v" + app.WantedVersion + "' to latest 'v" + latestVersion + "'? [y/n] ")
			char, _, err := keyboard.GetSingleKey()
			if err != nil {
				logs.Error("There was an error while reading your answer.", err)
			}
			if char == 'y' {
				versions[app.Name] = latestVersion
				app.WantedVersion = latestVersion
				Save(VersionsFilePath) // write versions in versionsFile
				logs.Info("Updated version of app '" + app.Name + "' from 'v" + app.InstalledVersion(AppsDirPath) + "' to 'v" + latestVersion + "'.")
			} else {
				logs.Info("Skipped app '" + app.Name + "'.")
			}
		}
	} else { // if versionsFile doesn't contain a version for app: ask to store
		fmt.Println("Do you want to set the version of app '" + app.Name + "' to latest v" + latestVersion + "? [y/n] ")
		char, _, err := keyboard.GetSingleKey()
		if err != nil {
			logs.Error("There was an error while reading your answer.", err)
		}
		if char == 'y' {
			versions[app.Name] = latestVersion
			app.WantedVersion = latestVersion
			logs.Info("Persisted updated version of app '" + app.Name + "' to 'v" + latestVersion + "'.")
		} else {
			logs.Info("Skipped app '" + app.Name + "'.")
		}
	}
	return app
}
