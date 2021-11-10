package apps

import (
	"github.com/thetillhoff/eac/internal/appVersions"
	"github.com/thetillhoff/eac/pkg/logs"
)

func Update(appNames []string, appsDirPath string, versionsFilePath string, dryRun bool, skipLocal bool, verbose bool) {
	logs.Verbose = verbose
	apps := apps(appNames, versionsFilePath)

	for _, appItem := range apps {
		_ = appVersions.Update(appItem)
	}
}
