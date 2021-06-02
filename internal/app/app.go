package app

import (
	"os"
	"runtime"

	"github.com/thetillhoff/eac/pkg/logs"
)

func (app App) WantedVersion() string {
	return app.wantedVersion
}

func (app App) LocalVersion(appsDirPath string) string {
	if app.localVersion == "" && !app.localVersionFailed {
		app.localVersion = app.getLocalVersion(appsDirPath)
	}
	return app.localVersion
}

// This will either return the local version or an empty string (when the script for getting the local version failed, f.e. if the app is not installed)
func (app App) getLocalVersion(appsDirPath string) string {
	localVersion, err := RunScript(app.shell, app.Name, appsDirPath, runtime.GOOS, app.getLocalVersionScript, app.continueOnError)
	if err != nil {
		logs.Warn("App '" + app.Name + "' is not installed or the getLocalVersion script doesn't work properly.")
		localVersion = ""
	}
	return localVersion
}

func (app App) GetLatestVersion(appsDirPath string, platform string) string {
	latestVersion, err := RunScript(app.shell, app.Name, appsDirPath, platform, app.getLatestVersionScript, app.continueOnError)
	if err != nil {
		logs.Err("There was an error while retrieving the latest version of app '"+app.Name+"':", app.continueOnError, err)
	}
	return latestVersion
}

func (app App) Install(appsDirPath string, platform string, version string) (App, string, error) {
	if version == "" && app.WantedVersion() == "" {
		latestVersion := app.GetLatestVersion(appsDirPath, platform)
		if latestVersion == "" {
			logs.Err("There was an error while retrieving the latest version of app '"+app.Name+"'.", app.continueOnError)
		}
		app.wantedVersion = latestVersion
	}

	out, err := RunScript(app.shell, app.Name, appsDirPath, platform, app.installScript, app.continueOnError, app.WantedVersion())

	return app, out, err
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
