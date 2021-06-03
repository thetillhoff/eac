package apps

import (
	"fmt"
	"io/fs"
	"io/ioutil"
	"os"
	"path"

	"github.com/thetillhoff/eac/pkg/logs"
)

func Create(appNames []string, flaggedPlatforms []string, shell string, appsDirPath string, verbose bool) {
	logs.Verbose = verbose
	platforms := ResolvePlatforms(flaggedPlatforms)

	if _, err := os.Stat(appsDirPath); os.IsNotExist(err) {
		logs.Err("Apps folder at '" + appsDirPath + "' doesn't exist.\nPlease run 'eac init' first.")
	} else if err != nil {
		logs.Err("There was an error while accessing the appsDir at '"+appsDirPath+"':", err)
	}

	for _, appName := range appNames {
		appPath := path.Join(appsDirPath, appName)
		if _, err := os.Stat(appPath); os.IsNotExist(err) { // if folder doesn't exist yet
			err := os.Mkdir(appPath, os.ModePerm)
			if err != nil {
				logs.Err("There was an error while creating the app at '"+appPath+"':", err)
			}
			logs.Success("Created app '" + appName + "'.")
		} else if err == nil { // if folder does exist
			logs.Err("App '" + appName + "' does already exist.")
		} else {
			logs.Err("There was an error while accessing the app at '"+appPath+"':", err)
		}

		for _, platform := range platforms {
			platformPath := path.Join(appPath, platform)
			if _, err := os.Stat(platformPath); os.IsNotExist(err) {
				err := os.Mkdir(platformPath, os.ModePerm) // ignore errors
				if err != nil {
					logs.Err("There was an error while creating the platform '"+platform+"' for app '"+appPath+"' at '"+platformPath+"':", err)
				}
				logs.Info("Created '" + platformPath + "' folder.")
			} else if err == nil {
				logs.Warn("Platform '" + platform + "' for app '" + appName + "' does already exist.")
			} else {
				logs.Err("There was an error while accessing the platform '"+platform+"' for app '"+appPath+"' at '"+platformPath+"':", err)
			}

			platformDemoFiles := demoFiles[platform]

			for filename, fileContent := range platformDemoFiles {

				fileContent = fmt.Sprintf(fileContent, appName)

				err := ioutil.WriteFile(path.Join(platformPath, filename), []byte(fileContent), fs.ModePerm)
				if err != nil {
					logs.Err("There was an error while writing to the file '"+filename+"' at '"+platformPath+"':", err)
				}

				logs.Info("Created '" + path.Join(platformPath, filename) + "' file.")
			}
			logs.Success("Created platform '" + platform + "' for app '" + appName + "'.")
		}
	}
}
