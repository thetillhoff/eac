package apps

import (
	"strings"

	"github.com/thetillhoff/eac/internal/app"
)

func apps(appNames []string, shell string, continueOnError bool) []app.App {
	apps := []app.App{}

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

		apps = append(apps, newApp(arg, app.WantedVersion(wantedVersion), app.Shell(shell), app.ContinueOnError(continueOnError)))
	}

	return apps
}
