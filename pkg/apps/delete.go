package apps

import (
	"os"
	"path"

	"github.com/thetillhoff/eac/pkg/logs"
)

func Delete(appNames []string, flaggedPlatforms []string, shell string, appsDirPath string, continueOnError bool) {
	for _, appName := range appNames { // for all apps
		appPath := path.Join(appsDirPath, appName)
		if _, err := os.Stat(appPath); os.IsNotExist(err) { // if folder doesn't exist
			logs.Err("Folder '" + appPath + "' for app '" + appName + "' doesn't exist.")
		} else if err != nil {
			logs.Err("There was a problem while accessing app '"+appName+"':", err)
		}

		if len(flaggedPlatforms) == 0 { // delete whole app
			err := os.RemoveAll(appPath)
			if err != nil {
				logs.Err("There was an error while deleting the app at '"+appPath+"':", err)
			}
			logs.Info("Deleted '" + appPath + "' folder and all its contents.")
		} else { // for specific platforms
			platforms := []string{}

			for _, platform := range flaggedPlatforms { // validate specified platforms
				if !IsValidPlatform(platform) { // references method 'isValidPlatform' in 'create.go'
					logs.Err("Unkown platform: " + platform)
				}
				platforms = append(platforms, platform)

				for _, platform := range platforms {
					platformPath := path.Join(appPath, platform)
					if _, err := os.Stat(platformPath); os.IsNotExist(err) { // if folder doesn't exist
						logs.Err("Folder '" + platformPath + "' for platform '" + platform + "' for app '" + appName + "' doesn't exist.")
					} else if err != nil {
						logs.Err("There was a problem while accessing platform '"+platform+"' for app '"+appName+"':", err)
					}
					err := os.RemoveAll(platformPath)
					if err != nil {
						logs.Err("There was a problem while deleting app '"+appName+"':", err)
					}
					logs.Info("Deleted '" + platformPath + "' folder.")
				}
			}
		}
	}
}
