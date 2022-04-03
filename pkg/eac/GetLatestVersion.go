package apps

import (
	"errors"

	"github.com/thetillhoff/eac/pkg/eac/internal/apps"
	"github.com/thetillhoff/eac/pkg/eac/internal/getVersion"
	"github.com/thetillhoff/eac/pkg/eac/internal/version"
)

func GetLatestVersion(appName string) (string, error) {
	var (
		app = apps.Apps[appName]
	)

	switch app.VersionProvider.Provider {
	case version.Plaintext:
		return getVersion.GetLatestPlaintextVersion(app.VersionProvider)
	case version.Json:
		return getVersion.GetLatestJsonVersion(app.VersionProvider)
	default:
		return "", errors.New("found misconfigured app: " + appName + ". Could not load latest version")
	}
}
