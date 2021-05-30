package apps

import (
	"github.com/thetillhoff/eac/pkg/logs"
)

func Validate(appNames []string, flaggedPlatforms []string, shell string, appsDirPath string, continueOnError bool) {

	apps := apps(appNames, shell, continueOnError)

	platforms := ResolvePlatforms(flaggedPlatforms)

	for _, appItem := range apps {
		for _, platform := range platforms {
			out, err := appItem.Validate(appsDirPath, platform)
			logs.Info("Output of configuration script:", out)
			if err != nil {
				logs.Err("There was an error while validating app '"+appItem.Name+"':", err)
			}
		}
	}
}
