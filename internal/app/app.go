package app

import (
	"log"
	"os"
	"os/exec"
	"path"
	"strings"
)

func (app App) WantedVersion() string {
	return app.wantedVersion
}

func (app App) LocalVersion() string {
	return app.localVersion
}

// This will either return the local version or an empty string (when the script for getting the local version failed, f.e. if the app is not installed)
func (app App) getLocalVersion(platform string) string {
	localVersion, err := RunScript(app.shell, app.Name, platform, app.getLocalVersionScript, app.continueOnError)
	if err != nil {
		log.Println("App '" + app.Name + "' is not installed.")
		localVersion = ""
	}
	return localVersion
}

func (app App) GetLatestVersion(platform string) string {
	latestVersion, err := RunScript(app.shell, app.Name, platform, app.getLatestVersionScript, app.continueOnError)
	if err != nil {
		log.Fatalln("There was an error while retrieving the latest version of app '" + app.Name + "':\n" + err.Error())
	}
	return latestVersion
}

func (app App) Install(platform string, latest bool) (string, error) { //TODO add parameter version of type string and pass it to the script as arg.
	if latest {
		latestVersion := app.GetLatestVersion(platform)
		if latestVersion == "" {
			log.Fatalln("There was an error while retrieving the latest versin of app '" + app.Name + "'.")
		}
		app.wantedVersion = latestVersion
	}
	return RunScript(app.shell, app.Name, platform, app.installScript, app.continueOnError)
}

func (app App) Configure(platform string) (string, error) {
	return RunScript(app.shell, app.Name, platform, app.configureScript, app.continueOnError)
}

func (app App) Uninstall(platform string) (string, error) {
	return RunScript(app.shell, app.Name, platform, app.uninstallScript, app.continueOnError)
}

func RunScript(shell string, appName string, platform string, script string, appcontinueOnError bool) (string, error) {
	cmd := exec.Command(shell, path.Join("apps", appName, script))
	outbytes, err := cmd.CombinedOutput()
	out := strings.TrimSuffix(string(outbytes), "\n")
	return out, err
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
