package apps

import (
	"github.com/thetillhoff/eac/internal/config"
	"github.com/thetillhoff/eac/pkg/apps/internal/appVersions"
	"github.com/thetillhoff/eac/pkg/logs"
)

// Checks whether updates for the provided apps are available.
// If yes, only the version is updated, not the app.
// TODO If yes, the user is asked whether only the version should be updated or the app should be upgraded as well.
// If no name is provided, all apps are checked.
func Update(appNames []string, conf config.Config) {
	logs.Verbose = conf.Verbose
	apps := parseApps(appNames, conf.VersionsFilePath)

	for _, appItem := range apps {
		logs.Info("Updating appversion for app '" + appItem.Name + "'")
		_ = appVersions.Update(appItem, conf.AppsDirPath, conf.VersionsFilePath)
	}
}
