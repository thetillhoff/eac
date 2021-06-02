package apps

import (
	"bufio"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/thetillhoff/eac/internal/app"
	"github.com/thetillhoff/eac/pkg/logs"
	"gopkg.in/yaml.v3"
)

var (
	versionsFromfile map[string]string
)

func getVersion(appName string) string {
	return versionsFromfile[appName]
}

func loadVersions(filepath string, continueOnError bool) {
	//TODO: allow multiple version files to be passed to the commands; -f, --values // -> flag in cmd/root.go, but functionallity here

	var mapObject map[string]interface{}
	fileContents, err := ioutil.ReadFile(filepath)
	if err != nil {
		logs.Err("There was an error while accessing the versionsFile at '"+filepath+"':", continueOnError, err)
	}
	yaml.Unmarshal([]byte(fileContents), &mapObject)

	checkedMap := map[string]string{}

	for key, value := range mapObject {
		if value, ok := value.(string); ok {
			checkedMap[key] = value
		} else {
			logs.Err("The version '"+value+"' for app '"+key+"' is not a string.", continueOnError)
		}
	}

	versionsFromfile = checkedMap
}

func updateAppVersion(app app.App, appsDirPath, platform string, continueOnError bool, versionsFilePath string, dryRun bool, skipLocal bool) error {
	logs.Info("Checking " + app.Name + " version ...")
	if app.LocalVersion(appsDirPath) == "" {
		return errors.New("Retrieval of local version for app '" + app.Name + "' failed.")
	}

	latestVersion := app.GetLatestVersion(appsDirPath, platform)
	if latestVersion == "" {
		return errors.New("Retrieval of latest version for app '" + app.Name + "' failed.")
	}

	if app.LocalVersion(appsDirPath) == latestVersion {
		logs.Info("App '" + app.Name + "' is already installed with the latest version (v" + app.LocalVersion(appsDirPath) + ").")
	} else {
		if !dryRun {
			reader := bufio.NewReader(os.Stdin)
			logs.Info("Do you want to update from v" + app.LocalVersion(appsDirPath) + " to v" + latestVersion + "? [y/n] ")
			text, err := reader.ReadString('\n')
			if err != nil {
				return err
			}
			text = strings.TrimSuffix(text, "\n") // remove line-break after input
			if text == "y" {
				versionsFromfile[app.Name] = latestVersion
				bFileContents, err := yaml.Marshal(&versionsFromfile)
				if err != nil {
					logs.Err("There was an error during the conversion of the updated versionsFile:", continueOnError, err)
				}
				f, err := os.Open(versionsFilePath)
				if err != nil {
					logs.Err("There was an error while opening the versionsFile at '"+versionsFilePath+"':", continueOnError, err)
				}
				defer f.Close()
				err = f.Truncate(0)
				if err != nil {
					logs.Err("There was an error while removing earlier contents in the versionsFile at '"+versionsFilePath+"':", continueOnError, err)
				}
				_, err = fmt.Fprint(f, bFileContents)
				if err != nil {
					logs.Err("There was an error while writing to the versionsFile at '"+versionsFilePath+"':", continueOnError, err)
				}
				//TODO: parse version file, update version of app, save yaml back to file
				logs.Info("Updated from v" + app.LocalVersion(appsDirPath) + " to v" + latestVersion + ".")
			} else {
				logs.Info("Skipped " + app.Name + ".")
			}
		} else {
			logs.Info("App '" + app.Name + "' could be upgraded from version '" + app.LocalVersion(appsDirPath) + "' to version '" + latestVersion + "'.")
		}
	}
	return nil
}
