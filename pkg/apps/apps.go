package apps

import (
	"strings"

	"github.com/thetillhoff/eac/internal/app"
	"github.com/thetillhoff/eac/pkg/logs"
)

func apps(appNames []string, versionsFilePath string) []app.App {
	apps := []app.App{}

	loadVersions(versionsFilePath)

	for _, arg := range appNames {
		wantedVersion := ""
		if strings.Contains(arg, "=") {
			splitted := strings.Split(arg, "=")
			arg, wantedVersion = splitted[0], splitted[1]
		}
		if wantedVersion == "" { // if wantedVersion is not set via `<app>=<version>`
			wantedVersion = getVersion(arg) // retrieve wantedVersion from file
		}

		//TOOD if wantedVersion is empty when installing, retrieve latest version first
		appItem := newApp(arg, app.WantedVersion(wantedVersion))
		logs.Info("app:", appItem)

		apps = append(apps, appItem)
	}

	return apps
}
