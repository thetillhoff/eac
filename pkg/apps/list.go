package apps

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

func List(appsDirPath string, versionsFilePath string, noVersion bool, seperator string) {
	loadVersions(versionsFilePath)

	files, err := ioutil.ReadDir(appsDirPath)
	if err != nil {
		log.Fatalln("There was an error while reading from appsDir at '" + appsDirPath + "':\n" + err.Error())
	}

	items := []string{}
	for _, file := range files {
		if noVersion {
			items = append(items, file.Name())
		} else {
			items = append(items, file.Name()+"=="+getVersion(file.Name()))
		}
		fmt.Println(strings.Join(items, seperator))
	}
}
