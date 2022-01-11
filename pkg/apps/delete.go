package apps

import (
	"os"
	"path"

	"github.com/thetillhoff/eac/pkg/logs"
)

// Deletes the files and folder structure for the provided list of appnames
func Delete(appNames []string, flaggedPlatforms []string, appsDirPath string, verbose bool) {
	logs.Verbose = verbose
	for _, appName := range appNames { // for all apps
		appPath := path.Join(appsDirPath, appName)
		if _, err := os.Stat(appPath); os.IsNotExist(err) { // if folder doesn't exist
			logs.Error("Folder '" + appPath + "' for app '" + appName + "' doesn't exist.")
		} else if err != nil {
			logs.Error("There was a problem while accessing app '"+appName+"':", err)
		}

		if len(flaggedPlatforms) == 0 { // delete whole app
			err := os.RemoveAll(appPath)
			if err != nil {
				logs.Error("There was an error while deleting the app at '"+appPath+"':", err)
			}
			logs.Success("Deleted app '" + appName + "'.")
		} else { // for specific platforms
			// platforms := ResolvePlatforms(flaggedPlatforms)
			platforms := []string{}

			for _, platform := range flaggedPlatforms { // validate specified platforms
				if !IsValidPlatform(platform) { // references method 'isValidPlatform' in 'create.go'
					logs.Error("Unkown platform: " + platform)
				}
				platforms = append(platforms, platform)
			}
			logs.Info("Selected platforms:", platforms)

			for _, platform := range platforms {
				platformPath := path.Join(appPath, platform)
				if _, err := os.Stat(platformPath); os.IsNotExist(err) { // if folder doesn't exist
					logs.Error("Folder '" + platformPath + "' for platform '" + platform + "' for app '" + appName + "' doesn't exist.")
				} else if err != nil {
					logs.Error("There was a problem while accessing platform '"+platform+"' for app '"+appName+"':", err)
				}
				err := os.RemoveAll(platformPath)
				if err != nil {
					logs.Error("There was a problem while deleting app '"+appName+"':", err)
				}
				logs.Success("Deleted platform '" + platformPath + "' for app '" + appName + "'.")
			}
		}
	}
}
