package eac

import (
	"github.com/thetillhoff/eac/pkg/eac/internal/apps"
	"github.com/thetillhoff/eac/pkg/eac/internal/getVersion"
	"github.com/thetillhoff/eac/pkg/eac/internal/version"
)

func GetLatestVersion(appName string) (string, error) {
	var (
		app = apps.Apps[appName]
	)

	if app.VersionProvider.Provider == version.Plaintext {
		return getVersion.GetLatestPlaintextVersion(app.VersionProvider)
	} else { // == version.Json
		return getVersion.GetLatestJsonVersion(app.VersionProvider)
	}
}
