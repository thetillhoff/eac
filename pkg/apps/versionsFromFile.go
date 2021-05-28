package apps

import (
	"bufio"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"

	"github.com/thetillhoff/eac/internal/app"
	"gopkg.in/yaml.v3"
)

var (
	versionsFromfile map[string]string
)

func getVersion(appName string) string {
	return versionsFromfile[appName]
}

func loadVersions(filepath string) {
	//TODO: allow multiple version files to be passed to the commands; -f, --values // -> flag in cmd/root.go, but functionallity here

	var mapObject map[string]interface{}
	fileContents, err := ioutil.ReadFile(filepath)
	if err != nil {
		log.Fatalln("There was an error while accessing the versionsFile at '" + filepath + "':\n" + err.Error())
	}
	yaml.Unmarshal([]byte(fileContents), &mapObject)

	checkedMap := map[string]string{}

	for key, value := range mapObject {
		if value, ok := value.(string); ok {
			checkedMap[key] = value
		} else {
			log.Fatalln("The version '" + value + "' for app '" + key + "' is not a string.")
		}
	}

	versionsFromfile = checkedMap
}

func updateAppVersion(app app.App, platform string, versionsFilePath string, dryRun bool, skipLocal bool) error {
	fmt.Println("Checking " + app.Name + " version ...")
	if app.LocalVersion() == "" {
		return errors.New("Retrieval of local version for app '" + app.Name + "' failed.")
	}

	latestVersion := app.GetLatestVersion(platform)
	if latestVersion == "" {
		return errors.New("Retrieval of latest version for app '" + app.Name + "' failed.")
	}

	if app.LocalVersion() == latestVersion {
		fmt.Println("App '" + app.Name + "' is already installed with the latest version (v" + app.LocalVersion() + ").")
	} else {
		if !dryRun {
			reader := bufio.NewReader(os.Stdin)
			fmt.Print("Do you want to update from v" + app.LocalVersion() + " to v" + latestVersion + "? [y/n] ")
			text, err := reader.ReadString('\n')
			if err != nil {
				return err
			}
			text = strings.TrimSuffix(text, "\n") // remove line-break after input
			if text == "y" {
				versionsFromfile[app.Name] = latestVersion
				bFileContents, err := yaml.Marshal(&versionsFromfile)
				if err != nil {
					log.Fatalln("There was an error during the conversion of the updated versionsFile:\n" + err.Error())
				}
				f, err := os.Open(versionsFilePath)
				if err != nil {
					log.Fatalln("There was an error while opening the versionsFile at '" + versionsFilePath + "':\n" + err.Error())
				}
				defer f.Close()
				err = f.Truncate(0)
				if err != nil {
					log.Fatalln("There was an error while removing earlier contents in the versionsFile at '" + versionsFilePath + "':\n" + err.Error())
				}
				_, err = fmt.Fprint(f, bFileContents)
				if err != nil {
					log.Fatalln("There was an error while writing to the versionsFile at '" + versionsFilePath + "':\n" + err.Error())
				}
				//TODO: parse version file, update version of app, save yaml back to file
				fmt.Println("Updated from v" + app.LocalVersion() + " to v" + latestVersion + ".")
			} else {
				fmt.Println("Skipped " + app.Name + ".")
			}
		} else {
			fmt.Println("App '" + app.Name + "' could be upgraded from version '" + app.LocalVersion() + "' to version '" + latestVersion + "'.")
		}
	}
	return nil
}
