package apps

import (
	"runtime"

	"github.com/thetillhoff/eac/pkg/logs"
)

func Update(appNames []string, appsDirPath string, versionsFilePath string, dryRun bool, skipLocal bool, verbose bool) {
	logs.Verbose = verbose
	apps := apps(appNames, versionsFilePath)

	for _, appItem := range apps {
		_ = updateAppVersion(appItem, appsDirPath, runtime.GOOS, versionsFilePath)
	}
}
