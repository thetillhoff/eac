package apps

import (
	"github.com/thetillhoff/eac/pkg/apps/internal/appVersions"
	"github.com/thetillhoff/eac/pkg/logs"
)

// Checks whether updates for the provided apps are available.
// If yes, only the version is updated, not the app.
// TODO If yes, the user is asked whether only the version should be updated or the app should be upgraded as well.
// If no name is provided, all apps are checked.
func Update(appNames []string, appsDirPath string, versionsFilePath string, dryRun bool, skipLocal bool, verbose bool) {
	logs.Verbose = verbose
	apps := parseApps(appNames, versionsFilePath)

	for _, appItem := range apps {
		_ = appVersions.Update(appItem)
	}
}
