package app

import (
	"os"
	"runtime"
	"strings"

	"github.com/thetillhoff/eac/pkg/logs"
)

// This will either return the local version or an empty string (when the script for getting the local version failed, f.e. if the app is not installed)
func (app App) InstalledVersion(appsDirPath string) string {
	var (
		installedVersion string
		err              error
	)
	if app.installedVersion == "" {
		installedVersion, err = RunScript(app.Name, appsDirPath, runtime.GOOS, app.getInstalledVersionScript)
		if err != nil {
			logs.Warn("App '"+app.Name+"' is not installed or the getInstalledVersion script doesn't work properly.", installedVersion, err)
			installedVersion = ""
		}
		if strings.Contains(installedVersion, app.Name+": not found") {
			logs.Info("App '" + app.Name + "' is not installed.")
			installedVersion = ""
		}
		if app.installedVersion != installedVersion {
			app.installedVersion = installedVersion
			logs.Info("Updated installedVersion for app '" + app.Name + "'.")
		}
	}
	return app.installedVersion
}

func (app App) LatestVersion(appsDirPath string, platform string) string {
	var (
		latestVersion string
		err           error
	)
	if app.latestVersion == "" {
		latestVersion, err = RunScript(app.Name, appsDirPath, platform, app.getLatestVersionScript)
		if err != nil {
			logs.Error("There was an error while retrieving the latest version of app '"+app.Name+"':", err)
			latestVersion = ""
		}
		if latestVersion == "" {
			logs.Error("There was an error while retrieving the latest version of app '"+app.Name+"'.", latestVersion)
		} else if strings.Contains(latestVersion, "\n") {
			logs.Error("The latestVersion of app '"+app.Name+"' was not a single line:", latestVersion)
		}
		app.latestVersion = latestVersion
		logs.Info("Updated latestVersion for app '" + app.Name + "'.")
	}
	return app.latestVersion
}

func (app App) Install(appsDirPath string, platform string, version string) (App, string, error) {
	if version == "" && app.WantedVersion == "" {
		latestVersion := app.LatestVersion(appsDirPath, platform)
		if latestVersion == "" {
			logs.Error("There was an error while retrieving the latest version of app '" + app.Name + "'.")
		}
		app.WantedVersion = latestVersion
	}

	out, err := RunScript(app.Name, appsDirPath, platform, app.installScript, app.WantedVersion)

	return app, out, err
}

func (app App) Configure(appsDirPath string, platform string) (string, error) {
	return RunScript(app.Name, appsDirPath, platform, app.configureScript)
}

func (app App) Uninstall(appsDirPath string, platform string) (string, error) {
	return RunScript(app.Name, appsDirPath, platform, app.uninstallScript)
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
