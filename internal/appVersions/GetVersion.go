package appVersions

import "github.com/thetillhoff/eac/pkg/logs"

func GetVersion(app string) string {
	if value, ok := versions[app]; ok {
		return value
	} else {
		logs.Info("No version for app '" + app + "' in versionsFile (yet).")
		return ""
	}
}
