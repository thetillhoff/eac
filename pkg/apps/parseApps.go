package apps

import (
	"os"
	"strings"

	"github.com/thetillhoff/eac/pkg/apps/internal/app"
	"github.com/thetillhoff/eac/pkg/apps/internal/appVersions"
	"github.com/thetillhoff/eac/pkg/logs"
)

// Parses the provided slice of strings and returns a list of apps.
// Includes version parsing and retrieval.
func parseApps(appNames []string, versionsFilePath string) []app.App {
	apps := []app.App{}

	appVersions.Load(versionsFilePath)

	for _, appName := range appNames {
		wantedVersion := ""
		if strings.Contains(appName, "==") {
			splitted := strings.Split(appName, "==")
			appName, wantedVersion = splitted[0], splitted[1]
		} else if strings.Contains(appName, "=") {
			logs.Error("Invalid appname '" + appName + "'.\nUse '==' if you want to specify a version.")
			os.Exit(1)
		}
		if wantedVersion == "" { // if wantedVersion is not set via `<app>=<version>`
			wantedVersion = appVersions.GetVersion(appName) // retrieve wantedVersion from file
		}

		//TODO If wantedVersion is empty when installing, retrieve latest version first
		appItem := newApp(appName, app.WantedVersion(wantedVersion))

		apps = append(apps, appItem)
	}

	return apps
}
