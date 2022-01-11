package apps

import (
	"github.com/thetillhoff/eac/pkg/logs"
)

// Checks whether the app configurations are set up in a valid way. TODO what exactly is validated here?
func Validate(appNames []string, flaggedPlatforms []string, appsDirPath string, verbose bool, versionsFilePath string) {
	logs.Verbose = verbose
	apps := parseApps(appNames, versionsFilePath)

	platforms := ResolvePlatforms(flaggedPlatforms)
	logs.Info("Selected platforms:", platforms)

	for _, appItem := range apps {
		for _, platform := range platforms {
			out, err := appItem.Validate(appsDirPath, platform)
			logs.Info("Output of validation:", out)
			if err != nil {
				logs.Error("There was an error while validating app '"+appItem.Name+"':", err)
			} else {
				logs.Success("Validation of app '" + appItem.Name + "' successful.")
			}
		}
	}
}
