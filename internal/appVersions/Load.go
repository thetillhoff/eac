package appVersions

import (
	"io/ioutil"

	"github.com/thetillhoff/eac/pkg/logs"
	"gopkg.in/yaml.v3"
)

func Load(versionFilePath string) {
	//TODO: allow multiple version files to be passed to the commands; -f, --values // -> flag in cmd/root.go, but functionality here

	var mapObject map[string]interface{}
	fileContents, err := ioutil.ReadFile(versionFilePath)
	if err != nil {
		logs.Err("There was an error while accessing the versionsFile at '"+versionFilePath+"':", err)
	}
	err = yaml.Unmarshal(fileContents, &mapObject)
	if err != nil {
		logs.Err("There was an error parsing the contents of the versionsfile:", err)
	}

	checkedMap := map[string]string{}

	for key, value := range mapObject {
		if value, ok := value.(string); ok {
			checkedMap[key] = value
		} else {
			logs.Err("The version '" + value + "' for app '" + key + "' is not a string.")
		}
	}

	versions = checkedMap
}
