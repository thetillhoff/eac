package app

import (
	"os"

	"github.com/thetillhoff/eac/pkg/logs"
)

func (app App) WantedVersion() string {
	return app.wantedVersion
}

func (app App) LocalVersion() string {
	return app.localVersion
}

// This will either return the local version or an empty string (when the script for getting the local version failed, f.e. if the app is not installed)
func (app App) getLocalVersion(appsDirPath string, platform string) string {
	localVersion, err := RunScript(app.shell, app.Name, appsDirPath, platform, app.getLocalVersionScript, app.continueOnError)
	if err != nil {
		logs.Warn("App '" + app.Name + "' is not installed.")
		localVersion = ""
	}
	return localVersion
}

func (app App) GetLatestVersion(appsDirPath string, platform string) string {
	latestVersion, err := RunScript(app.shell, app.Name, appsDirPath, platform, app.getLatestVersionScript, app.continueOnError)
	if err != nil {
		logs.Err("There was an error while retrieving the latest version of app '"+app.Name+"':", err)
	}
	return latestVersion
}

func (app App) Install(appsDirPath string, platform string, latest bool) (string, error) { //TODO add parameter version of type string and pass it to the script as arg.
	if latest || app.WantedVersion() == "" {
		latestVersion := app.GetLatestVersion(appsDirPath, platform)
		if latestVersion == "" {
			logs.Err("There was an error while retrieving the latest version of app '" + app.Name + "'.")
		}
		app.wantedVersion = latestVersion
	}

	return RunScript(app.shell, app.Name, appsDirPath, platform, app.installScript, app.continueOnError, app.WantedVersion())
}

func (app App) Configure(appsDirPath string, platform string) (string, error) {
	return RunScript(app.shell, app.Name, appsDirPath, platform, app.configureScript, app.continueOnError)
}

func (app App) Uninstall(appsDirPath string, platform string) (string, error) {
	return RunScript(app.shell, app.Name, appsDirPath, platform, app.uninstallScript, app.continueOnError)
}

func testFiles(filepaths ...string) []string {

	missingFiles := []string{}

	for _, filepath := range filepaths {
		if _, err := os.Stat(filepath); os.IsNotExist(err) {
			missingFiles = append(missingFiles, filepath)
		}
	}

	return missingFiles
}
