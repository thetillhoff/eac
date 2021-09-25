package apps

import (
	"github.com/thetillhoff/eac/pkg/logs"
)

func Validate(appNames []string, flaggedPlatforms []string, appsDirPath string, verbose bool, versionsFilePath string) {
	logs.Verbose = verbose
	apps := apps(appNames, versionsFilePath)

	platforms := ResolvePlatforms(flaggedPlatforms)

	for _, appItem := range apps {
		for _, platform := range platforms {
			out, err := appItem.Validate(appsDirPath, platform)
			logs.Info("Output of validation:", out)
			if err != nil {
				logs.Err("There was an error while validating app '"+appItem.Name+"':", err)
			} else {
				logs.Success("Validation of app '" + appItem.Name + "' successful.")
			}
		}
	}
}
