package apps

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/eiannone/keyboard"
	"github.com/thetillhoff/eac/internal/app"
	"github.com/thetillhoff/eac/pkg/logs"
	"gopkg.in/yaml.v3"
)

var (
	versionsFromFile map[string]string
)

func getVersion(appName string) string {
	if value, ok := versionsFromFile[appName]; ok {
		return value
	} else {
		logs.Warn("No version for app '" + appName + "' in versionsFile yet.")
		logs.Info("versionsFile: ", versionsFromFile)
		return ""
	}
}

func loadVersions(filepath string) {
	//TODO: allow multiple version files to be passed to the commands; -f, --values // -> flag in cmd/root.go, but functionality here

	var mapObject map[string]interface{}
	fileContents, err := ioutil.ReadFile(filepath)
	if err != nil {
		logs.Err("There was an error while accessing the versionsFile at '"+filepath+"':", err)
	}
	err = yaml.Unmarshal(fileContents, &mapObject)
	if err != nil {
		logs.Err("There was an error parsing the contents of the versionsfile:", false, err)
	}

	checkedMap := map[string]string{}

	for key, value := range mapObject {
		if value, ok := value.(string); ok {
			checkedMap[key] = value
		} else {
			logs.Err("The version '" + value + "' for app '" + key + "' is not a string.")
		}
	}

	logs.Info("Read the following versions from versionsfile: ", checkedMap)

	versionsFromFile = checkedMap
}

func saveVersions(filepath string) {
	bFileContents, err := yaml.Marshal(&versionsFromFile)
	if err != nil {
		logs.Err("There was an error during the conversion of the updated versionsFile:", err)
	}
	f, err := os.OpenFile(filepath, os.O_WRONLY, os.ModePerm)
	if err != nil {
		logs.Err("There was an error while opening the versionsFile at '"+filepath+"':", err)
	}
	defer f.Close()
	_, err = f.WriteString(string(bFileContents))
	if err != nil {
		logs.Err("There was an error while writing to the versionsFile at '"+filepath+"':", err)
	}
}

func updateAppVersion(app app.App, appsDirPath, platform string, versionsFilePath string) app.App {
	logs.Info("Checking latest " + app.Name + " version ...")
	latestVersion := app.LatestVersion(appsDirPath, platform) // retrieve latest Version
	if latestVersion == "" {
		logs.Err("Retrieval of latest version for app '"+app.Name+"' failed.", app)
	}
	logs.Info("Latest version for app '" + app.Name + "' retrieved:  'v" + latestVersion + "'.")

	if app.WantedVersion != "" { // if VersionsFile contains a version for app
		if app.WantedVersion == latestVersion { // if version is already latest
			logs.Info("Version of app '" + app.Name + "' is already set to latest version (v" + app.WantedVersion + ").")
		} else { // if version is set but not latest: ask to upgrade
			fmt.Println("Do you want to update the version of app '" + app.Name + "' from 'v" + app.WantedVersion + "' to latest 'v" + latestVersion + "'? [y/n] ")
			char, _, err := keyboard.GetSingleKey()
			if err != nil {
				logs.Err("There was an error while reading your answer.", err)
			}
			if char == 'y' {
				versionsFromFile[app.Name] = latestVersion
				app.WantedVersion = latestVersion
				saveVersions(versionsFilePath) // write versions in versionsFile
				logs.Info("Updated version of app '" + app.Name + "' from 'v" + app.LocalVersion(appsDirPath) + "' to 'v" + latestVersion + "'.")
			} else {
				logs.Info("Skipped app '" + app.Name + "'.")
			}
		}
	} else { // if versionsFile doesn't contain a version for app: ask to store
		fmt.Println("Do you want to set the version of app '" + app.Name + "' to latest v" + latestVersion + "? [y/n] ")
		char, _, err := keyboard.GetSingleKey()
		if err != nil {
			logs.Err("There was an error while reading your answer.", err)
		}
		if char == 'y' {
			versionsFromFile[app.Name] = latestVersion
			app.WantedVersion = latestVersion
			saveVersions(versionsFilePath) // write versions in versionsFile
			logs.Info("Set version of app '" + app.Name + "' to 'v" + latestVersion + "'.")
		} else {
			logs.Info("Skipped app '" + app.Name + "'.")
		}
	}
	return app
}
