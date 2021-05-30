package apps

import (
	"runtime"

	"github.com/thetillhoff/eac/pkg/logs"
)

func Update(appNames []string, shell string, appsDirPath string, continueOnError bool, versionsFilePath string, dryRun bool, skipLocal bool) {

	apps := apps(appNames, shell, continueOnError)

	for _, appItem := range apps {
		err := updateAppVersion(appItem, appsDirPath, runtime.GOOS, versionsFilePath, dryRun, skipLocal)
		if err != nil {
			logs.Err("There was an error while updating the version for app '"+appItem.Name+"':", err)
		}
	}
}
